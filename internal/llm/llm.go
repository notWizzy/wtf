package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "http://localhost:11434/api/chat"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Response struct {
	Message Message `json:"message"`
}

func Explain(errorText string) (string, error) {
	body := Request{
		Model:  "llama3.2",
		Stream: false,
		Messages: []Message{
			{
				Role: "user",
				Content: fmt.Sprintf(`You are a terminal error explainer. Be concise and direct.

Given this error output:
%s

Respond in exactly this format, no markdown, no asterisks, no bullet symbols:

WHAT HAPPENED: one sentence explaining what went wrong
WHY: one sentence explaining the root cause
FIX: one sentence telling exactly what to do`, errorText),
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to build request: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed — is Ollama running? Try: ollama serve: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var result Response
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if result.Message.Content == "" {
		return "", fmt.Errorf("empty response from Ollama")
	}

	return result.Message.Content, nil
}
