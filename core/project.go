package core

import (
	"database/sql"
	"log"
)

func FindOrCreateProject(db *sql.DB, name string) int {
	var id int
	err := db.QueryRow("SELECT id FROM projects WHERE name = ?", name).Scan(&id)
	if err == nil {
		return id
	}

	result, err := db.Exec("INSERT INTO projects (name) VALUES (?)", name)
	if err != nil {
		log.Fatalf("Failed to insert project: %v", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Failed to get last insert ID: %v", err)
	}
	id = int(lastInsertId)

	return int(id)
}
