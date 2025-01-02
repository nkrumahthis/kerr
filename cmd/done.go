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
			listDone()
		} else {
			handleDone(db, args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func listDone() {
	principles := core.GetPrinciples()
	for _, principle := range principles {
		fmt.Println("Principle: " + principle.Description)
		for _, tag := range principle.Tags {
			fmt.Println(" - Tag: " + tag)
		}
	}
}

func handleDone(db *sql.DB, name string) {
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

	if chosenPrinciple.Description == "" {
		fmt.Println("Principle not found. Would you like to create a new one? (yes/no)")
		var response string
		fmt.Scanln(&response)
		if response == "yes" {
			fmt.Println("Enter the description of the new principle:")
			var description string
			fmt.Scanln(&description)
			newPrinciple := core.Principle{
				Description: description,
				Tags:        []string{name},
			}
			principles = append(principles, newPrinciple)
			core.SavePrinciples(principles)
			chosenPrinciple = newPrinciple
		} else {
			fmt.Println("No new principle created.")
			return
		}
	}

	fmt.Println("You have successfully done something: " + chosenPrinciple.Description)
}

