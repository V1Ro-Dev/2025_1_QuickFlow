package minio

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/sync/errgroup"

	minioconfig "quickflow/file_service/config/minio"
	threadsafeslice "quickflow/pkg/thread-safe-slice"
	"quickflow/shared/logger"
	"quickflow/shared/models"
)

type MinioRepository struct {
	client                *minio.Client
	PostsBucketName       string
	AttachmentsBucketName string
	ProfileBucketName     string
	StickerBuckerName     string
	PublicUrlRoot         string
}

func NewMinioRepository(cfg *minioconfig.MinioConfig) (*MinioRepository, error) {
	client, err := minio.New(cfg.MinioInternalEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioRootUser, cfg.MinioRootPassword, ""),
		Secure: cfg.MinioUseSSL,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create minio client: %v", err)
	}

	exists, err := client.BucketExists(context.Background(), cfg.PostsBucketName)
	if err != nil {
		return nil, fmt.Errorf("could not check if bucket exists: %v", err)
	}

	if !exists {
		err = client.MakeBucket(context.Background(), cfg.PostsBucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("could not create bucket: %v", err)
		}
	}

	return &MinioRepository{
		client:                client,
		PostsBucketName:       cfg.PostsBucketName,
		AttachmentsBucketName: cfg.AttachmentsBucketName,
		ProfileBucketName:     cfg.ProfileBucketName,
		StickerBuckerName:     cfg.StickerBuckerName,
		PublicUrlRoot:         fmt.Sprintf("%s://%s", cfg.Scheme, cfg.MinioPublicEndpoint),
	}, nil
}

// UploadFile uploads file to MinIO and returns a public URL.
func (m *MinioRepository) UploadFile(ctx context.Context, file *models.File) (string, error) {
	var err error
	uuID := uuid.New()
	fileName := uuID.String() + file.Ext

	if file.DisplayType == models.DisplayTypeSticker {
		_, err = m.client.PutObject(ctx, m.StickerBuckerName, fileName, file.Reader, file.Size, minio.PutObjectOptions{
			ContentType: file.MimeType,
		})
	} else {
		_, err = m.client.PutObject(ctx, m.PostsBucketName, fileName, file.Reader, file.Size, minio.PutObjectOptions{
			ContentType: file.MimeType,
		})
	}
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("could not upload file %v: %v", file.Name, err))
		return "", fmt.Errorf("could not upload file: %v", err)
	}

	publicURL := fmt.Sprintf("%s/%s/%s", m.PublicUrlRoot, m.PostsBucketName, fileName)
	logger.Info(ctx, fmt.Sprintf("File successfully loaded: %v, url: %v", file.Name, publicURL))
	return publicURL, nil
}

func (m *MinioRepository) UploadManyImages(ctx context.Context, files []*models.File) ([]string, error) {
	urls := threadsafeslice.NewThreadSafeSliceN[string](len(files))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)

	for i, file := range files {
		i := i
		file := file // https://golang.org/doc/faq#closures_and_goroutines
		uuID := uuid.New()
		fileName := uuID.String() + file.Ext

		wg.Go(func() error {
			var err error
			if file.DisplayType == models.DisplayTypeSticker {
				_, err = m.client.PutObject(ctx, m.StickerBuckerName, fileName, file.Reader, file.Size, minio.PutObjectOptions{
					ContentType: file.MimeType,
				})
				if err != nil {
					return fmt.Errorf("could not upload file: %v, err: %v", file.Name, err)
				}
			} else {
				_, err = m.client.PutObject(ctx, m.PostsBucketName, fileName, file.Reader, file.Size, minio.PutObjectOptions{
					ContentType: file.MimeType,
				})
				if err != nil {
					return fmt.Errorf("could not upload file: %v, err: %v", file.Name, err)
				}
			}

			publicURL := fmt.Sprintf("%s/%s/%s", m.PublicUrlRoot, m.PostsBucketName, fileName)
			err = urls.SetByIdx(i, publicURL)
			if err != nil {
				return fmt.Errorf("could not upload file: %v, err: %v", file.Name, err)
			}
			return nil
		})
	}

	if err := wg.Wait(); err != nil {
		return nil, err
	}
	return urls.GetSliceCopy(), nil
}

// GetFileURL returns a public URL for the file.
func (m *MinioRepository) GetFileURL(_ context.Context, fileName string) (string, error) {
	return fmt.Sprintf("%s/%s/%s", m.PublicUrlRoot, m.PostsBucketName, fileName), nil
}

// DeleteFile deletes a file from MinIO.
func (m *MinioRepository) DeleteFile(ctx context.Context, fileName string) error {
	err := m.client.RemoveObject(ctx, m.PostsBucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("could not delete file: %v", err)
	}
	return nil
}
