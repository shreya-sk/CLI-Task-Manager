package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Emoji     string    `json:"emoji"`
	Category  string    `json:"category" default:"personal"`
	Priority  string    `json:"priority" default:"p2"`
}

type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func (l *TaskList) saveTasks(filename string) error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
func (l *TaskList) Loadtasks(filename string) error {
	// Check if file exists
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, initialize empty task list
			l.Tasks = []Task{}
			l.NextID = 1 // If you're using the NextID approach
			return l.saveTasks(filename)
		}
		return fmt.Errorf("failed to access file: %w", err)
	}

	// File exists but is empty
	if fileInfo.Size() == 0 {
		l.Tasks = []Task{}
		l.NextID = 1 // If you're using the NextID approach
		return l.saveTasks(filename)
	}

	// File exists and has content, try to read it
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Try to unmarshal the JSON
	err = json.Unmarshal(b, l)
	if err != nil {
		// Invalid JSON content - this could be corrupted data
		// Best approach: create a backup and initialize a new file
		backupName := filename + ".bak"
		os.Rename(filename, backupName)
		fmt.Printf("Warning: Invalid data in tasks file. A backup was created at %s\n", backupName)

		// Initialize a new empty task list
		l.Tasks = []Task{}
		l.NextID = 1 // If you're using the NextID approach
		return l.saveTasks(filename)
	}

	return nil
}

func (l *TaskList) addTask(title, category, p string) {
	task := Task{
		ID:        l.NextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		Emoji:     "❌",
		Category:  category,
		Priority:  p,
	}
	l.Tasks = append(l.Tasks, task)
	l.NextID++
	color.Green("✅ Task added successfully!")
}

func (l *TaskList) completeTask(id int) error {
	found := false
	for i := range l.Tasks {
		if l.Tasks[i].ID == id {
			l.Tasks[i].Completed = true
			l.Tasks[i].Emoji = "✅"
			found = true
			break
		}
	}

	if !found {
		color.Red("❌ Task with ID %d not found", id)
		return fmt.Errorf("task not found")
	}

	color.Green("✅ Task marked as completed!")
	return nil
}

func (l *TaskList) listTasks() {
	if len(l.Tasks) == 0 {
		color.Yellow("No tasks found.")
		return
	}

	// Define column widths
	idWidth := 4
	titleWidth := 16
	statusWidth := 10
	priorityWidth := 8
	categoryWidth := 10
	dateWidth := 18

	// Create clean divider line
	divider := "+-" + strings.Repeat("-", idWidth) +
		"-+-" + strings.Repeat("-", titleWidth) +
		"-+-" + strings.Repeat("-", statusWidth) +
		"-+-" + strings.Repeat("-", priorityWidth) +
		"-+-" + strings.Repeat("-", categoryWidth) +
		"-+-" + strings.Repeat("-", dateWidth) + "-+"

	// Print header
	fmt.Println(divider)
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		idWidth, "ID",
		titleWidth, "Title",
		statusWidth, "Status",
		priorityWidth, "Priority",
		categoryWidth, "Category",
		dateWidth, "Created At")
	fmt.Println(divider)

	// Print each task
	for _, task := range l.Tasks {
		// Truncate title if too long
		title := task.Title
		if len(title) > titleWidth {
			title = title[:titleWidth-3] + "..."
		}

		// Format status with colors and standard width
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}

		// Make sure priority has a value
		priority := task.Priority
		if priority == "" {
			priority = "N/A"
		}

		// Format date more compactly
		timeFormat := task.CreatedAt.Format("Jan 02 15:04")

		// Print status with color in-line
		fmt.Printf("| %-*d | %-*s | ", idWidth, task.ID, titleWidth, title)

		if task.Completed {
			color.New(color.FgGreen).Printf("%-*s", statusWidth, status)
		} else {
			color.New(color.FgYellow).Printf("%-*s", statusWidth, status)
		}

		fmt.Printf(" | %-*s | %-*s | %-*s |\n",
			priorityWidth, priority,
			categoryWidth, task.Category,
			dateWidth, timeFormat)
	}

	// Print bottom divider
	fmt.Println(divider)
}

func (l *TaskList) deleteTask(id int) error {
	found := false
	for i := range l.Tasks {
		if l.Tasks[i].ID == id {
			l.Tasks = append(l.Tasks[:i], l.Tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		color.Red("❌ Task with ID %d not found", id)
		return fmt.Errorf("task not found")
	}

	color.Yellow("🗑️ Task deleted successfully!")
	return nil
}

func (l *TaskList) searchTask(test string) bool {
	// Implement search functionality
	found := false
	var subs []string
	var id []int
	var categories []string
	for i := range l.Tasks {
		if (strings.Contains(l.Tasks[i].Title, test)) || (strings.Contains(l.Tasks[i].Category, test)) || (strings.Contains(l.Tasks[i].Priority, test)) {
			subs = append(subs, l.Tasks[i].Title)
			id = append(id, l.Tasks[i].ID)
			categories = append(categories, l.Tasks[i].Category)
			found = true
		}
	}
	color.Green("🔎 Found %d tasks with the search term: '%s'", len(subs), test)
	for i := range subs {
		color.White("ID: %d  |   Title: %s |  Category: %s ", id[i], subs[i], categories[i])
	}
	fmt.Println()
	return found

}

func (L *TaskList) stats() {
	color.Cyan("📊 Task List Statistics:")
	color.Cyan("====================================")

	totalT := len(L.Tasks)
	created_weekly := 0
	created_monthly := 0
	finished_weekly := 0
	finished_monthly := 0
	categoryMap := make(map[string]int)

	var completeT int
	for i := range L.Tasks {
		categoryMap[L.Tasks[i].Category]++
		if L.Tasks[i].Completed {
			completeT++
			if L.Tasks[i].CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
				finished_monthly++
			}

			if L.Tasks[i].CreatedAt.After(time.Now().AddDate(0, 0, -7)) {
				finished_weekly++
			}
		}

		if L.Tasks[i].CreatedAt.After(time.Now().AddDate(0, 0, -7)) {
			created_weekly++
		}
		if L.Tasks[i].CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			created_monthly++
		}

	}
	fmt.Printf("%s%s\n",
		color.BlueString("Total Tasks: "),
		color.WhiteString("%d", totalT))
	//remainT := totalT - completeT
	completionRate := float64(completeT) / float64(totalT) * 100

	weeklyRate := 0.0
	if created_weekly > 0 {
		weeklyRate = float64(finished_weekly) / float64(created_weekly) * 100
	}

	monthlyRate := 0.0
	if created_monthly > 0 {
		monthlyRate = float64(finished_monthly) / float64(created_monthly) * 100
	}

	fmt.Printf("%s%s\n",
		color.BlueString("Completion Rate: "),
		color.WhiteString("%.1f%%", completionRate))

	fmt.Printf("%s%s\n",
		color.BlueString("Weekly Completion Rate: "),
		color.WhiteString("%.1f%% (%d/%d)", weeklyRate, finished_weekly, created_weekly))

	fmt.Printf("%s%s\n",
		color.BlueString("Monthly Completion Rate: "),
		color.WhiteString("%.1f%% (%d/%d)", monthlyRate, finished_monthly, created_monthly))
	fmt.Println("~")

	color.Blue("Tasks by priority:")
	for category, count := range categoryMap {
		fmt.Printf("  %s: %s (%s)\n",
			color.WhiteString(category),
			color.WhiteString(fmt.Sprintf("%d", count)),
			color.WhiteString(fmt.Sprintf("%.1f%%", float64(count)/float64(totalT)*100)))
	}
	color.Cyan("====================================")
}
