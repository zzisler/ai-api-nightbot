package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey         string
	SecretToken    string
	ProxyUrl       string
	Prompt         string
	PromptTemplate string
}

var templates = map[string]string{
	"troll": "%s задал вопрос: %s. высмей его.",
}

func Load() (*Config, error) {

	godotenv.Load()

	cfg := &Config{}

	cfg.ApiKey = os.Getenv("API_KEY")
	if cfg.ApiKey == "" {
		return nil, fmt.Errorf("API_KEY is not set")
	}

	cfg.SecretToken = os.Getenv("SECRET_TOKEN")
	if cfg.SecretToken == "" {
		return nil, fmt.Errorf("SECRET_TOKEN is not set")
	}

	cfg.ProxyUrl = os.Getenv("PROXY_URL")

	data, err := os.ReadFile("internal/prompt/prompt.txt")
	if err != nil {
		return nil, fmt.Errorf("PROMPT error: %s", err)
	}
	cfg.Prompt = string(data)

	tmplKey := os.Getenv("PROMPT_TEMPLATE")
	tmpl, ok := templates[tmplKey]
	if !ok {
		return nil, fmt.Errorf("TEMPLATE error: %s", tmplKey)
	}
	cfg.PromptTemplate = tmpl

	log.Println("config loaded")

	return cfg, nil

}
