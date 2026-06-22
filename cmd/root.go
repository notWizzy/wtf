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
		result, err := llm.Explain("ModuleNotFoundError: No module named 'pandas'")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
