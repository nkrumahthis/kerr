package cmd

import (
	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
)

var listCmd = &cobra.Command{
	Use:   "list [tags...]",
	Short: "List principles for doing a great job",
	Long: `Show actionable items and reminders based on the book "How to Be Great at Your Job". 
	Provide tags to filter principles, e.g., 'kerr list tasks' or 'kerr list meetings'.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.ListActions()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}