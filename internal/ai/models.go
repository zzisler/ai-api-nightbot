package ai

import (
	"net/http"
)

type Client struct {
	apiKey       string
	httpClient   *http.Client
	proxyUrl     string
	systemPrompt string
}

type Request struct {
	Model               string       `json:"model"`
	Messages            []ReqMessage `json:"messages"`
	Temperature         float64      `json:"temperature"`
	MaxCompletionTokens int          `json:"max_completion_tokens"`
}

type ReqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"` // это можно убрать я думаю
	Index        int     `json:"index"`         // и это тоже можно убрать
	Message      Message `json:"message"`
}

type Message struct {
	Content   string `json:"content"`
	Reasoning string `json:"reasoning"`
	Role      string `json:"role"`
}
