package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func NewClient(apiKey, prompt string) (*Client, error) {

	httpClient := &http.Client{
		Timeout: 3500 * time.Millisecond,
	}

	proxyUrl := os.Getenv("PROXY_URL")
	if proxyUrl != "" {
		parsed, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}
		httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(parsed),
		}
	}

	return &Client{
		apiKey:       apiKey,
		httpClient:   httpClient,
		proxyUrl:     proxyUrl,
		systemPrompt: prompt,
	}, nil

}

func (c *Client) AiReq(messages []ReqMessage) (*Message, error) {
	url := "https://api.cerebras.ai/v1/chat/completions"

	messages = append([]ReqMessage{
		{
			Role:    "system",
			Content: string(c.systemPrompt),
		},
	}, messages...)

	data := Request{
		Model:               "gpt-oss-120b",
		Messages:            messages,
		Temperature:         0.9,
		MaxCompletionTokens: 500,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cerebras api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("empty choices, raw: %s", string(body))
	}

	return &response.Choices[0].Message, nil

}
