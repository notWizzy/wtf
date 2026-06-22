package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF5F87"))

	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700"))

	explanationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF"))

	successStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF87"))
)

func PrintExplanation(command string, stderr string, explanation string) {
	fmt.Println()
	fmt.Println(labelStyle.Render("⚡ wtf") + "  " + commandStyle.Render("$ "+command))
	fmt.Println()
	fmt.Println(explanationStyle.Render(explanation))
	fmt.Println()
}

func PrintSuccess(command string) {
	fmt.Println()
	fmt.Println(successStyle.Render("✓ " + command + " ran successfully"))
	fmt.Println()
}
