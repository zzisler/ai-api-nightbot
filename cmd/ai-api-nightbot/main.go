package main

import (
	"ai-nightbot/internal/ai"
	"ai-nightbot/internal/config"
	"ai-nightbot/internal/handler"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	aiClient, err := ai.NewClient(cfg.ApiKey, cfg.Prompt)
	if err != nil {
		log.Fatal(err)
	}

	h := &handler.Handler{
		AiClient: aiClient,
		Cfg:      cfg,
	}

	http.HandleFunc("/api/troll", h.Troll)

	http.ListenAndServe(":7777", nil)

}
