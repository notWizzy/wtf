package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF5F87")).
			MarginBottom(1)

	commandStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFD700")).
			PaddingLeft(2)

	explanationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				PaddingLeft(2).
				PaddingRight(4)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF5F87")).
			Padding(1, 2).
			MarginTop(1)
)

func PrintExplanation(command string, explanation string) {
	title := titleStyle.Render("⚡ wtf")
	cmd := commandStyle.Render("$ " + command)
	body := explanationStyle.Render(explanation)
	box := boxStyle.Render(body)

	fmt.Println(title)
	fmt.Println(cmd)
	fmt.Println(box)
}
