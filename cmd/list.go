package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
)

var listCmd = &cobra.Command{
	Use:   "list [tags...]",
	Short: "List principles for doing a great job",
	Long: `Show actionable items and reminders based on the book "How to Be Great at Your Job". 
	Provide tags to filter principles, e.g., 'kerr list tasks' or 'kerr list meetings'.`,
	Run: func(cmd *cobra.Command, args []string) {
		principles := core.GetPrinciples()

		if len(args) == 0 {
			fmt.Println("Things to focus on to do a great job:")
			for i, principle := range principles {
				fmt.Printf("%d. %s\n", i+1, principle.Description)
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}