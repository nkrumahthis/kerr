package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"

	"github.com/spf13/cobra"
	"nkrumahsarpong.com/kerr/core"
	"nkrumahsarpong.com/kerr/db"
)

type task struct {
	ID       int
	Name     string
	Deadline string
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks within the current project",
	Run: func(cmd *cobra.Command, args []string) {
		db := db.GetDB()
		defer db.Close()

		if len(args) == 0 {
			listTasks(db)
		} else {
			handleNewTask(db, args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}

func listTasks(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, deadline FROM tasks")
	if err != nil {
		log.Fatalf("Failed to list tasks: %v", err)
	}
	defer rows.Close()

	tasks := []task{}
	for rows.Next() {
		t := task{}
		if err := rows.Scan(&t.ID, &t.Name, &t.Deadline); err != nil {
			log.Fatalf("Failed to scan task: %v", err)
		}
		tasks = append(tasks, t)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Tasks:")
	for _, t := range tasks {
		fmt.Printf("%d. %s (Deadline: %s)\n", t.ID, t.Name, t.Deadline)
	}
}

func insertTask(db *sql.DB, projectID int, taskName, deadline string, steps []string) {
	res, err := db.Exec("INSERT INTO tasks (project_id, name, deadline) VALUES (?, ?, ?)", projectID, taskName, deadline)
	if err != nil {
		log.Fatalf("Failed to insert task: %v", err)
	}
	taskID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("Failed to retrieve last insert id: %v", err)
		return
	}

	for _, step := range steps {
		_, err := db.Exec("INSERT INTO steps (task_id, description) VALUES (?, ?)", taskID, step)
		if err != nil {
			log.Fatalf("Failed to insert step: %v", err)
		}
	}
}

func handleNewTask(db *sql.DB, taskName string) {
	fmt.Printf("Creating task '%s'...\n", taskName)
	projectID := core.FindOrCreateProject(db, core.GetCurrentDir())
	if projectID == 0 {
		log.Fatalf("Failed to find or create project")
		return
	}

	fmt.Print("Deadline (e.g., 'today', 'in 2 hours', 'next Friday'): ")
	var deadlineInput string
	fmt.Scanln(&deadlineInput)
	deadline, err := parseNaturalLanguageDeadline(deadlineInput)

	if err != nil {
		log.Fatalf("Failed to parse deadline: %v", err)
		return
	}

	fmt.Println("Enter task steps (one per line, finish with an empty line):")
	steps := []string{}
	for {
		var step string
		fmt.Scanln(&step)
		if step == "" {
			break
		}
		steps = append(steps, step)
	}

	insertTask(db, projectID, taskName, deadline.Format(time.RFC3339), steps)
}

func parseNaturalLanguageDeadline(input string) (time.Time, error) {
	w := when.New(nil)
	w.Add(en.All...)

	now := time.Now()
	result, err := w.Parse(input, now)
	if err != nil || result == nil {
		return time.Time{}, fmt.Errorf("unable to parse deadline")
	}

	return result.Time, nil
}
