package main

import (
	_ "github.com/mattn/go-sqlite3"
	"nkrumahsarpong.com/kerr/cmd"
	"nkrumahsarpong.com/kerr/core"
)

func main() {
	db := core.OpenDatabase(core.EnsureDatabase())
	core.InitializeTables(db)
	cmd.Execute()
}
