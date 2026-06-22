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
	"os/exec"
	"strings"

	"github.com/notWizzy/wtf/internal/llm"
	"github.com/notWizzy/wtf/internal/output"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wtf [command]",
	Short: "Explains your last terminal error in plain English",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("wtf is working!")
	// },
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		commandArgs := args[1:]
		fullCommand := strings.Join(args, " ")

		exitCode, stderr := runCommand(command, commandArgs)

		if exitCode == 0 {
			output.PrintSuccess(fullCommand)
			return
		}

		explanation, err := llm.Explain(stderr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error from LLM:", err)
			os.Exit(1)
		}

		output.PrintExplanation(fullCommand, stderr, explanation)
	},
}

func runCommand(command string, args []string) (int, string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout

	stderrBytes, err := cmd.StderrPipe()
	if err != nil {
		return 1, fmt.Sprintf("failed to capture stderr: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return 1, fmt.Sprintf("failed to start command: %v", err)
	}

	buf := new(strings.Builder)
	buffer := make([]byte, 1024)
	for {
		n, err := stderrBytes.Read(buffer)
		if n > 0 {
			buf.Write(buffer[:n])
		}
		if err != nil {
			break
		}
	}

	cmd.Wait()

	exitCode := 0
	if cmd.ProcessState != nil {
		exitCode = cmd.ProcessState.ExitCode()
	}

	return exitCode, buf.String()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
