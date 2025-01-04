package cmd

import (
	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
)

var doneCmd = &cobra.Command{
	Use: "done [action codes]",
	Short: "Log completed actions and update XP score.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		core.LogAchievements(args, true)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
