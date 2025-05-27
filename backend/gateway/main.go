package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"quickflow/config"
	"quickflow/config/cors"
	minioConfig "quickflow/config/minio"
	postgresConfig "quickflow/config/postgres"
	redisConfig "quickflow/config/redis"
	"quickflow/config/server"
	validationCfg "quickflow/config/validation"
	"quickflow/gateway/internal"
)

func resolveConfigPath(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	if _, ok := os.LookupEnv("RUNNING_IN_CONTAINER"); ok {
		return filepath.Join("/config", rel)
	}
	return filepath.Join("../deploy/config", rel)
}

func initCfg() (*config.Config, error) {

	serverConfigPath := flag.String("server-config", "feeder/config.toml", "Path to config file")
	corsConfigPath := flag.String("cors-config", "cors/config.toml", "Path to CORS config file")
	minioConfigPath := flag.String("minio-config", "minio/config.toml", "Path to Minio config file")
	validationConfig := flag.String("validation-config", "validation/config.toml", "Path to Validation config file")
	flag.Parse()

	serverCfg, err := server_config.Parse(resolveConfigPath(*serverConfigPath))
	if err != nil {
		return nil, fmt.Errorf("failed to load project server configuration: %v", err)
	}

	corsCfg, err := cors_config.ParseCORS(resolveConfigPath(*corsConfigPath))
	if err != nil {
		return nil, fmt.Errorf("failed to load project CORS configuration: %v", err)
	}

	minioCfg, err := minioConfig.ParseMinio(resolveConfigPath(*minioConfigPath))
	if err != nil {
		return nil, fmt.Errorf("failed to load project minio configuration: %v", err)
	}

	validationConf, err := validationCfg.NewValidationConfig(resolveConfigPath(*validationConfig))
	if err != nil {
		return nil, fmt.Errorf("failed to load project validation configuration: %v", err)
	}

	postgresCfg := postgresConfig.NewPostgresConfig()
	redisCfg := redisConfig.NewRedisConfig()

	return &config.Config{
		PostgresConfig:   postgresCfg,
		ServerConfig:     serverCfg,
		CORSConfig:       corsCfg,
		MinioConfig:      minioCfg,
		RedisConfig:      redisCfg,
		ValidationConfig: validationConf,
	}, nil
}

func main() {

	appCfg, err := initCfg()
	if err != nil {
		log.Fatalf("failed to initialize configuration: %v", err)
	}

	if err = internal.Run(appCfg); err != nil {
		log.Fatalf("failed to start QuickFlow: %v", err)
	}
}
