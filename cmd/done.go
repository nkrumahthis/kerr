package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
)

var doneCmd = &cobra.Command{
	Use: "done",
	Short: "Successfully did something",
	Run: func(cmd *cobra.Command, args []string) {
		dbPath := core.EnsureDatabase()
		db := core.OpenDatabase(dbPath)
		defer db.Close()

		if len(args) == 0 {
			listDone(db)
		} else {
			handleDone(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func listDone(db *sql.DB) {
	fmt.Println("List of things you have done:")
}

func handleDone(name string) {
	var principles = core.GetPrinciples()
	var chosenPrinciple core.Principle
	for _, principle := range principles {
		for _, tag := range principle.Tags {
			if tag == name {
				chosenPrinciple = principle
				break
			}
		}
	}
	fmt.Println("You have successfully done something: " + chosenPrinciple.Description)
}

