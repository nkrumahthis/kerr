package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kerr",
	Short: "Be great at your job, Get things done, get the credit, get ahead",
	Long: "Kerr is a CLI tool that helps you get things done, get the credit, and get ahead in your career",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `kerr [command]` to interact with the tool.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}