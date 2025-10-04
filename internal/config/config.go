package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"github.com/jdschrack/go-agents/internal/log"
)

type Config struct {
	Port          string
	GCPProjectID  string
	DatabasePath  string
	GithubURL     string
	VertexAIKey   string
	VertexAIModel string
}

func LoadConfig(ctx context.Context) *Config {
	logger := log.FromContext(ctx)
	err := godotenv.Load()

	err = godotenv.Overload(".env.local")
	if err != nil {
		logger.Warn().Msgf("No .env.local file found, proceeding without machine environment variables")
	}

	getEnv := func(key, defaultValue string) string {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
		return defaultValue
	}

	return &Config{
		Port:          getEnv("PORT", "8080"),
		GCPProjectID:  getEnv("GCP_PROJECT_ID", ""),
		DatabasePath:  getEnv("DATABASE_PATH", "data.db"),
		GithubURL:     getEnv("GITHUB_URL", ""),
		VertexAIKey:   getEnv("VERTEX_AI_KEY", ""),
		VertexAIModel: getEnv("VERTEX_AI_MODEL", ""),
	}
}
