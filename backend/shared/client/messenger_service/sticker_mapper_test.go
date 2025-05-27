package messenger_service

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"quickflow/shared/client/file_service"
	"quickflow/shared/models"
	"quickflow/shared/proto/file_service/mocks"
	pb "quickflow/shared/proto/messenger_service"
)

func TestMapStickerPackToProto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокируем вызов file_service
	file_service.ModelFilesToProto = func(files []models.File) []*pb.File {
		return []*pb.File{
			{Id: "sticker1-id"},
			{Id: "sticker2-id"},
		}
	}
	defer func() {
		file_service.ModelFilesToProto = file_service.ModelFilesToProtoOriginal
	}()

	now := time.Now()
	tests := []struct {
		name     string
		input    *models.StickerPack
		expected *pb.StickerPack
	}{
		{
			name: "Full StickerPack",
			input: &models.StickerPack{
				Id:        uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				Name:      "test-pack",
				CreatorId: uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
				CreatedAt: now,
				UpdatedAt: now,
				Stickers: []models.File{
					{Name: "sticker1"},
					{Name: "sticker2"},
				},
			},
			expected: &pb.StickerPack{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				Name:      "test-pack",
				CreatorId: "123e4567-e89b-12d3-a456-426614174001",
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
				Stickers: []*pb.File{
					{Id: "sticker1-id"},
					{Id: "sticker2-id"},
				},
			},
		},
		{
			name: "Empty Stickers",
			input: &models.StickerPack{
				Id:        uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				Name:      "empty-pack",
				CreatorId: uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
				CreatedAt: now,
				UpdatedAt: now,
				Stickers:  []models.File{},
			},
			expected: &pb.StickerPack{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				Name:      "empty-pack",
				CreatorId: "123e4567-e89b-12d3-a456-426614174001",
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
				Stickers:  []*pb.File{},
			},
		},
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapStickerPackToProto(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMapStickerPacksToProto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	tests := []struct {
		name     string
		input    []*models.StickerPack
		expected []*pb.StickerPack
	}{
		{
			name: "Multiple Packs",
			input: []*models.StickerPack{
				{
					Id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
					Name: "pack1",
				},
				{
					Id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
					Name: "pack2",
				},
			},
			expected: []*pb.StickerPack{
				{
					Id:   "123e4567-e89b-12d3-a456-426614174000",
					Name: "pack1",
				},
				{
					Id:   "123e4567-e89b-12d3-a456-426614174001",
					Name: "pack2",
				},
			},
		},
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "Empty Slice",
			input:    []*models.StickerPack{},
			expected: []*pb.StickerPack{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapStickerPacksToProto(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMapProtoToStickerPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокируем вызов file_service
	mockFileService := file_service.NewMockFileService(ctrl)
	file_service.ProtoFilesToModels = func(files []*pb.File) []models.File {
		return []models.File{
			{Name: "sticker1"},
			{Name: "sticker2"},
		}
	}
	defer func() {
		file_service.ProtoFilesToModels = file_service.ProtoFilesToModelsOriginal
	}()

	now := time.Now()
	tests := []struct {
		name        string
		input       *pb.StickerPack
		expected    *models.StickerPack
		expectError bool
	}{
		{
			name: "Valid StickerPack",
			input: &pb.StickerPack{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				Name:      "test-pack",
				CreatorId: "123e4567-e89b-12d3-a456-426614174001",
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
				Stickers: []*pb.File{
					{Id: "sticker1-id"},
					{Id: "sticker2-id"},
				},
			},
			expected: &models.StickerPack{
				Id:        uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				Name:      "test-pack",
				CreatorId: uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
				CreatedAt: now,
				UpdatedAt: now,
				Stickers: []models.File{
					{Name: "sticker1"},
					{Name: "sticker2"},
				},
			},
		},
		{
			name: "Invalid UUID in Id",
			input: &pb.StickerPack{
				Id: "invalid-uuid",
			},
			expectError: true,
		},
		{
			name: "Invalid UUID in CreatorId",
			input: &pb.StickerPack{
				Id:        "123e4567-e89b-12d3-a456-426614174000",
				CreatorId: "invalid-uuid",
			},
			expectError: true,
		},
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MapProtoToStickerPack(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.input == nil {
					assert.Nil(t, result)
				} else {
					assert.Equal(t, tt.expected.Id, result.Id)
					assert.Equal(t, tt.expected.Name, result.Name)
					assert.Equal(t, tt.expected.CreatorId, result.CreatorId)
					assert.Equal(t, tt.expected.CreatedAt.Unix(), result.CreatedAt.Unix())
					assert.Equal(t, tt.expected.UpdatedAt.Unix(), result.UpdatedAt.Unix())
					assert.Equal(t, len(tt.expected.Stickers), len(result.Stickers))
				}
			}
		})
	}
}

func TestMapProtoToStickerPack_FileConversionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Мокируем вызов file_service чтобы вернуть ошибку
	mockFileService := file_service.NewMockFileService(ctrl)
	file_service.ProtoFilesToModels = func(files []*pb.File) []models.File {
		return nil
	}
	defer func() {
		file_service.ProtoFilesToModels = file_service.ProtoFilesToModelsOriginal
	}()

	input := &pb.StickerPack{
		Id:        "123e4567-e89b-12d3-a456-426614174000",
		CreatorId: "123e4567-e89b-12d3-a456-426614174001",
		Stickers:  []*pb.File{{Id: "sticker1-id"}},
	}

	_, err := MapProtoToStickerPack(input)
	assert.NoError(t, err) // Ожидаем что даже с nil стикерами маппинг продолжится
}
