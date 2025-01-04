package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
	"nkrumahsarpong.com/kerr/db"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project in current folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing project...")
		projectName := filepath.Base(core.GetCurrentDir())
		db := db.GetDB()
		defer db.Close()
		projectID := core.FindOrCreateProject(db, projectName)
		fmt.Printf("Initialized project '%s' (ID: %d)\n", projectName, projectID)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
