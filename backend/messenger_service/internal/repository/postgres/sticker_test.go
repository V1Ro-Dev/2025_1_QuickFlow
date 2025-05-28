package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"quickflow/messenger_service/internal/repository/postgres"
	"quickflow/messenger_service/internal/repository/postgres-models"
	"quickflow/shared/models"
)

func TestAddStickerPack_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	stickerPack := models.StickerPack{
		Id:        uuid.New(),
		Name:      "Sample Sticker Pack",
		CreatorId: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stickers: []*models.File{
			{
				URL:         "http://example.com/sticker1.png",
				DisplayType: models.DisplayTypeSticker,
				Name:        "sticker1",
			},
		},
	}

	// Моки для вставки стикерпака и стикеров
	mock.ExpectExec(`INSERT INTO sticker_pack`).
		WithArgs(stickerPack.Id, stickerPack.Name, stickerPack.CreatedAt, stickerPack.UpdatedAt, stickerPack.CreatorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	for _, sticker := range stickerPack.Stickers {
		mock.ExpectExec(`INSERT INTO sticker`).
			WithArgs(stickerPack.Id, sticker.URL).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	// Вызов метода
	err = repo.AddStickerPack(context.Background(), stickerPack)

	// Проверка
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetStickerPack_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	packId := uuid.New()
	expectedStickerPack := postgres_models.StickerPackPostgres{
		ID:        pgtype.UUID{Bytes: packId, Valid: true},
		Name:      pgtype.Text{String: "Sample Sticker Pack", Valid: true},
		CreatorID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Stickers: []postgres_models.PostgresFile{
			{URL: pgtype.Text{String: "http://example.com/sticker1.png", Valid: true}},
		},
	}

	// Моки для получения стикерпака
	mock.ExpectQuery(`SELECT id, name, creator_id, created_at, updated_at FROM sticker_pack`).
		WithArgs(packId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "creator_id", "created_at", "updated_at"}).
			AddRow(expectedStickerPack.ID, expectedStickerPack.Name, expectedStickerPack.CreatorID, expectedStickerPack.CreatedAt, expectedStickerPack.UpdatedAt))

	// Моки для получения стикеров
	mock.ExpectQuery(`SELECT sticker_url FROM sticker WHERE sticker_pack_id = \$1`).
		WithArgs(packId).
		WillReturnRows(sqlmock.NewRows([]string{"sticker_url"}).AddRow("http://example.com/sticker1.png"))

	// Вызов метода
	stickerPack, err := repo.GetStickerPack(context.Background(), packId)

	// Проверка
	assert.NoError(t, err)
	assert.Equal(t, "Sample Sticker Pack", stickerPack.Name)
	assert.Len(t, stickerPack.Stickers, 1)
	assert.Equal(t, "http://example.com/sticker1.png", stickerPack.Stickers[0].URL)
}

func TestGetStickerPacks_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	stickerPack1 := postgres_models.StickerPackPostgres{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Name:      pgtype.Text{String: "Pack 1", Valid: true},
		CreatorID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	stickerPack2 := postgres_models.StickerPackPostgres{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Name:      pgtype.Text{String: "Pack 2", Valid: true},
		CreatorID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	// Моки для получения стикерпака
	mock.ExpectQuery(`SELECT sp.id, sp.name, sp.creator_id, sp.created_at, sp.updated_at`).
		WithArgs(10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "creator_id", "created_at", "updated_at"}).
			AddRow(stickerPack1.ID, stickerPack1.Name, stickerPack1.CreatorID, stickerPack1.CreatedAt, stickerPack1.UpdatedAt).
			AddRow(stickerPack2.ID, stickerPack2.Name, stickerPack2.CreatorID, stickerPack2.CreatedAt, stickerPack2.UpdatedAt))

	// Моки для получения стикеров
	mock.ExpectQuery(`SELECT sticker_url FROM sticker WHERE sticker_pack_id = \$1`).
		WithArgs(stickerPack1.ID).
		WillReturnRows(sqlmock.NewRows([]string{"sticker_url"}).AddRow("http://example.com/sticker1.png"))

	mock.ExpectQuery(`SELECT sticker_url FROM sticker WHERE sticker_pack_id = \$1`).
		WithArgs(stickerPack2.ID).
		WillReturnRows(sqlmock.NewRows([]string{"sticker_url"}).AddRow("http://example.com/sticker2.png"))

	// Вызов метода
	stickerPacks, err := repo.GetStickerPacks(context.Background(), uuid.Nil, 10, 0)

	// Проверка
	assert.NoError(t, err)
	assert.Len(t, stickerPacks, 2)
	assert.Equal(t, "Pack 1", stickerPacks[0].Name)
	assert.Equal(t, "http://example.com/sticker1.png", stickerPacks[0].Stickers[0].URL)
	assert.Equal(t, "Pack 2", stickerPacks[1].Name)
	assert.Equal(t, "http://example.com/sticker2.png", stickerPacks[1].Stickers[0].URL)
}

func TestDeleteStickerPack_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	userId := uuid.New()
	packId := uuid.New()

	// Моки для проверки владельца стикерпака
	mock.ExpectQuery(`SELECT creator_id FROM sticker_pack WHERE id = \$1`).
		WithArgs(packId).
		WillReturnRows(sqlmock.NewRows([]string{"creator_id"}).AddRow(userId))

	// Моки для удаления стикеров
	mock.ExpectExec(`DELETE FROM sticker WHERE sticker_pack_id = \$1`).
		WithArgs(packId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Моки для удаления стикерпака
	mock.ExpectExec(`DELETE FROM sticker_pack WHERE id = \$1`).
		WithArgs(packId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Вызов метода
	err = repo.DeleteStickerPack(context.Background(), userId, packId)

	// Проверка
	assert.NoError(t, err)
}

func TestGetStickerPackByName_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	stickerPackName := "Sample Sticker Pack"
	expectedStickerPack := postgres_models.StickerPackPostgres{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Name:      pgtype.Text{String: stickerPackName, Valid: true},
		CreatorID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	// Моки для получения стикерпака по имени
	mock.ExpectQuery(`SELECT id, name, creator_id, created_at, updated_at FROM sticker_pack`).
		WithArgs(stickerPackName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "creator_id", "created_at", "updated_at"}).
			AddRow(expectedStickerPack.ID, expectedStickerPack.Name, expectedStickerPack.CreatorID, expectedStickerPack.CreatedAt, expectedStickerPack.UpdatedAt))

	// Моки для получения стикеров
	mock.ExpectQuery(`SELECT sticker_url FROM sticker WHERE sticker_pack_id = \$1`).
		WithArgs(expectedStickerPack.ID).
		WillReturnRows(sqlmock.NewRows([]string{"sticker_url"}).AddRow("http://example.com/sticker1.png"))

	// Вызов метода
	stickerPack, err := repo.GetStickerPackByName(context.Background(), stickerPackName)

	// Проверка
	assert.NoError(t, err)
	assert.Equal(t, stickerPackName, stickerPack.Name)
	assert.Equal(t, "http://example.com/sticker1.png", stickerPack.Stickers[0].URL)
}

func TestBelongsTo_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание репозитория
	repo := postgres.NewPostgresStickerRepository(db)

	// Тестовые данные
	userId := uuid.New()
	packId := uuid.New()

	// Моки для проверки принадлежности стикерпака
	mock.ExpectQuery(`SELECT creator_id FROM sticker_pack WHERE id = \$1`).
		WithArgs(packId).
		WillReturnRows(sqlmock.NewRows([]string{"creator_id"}).AddRow(userId))

	// Вызов метода
	belongs, err := repo.BelongsTo(context.Background(), userId, packId)

	// Проверка
	assert.NoError(t, err)
	assert.True(t, belongs)
}
