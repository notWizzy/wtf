package history

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLastCommand() (string, error) {
	historyFile := os.Getenv("HOME") + "/.zsh_history"

	file, err := os.Open(historyFile)
	if err != nil {
		return "", fmt.Errorf("could not open shell history: %w", err)
	}
	defer file.Close()

	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lastLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading history file: %w", err)
	}

	if lastLine == "" {
		return "", fmt.Errorf("no commands found in history")
	}

	if strings.Contains(lastLine, ";") {
		parts := strings.SplitN(lastLine, ";", 2)
		lastLine = parts[1]
	}

	return lastLine, nil
}
