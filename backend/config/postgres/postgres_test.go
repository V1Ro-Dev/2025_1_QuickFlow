package postgres_config

import (
	"log"
	"os"
	"testing"
)

func TestNewPostgresConfig_WithEnvVar(t *testing.T) {
	// Устанавливаем переменную окружения
	err := os.Setenv("DATABASE_URL", "postgresql://custom-user:password@localhost:5432/custom_db")
	if err != nil {
		log.Fatalf("failed to set DATABASE_URL: %v", err)
	}
	defer func() {
		err = os.Unsetenv("DATABASE_URL")
		if err != nil {
			log.Fatalf("failed to unset DATABASE_URL: %v", err)
		}
	}() // Очищаем после теста

	// Создаем новый объект конфигурации
	postgresConfig := NewPostgresConfig()

	// Проверяем, что возвращается правильный URL
	expectedURL := "postgresql://custom-user:password@localhost:5432/custom_db"
	if postgresConfig.GetURL() != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, postgresConfig.GetURL())
	}
}

func TestNewPostgresConfig_WithoutEnvVar(t *testing.T) {
	// Убираем переменную окружения, если она установлена
	err := os.Unsetenv("DATABASE_URL")
	if err != nil {
		log.Fatalf("failed to unset DATABASE_URL: %v", err)
	}

	// Создаем новый объект конфигурации
	postgresConfig := NewPostgresConfig()

	// Проверяем, что возвращается URL по умолчанию
	expectedURL := defaultDataBaseURL
	if postgresConfig.GetURL() != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, postgresConfig.GetURL())
	}
}
