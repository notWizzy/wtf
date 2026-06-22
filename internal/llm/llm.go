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
				Content: fmt.Sprintf(`You are a senior software engineer explaining a terminal failure to a fellow engineer. You write with clarity, precision, and natural human tone — no robotic phrasing, no filler words, no technical jargon unless it adds meaning.

The command failed with this error:

%s

You MUST respond in exactly this format and no other. Do not deviate. Do not add any text before or after. Do not remove the labels.

WHAT HAPPENED: [one crisp sentence naming exactly what failed — the specific file, path, module, or resource. Never say "the command failed" generically.]

WHY: [one or two sentences explaining the precise root cause. If the error contains a specific path, file name, or error code that reveals the cause, use it naturally — do not reference error codes by name like "errno 2", instead say what they mean: "the file does not exist at that path".]

FIX: [one or two sentences giving the most direct fix using standard shell commands. Prefer pwd, ls, touch, mkdir, pip install, brew install over language-specific workarounds. If a file is missing, give the exact touch or mkdir command with the full path from the error.]`, errorText),
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
