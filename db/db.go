package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const appDataFolder = ".kerr"
const dbFileName = "kerr.db"

func openDatabase(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	return db
}

func ensureDatabase() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	appDataPath := filepath.Join(homeDir, appDataFolder)
	if _, err := os.Stat(appDataPath); os.IsNotExist(err) {
		if err := os.MkdirAll(appDataPath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create app data folder: %v", err)
		}
	}

	dbPath := filepath.Join(appDataPath, dbFileName)
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if _, err := os.Create(dbPath); err != nil {
			log.Fatalf("Failed to create database file: %v", err)
		}
	}

	return dbPath
}

func initializeTables(db *sql.DB) {
	projectTable := `
		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL
		);
	`
	taskTable := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			deadline TEXT NOT NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id)
		);
	`
	stepTable := `
		CREATE TABLE IF NOT EXISTS steps (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task_id INTEGER NOT NULL,
			description TEXT NOT NULL,
			FOREIGN KEY (task_id) REFERENCES tasks(id)
		);
	`

	logsTable := `
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			action_code TEXT NOT NULL,
			score INTEGER NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	actionTable := `
		CREATE TABLE IF NOT EXISTS actions (
			code TEXT PRIMARY KEY,
			description TEXT NOT NULL,
			score INTEGER NOT NULL,
			process TEXT NOT NULL
		);
	`

	totalsTable := `
		CREATE TABLE IF NOT EXISTS totals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			total INTEGER NOT NULL
		);
	`

	_, err := db.Exec(projectTable)
	if err != nil {
		log.Fatalf("Failed to create projects table: %v", err)
	}

	_, err = db.Exec(taskTable)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}

	_, err = db.Exec(stepTable)
	if err != nil {
		log.Fatalf("Failed to create steps table: %v", err)
	}

	_, err = db.Exec(logsTable)
	if err != nil {
		log.Fatalf("Failed to create score table: %v", err)
	}

	_, err = db.Exec(totalsTable)
	if err != nil {
		log.Fatalf("Failed to create totals table: %v", err)
	}

	_, err = db.Exec(actionTable)
	if err != nil {
		log.Fatalf("Failed to create actions table: %v", err)
	}
}

func seedDatabase(db *sql.DB) error {
	file, err := os.Open("./core/actions.json")
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %v", err)
	}

	// Parse the JSON data
	var actions []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		Score       int    `json:"score"`
		Process     string `json:"process"`
	}
	if err := json.Unmarshal(byteValue, &actions); err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	// Insert data into the database
	for _, action := range actions {
		query := `INSERT INTO actions (code, description, score, process) VALUES (?, ?, ?, ?)`
		_, err := db.Exec(query, action.Code, action.Description, action.Score, action.Process)
		if err != nil {
			return fmt.Errorf("failed to insert action into database: %v", err)
		}
	}

	return nil
}

var instance *sql.DB

func init() {
	dbPath := ensureDatabase()
	instance = openDatabase(dbPath)
	initializeTables(instance)
	seedDatabase(instance)
}

func GetDB() *sql.DB {
	fmt.Println("Database initialized")
	return instance
}
