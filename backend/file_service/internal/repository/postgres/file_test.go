package postgres_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"quickflow/file_service/internal/repository/postgres"
	"quickflow/shared/models"
	"testing"
)

func TestAddFileRecord_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра репозитория
	repo := postgres.NewPostgresFileRepository(db)

	// Тестовые данные
	file := &models.File{URL: "http://example.com/file1", Name: "file1"}

	// Ожидаем выполнение SQL-запроса
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(file.URL, file.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Вызов метода
	err = repo.AddFileRecord(context.Background(), file)

	// Проверка
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddFileRecord_SQLQueryError(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра репозитория
	repo := postgres.NewPostgresFileRepository(db)

	// Тестовые данные
	file := &models.File{URL: "http://example.com/file1", Name: "file1"}

	// Ожидаем выполнение SQL-запроса с ошибкой
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(file.URL, file.Name).
		WillReturnError(errors.New("query error"))

	// Вызов метода
	err = repo.AddFileRecord(context.Background(), file)

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "query error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddFilesRecords_Success(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра репозитория
	repo := postgres.NewPostgresFileRepository(db)

	// Тестовые данные
	files := []*models.File{
		{URL: "http://example.com/file1", Name: "file1"},
		{URL: "http://example.com/file2", Name: "file2"},
	}

	// Ожидаем начало транзакции
	mock.ExpectBegin()

	// Ожидаем выполнение SQL-запросов для каждого файла
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(files[0].URL, files[0].Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(files[1].URL, files[1].Name).
		WillReturnResult(sqlmock.NewResult(2, 1))

	// Ожидаем успешное завершение транзакции
	mock.ExpectCommit()

	// Вызов метода
	err = repo.AddFilesRecords(context.Background(), files)

	// Проверка
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddFilesRecords_TxError(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра репозитория
	repo := postgres.NewPostgresFileRepository(db)

	// Тестовые данные
	files := []*models.File{
		{URL: "http://example.com/file1", Name: "file1"},
		{URL: "http://example.com/file2", Name: "file2"},
	}

	// Ожидаем начало транзакции
	mock.ExpectBegin()

	// Ожидаем выполнение первого SQL-запроса
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(files[0].URL, files[0].Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Ожидаем, что второй запрос вызовет ошибку
	mock.ExpectExec(`INSERT INTO files \(file_url, filename\) VALUES \(\$1, \$2\)`).
		WithArgs(files[1].URL, files[1].Name).
		WillReturnError(errors.New("insert error"))

	// Ожидаем откат транзакции
	mock.ExpectRollback()

	// Вызов метода
	err = repo.AddFilesRecords(context.Background(), files)

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "insert error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddFilesRecords_FileIsNil(t *testing.T) {
	// Инициализация мока
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра репозитория
	repo := postgres.NewPostgresFileRepository(db)

	// Вызов метода с nil файлом
	err = repo.AddFilesRecords(context.Background(), nil)

	// Проверка
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
