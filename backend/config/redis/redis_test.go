package redis_config

import (
	"log"
	"os"
	"testing"
)

func TestNewRedisConfig_WithEnvVar(t *testing.T) {
	// Устанавливаем переменную окружения
	err := os.Setenv("REDIS_URL", "redis://custom-redis:6379")
	if err != nil {
		log.Fatalf("failed to set REDIS_URL: %v", err)
	}
	defer func() {
		err := os.Unsetenv("REDIS_URL")
		if err != nil {
			log.Fatalf("failed to unset REDIS_URL: %v", err)
		}
	}() // Очищаем после теста

	// Создаем новый объект конфигурации
	redisConfig := NewRedisConfig()

	// Проверяем, что возвращается правильный URL
	expectedURL := "redis://custom-redis:6379"
	if redisConfig.GetURL() != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, redisConfig.GetURL())
	}
}

func TestNewRedisConfig_WithoutEnvVar(t *testing.T) {
	// Убираем переменную окружения, если она установлена
	err := os.Unsetenv("REDIS_URL")
	if err != nil {
		log.Fatalf("failed to unset REDIS_URL: %v", err)
	}

	// Создаем новый объект конфигурации
	redisConfig := NewRedisConfig()

	// Проверяем, что возвращается URL по умолчанию
	expectedURL := backupURL
	if redisConfig.GetURL() != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, redisConfig.GetURL())
	}
}
