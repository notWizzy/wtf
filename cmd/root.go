package cmd

// import (
// 	"fmt"
// 	"os"

// 	"github.com/notWizzy/wtf/internal/llm"
// 	"github.com/spf13/cobra"
// )

import (
	"fmt"
	"os"

	"github.com/notWizzy/wtf/internal/history"
	"github.com/notWizzy/wtf/internal/llm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtf",
	Short: "Explains your last terminal error in plain English",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("wtf is working!")
	// },
	Run: func(cmd *cobra.Command, args []string) {
		lastCmd, err := history.GetLastCommand()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading history:", err)
			os.Exit(1)
		}

		fmt.Println("Last command:", lastCmd)
		fmt.Println("Asking Ollama...")

		explanation, err := llm.Explain(lastCmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error from LLM:", err)
			os.Exit(1)
		}

		fmt.Println(explanation)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
