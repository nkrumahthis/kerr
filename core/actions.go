package core

import (
	"fmt"
	"log"
	"strings"
)

type Action struct {
	Code        string
	Description string
	Score       int
	Process     string
}

func LogAchievements(codes []string, isSuccess bool) {
	db := OpenDatabase(EnsureDatabase())
	defer db.Close()

	query := "SELECT code, score FROM actions WHERE code IN (" + strings.Repeat("?,", len(codes)-1) + "?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Failed to prepare query: %v", err)
	}
	rows, err := stmt.Query(toInterfaceSlice(codes)...)
	if err != nil {
		log.Fatalf("Failed to query actions: %v", err)
	}
	defer rows.Close()

	totalScore := 0
	for rows.Next() {
		var code string
		var score int
		if err := rows.Scan(&code, &score); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		// Adjust score based on success or failure
		adjustedScore := score
		if !isSuccess {
			adjustedScore = -score
		}

		// log the action
		logQuery := "INSERT INTO logs (action_code, score) VALUES (?, ?)"
		if _, err := db.Exec(logQuery, code, adjustedScore); err != nil {
			log.Printf("Failed to log action %s: %v", code, err)
			continue
		}
		status := "gained"
		if !isSuccess {
			status = "lost"
		}
		fmt.Printf("Logged action: %s (%s %d XP)\n", code, status, adjustedScore)
		totalScore += adjustedScore
	}

	// Update the total score
	updateTotal := "UPDATE totals SET total = total + ? WHERE name = 'xp'"
	if _, err := db.Exec(updateTotal, totalScore); err != nil {
		log.Fatalf("Failed to update total score: %v", err)
	}
	fmt.Printf("Total XP: %d\n", totalScore)
}

func ListActions() {
	db := OpenDatabase(EnsureDatabase())
	defer db.Close()

	rows, err := db.Query("SELECT code, description, score FROM actions")
	if err != nil {
		log.Fatalf("Failed to list actions: %v", err)
	}
	defer rows.Close()

	actions := []Action{}
	for rows.Next() {
		a := Action{}
		if err := rows.Scan(&a.Code, &a.Description, &a.Score); err != nil {
			log.Fatalf("Failed to scan action: %v", err)
		}
		actions = append(actions, a)
	}

	if len(actions) == 0 {
		fmt.Println("No actions found")
		return
	}

	fmt.Println("Actions:")
	for _, a := range actions {
		fmt.Printf(" - %s. %s (%d XP)\n", a.Code, a.Description, a.Score)
	}

}

func toInterfaceSlice(strs []string) []interface{} {
	iface := make([]interface{}, len(strs))
	for i, v := range strs {
		iface[i] = v
	}
	fmt.Println(iface)
	return iface
}
