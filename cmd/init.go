package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize a new project in current folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing project...")
		projectName := filepath.Base(core.GetCurrentDir())
		dbPath := core.EnsureDatabase()
		db := core.OpenDatabase(dbPath)
		defer db.Close()
		core.InitializeTables(db)
		projectID := core.FindOrCreateProject(db, projectName)
		fmt.Printf("Initialized project '%s' (ID: %d)\n", projectName, projectID)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}