package grpc

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dto "quickflow/shared/client/file_service"
	"quickflow/shared/logger"
	"quickflow/shared/models"
	pb "quickflow/shared/proto/file_service"
)

type FileUseCase interface {
	UploadFile(ctx context.Context, fileModel *models.File) (string, error)
	UploadManyMedia(ctx context.Context, files []*models.File) ([]string, error)
	GetFileURL(ctx context.Context, filename string) (string, error)
	DeleteFile(ctx context.Context, filename string) error
}

type FileServiceServer struct {
	pb.UnimplementedFileServiceServer
	fileUC FileUseCase
}

func NewFileServiceServer(fileUC FileUseCase) *FileServiceServer {
	return &FileServiceServer{fileUC: fileUC}
}

func (s *FileServiceServer) UploadFile(stream pb.FileService_UploadFileServer) error {
	var (
		fileInfo *pb.File
		tempFile *os.File
	)

	ctx := stream.Context()
	logger.Info(ctx, "Started streaming UploadFile request")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					return
				}
			}(tempFile.Name())

			f, err := os.Open(tempFile.Name())
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to reopen temp file: %v", err))
				return err
			}
			defer func(f *os.File) {
				err = f.Close()
				if err != nil {
					return
				}
			}(f)

			data, err := io.ReadAll(f)
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to read file: %v", err))
				return err
			}

			fileInfo.File = data
			fileURL, err := s.fileUC.UploadFile(ctx, dto.ProtoFileToModel(fileInfo))
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Upload usecase failed: %v", err))
				return err
			}

			logger.Info(ctx, fmt.Sprintf("File uploaded successfully: %s", fileURL))
			return stream.SendAndClose(&pb.UploadFileResponse{FileUrl: fileURL})
		}

		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Error receiving stream: %v", err))
			return err
		}

		switch x := req.Data.(type) {
		case *pb.UploadFileRequest_Info:
			fileInfo = x.Info
			tempFile, err = os.CreateTemp("", "upload-*")
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to create temp file: %v", err))
				return err
			}
			defer func(tempFile *os.File) {
				err := tempFile.Close()
				if err != nil {
					return
				}
			}(tempFile)

		case *pb.UploadFileRequest_Chunk:
			if tempFile == nil {
				return status.Errorf(codes.InvalidArgument, "FileInfo must be sent before chunks")
			}
			_, err := tempFile.Write(x.Chunk)
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to write chunk: %v", err))
				return err
			}
		}
	}
}

func (s *FileServiceServer) UploadManyFiles(stream pb.FileService_UploadManyFilesServer) error {
	var (
		currentInfo *pb.File
		tempFile    *os.File
		ctx         = stream.Context()
	)

	defer func() {
		if tempFile != nil {
			err := tempFile.Close()
			if err != nil {
				return
			}
			err = os.Remove(tempFile.Name())
			if err != nil {
				return
			} // чистим tmp
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Обработка последнего файла, если был начат
			if currentInfo != nil && tempFile != nil {
				fileURL, err := s.finalizeUploadedFile(ctx, currentInfo, tempFile)
				if err != nil {
					logger.Error(ctx, fmt.Sprintf("Error finalizing last file: %v", err))
					return err
				}
				if err := stream.Send(&pb.UploadFileResponse{
					FileUrl: fileURL,
				}); err != nil {
					logger.Error(ctx, fmt.Sprintf("Failed to send response: %v", err))
					return err
				}
			}
			logger.Info(ctx, "All files received and processed")
			return nil
		}
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Error receiving upload stream: %v", err))
			return err
		}

		switch data := req.Data.(type) {
		case *pb.UploadFileRequest_Info:
			// Завершаем предыдущий файл
			if currentInfo != nil && tempFile != nil {
				fileURL, err := s.finalizeUploadedFile(ctx, currentInfo, tempFile)
				if err != nil {
					logger.Error(ctx, fmt.Sprintf("Error finalizing file: %v", err))
					return err
				}
				if err := stream.Send(&pb.UploadFileResponse{
					FileUrl: fileURL,
				}); err != nil {
					logger.Error(ctx, fmt.Sprintf("Failed to send response: %v", err))
					return err
				}
				err = tempFile.Close()
				if err != nil {
					return err
				}
				err = os.Remove(tempFile.Name())
				if err != nil {
					return err
				}
			}

			// Начинаем новый файл
			currentInfo = data.Info
			tempFile, err = os.CreateTemp("", "upload-*")
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to create temp file: %v", err))
				return err
			}

		case *pb.UploadFileRequest_Chunk:
			if tempFile == nil {
				return status.Errorf(codes.InvalidArgument, "FileInfo must be sent before chunks")
			}
			_, err := tempFile.Write(data.Chunk)
			if err != nil {
				logger.Error(ctx, fmt.Sprintf("Failed to write chunk: %v", err))
				return err
			}
		}
	}
}

// finalizeUploadedFile обрабатывает файл после его получения
func (s *FileServiceServer) finalizeUploadedFile(
	ctx context.Context,
	info *pb.File,
	tempFile *os.File,
) (string, error) {
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			return
		}
	}(tempFile)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			return
		}
	}(tempFile.Name())

	f, err := os.Open(tempFile.Name())
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Failed to reopen temp file: %v", err))
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	data, err := io.ReadAll(f)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Failed to read file: %v", err))
		return "", err
	}

	info.File = data

	fileURL, err := s.fileUC.UploadFile(ctx, dto.ProtoFileToModel(info))
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Upload usecase failed: %v", err))
		return "", err
	}

	return fileURL, nil
}

func (s *FileServiceServer) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	logger.Info(ctx, "Received DeleteFile request")

	err := s.fileUC.DeleteFile(ctx, req.FileUrl)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Failed to delete file: %v", err))
		return &pb.DeleteFileResponse{Success: false}, err
	}

	logger.Info(ctx, "Successfully deleted file")
	return &pb.DeleteFileResponse{Success: true}, nil
}
