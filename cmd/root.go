package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtf",
	Short: "Explains your last terminal error in plain English",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wtf is working!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
