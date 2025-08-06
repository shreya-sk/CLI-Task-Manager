package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

const filename = "data/tasks.json"

func main() {
	clearScreen()
	taskList := TaskList{}

	// Try to load existing tasks
	err := taskList.Loadtasks(filename)
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		scanner.Scan()
		choice := scanner.Text()
		//fmt.Print("\n")
		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter task category (personal, work, fitness etc.) - press enter to skip: ")
			scanner.Scan()
			category := scanner.Text()
			if category == "" {
				category = "personal"
			}
			fmt.Print("Enter priority (P1/P2/P3) - press enter to skip: ")
			scanner.Scan()
			pri := scanner.Text()
			if pri == "" {
				pri = "P2"
			}
			taskList.addTask(title, category, pri)
			taskList.saveTasks(filename)

		case "2":
			fmt.Println("Tasks:")
			taskList.listTasks()

		case "3":
			fmt.Print("Enter task ID to complete: ")
			var id int
			fmt.Scanln(&id)
			taskList.completeTask(id)
			taskList.saveTasks(filename)
			fmt.Println("")

			// In your main function, modify the delete case:
		case "4":
			fmt.Print("Enter task ID to delete: ")
			var idStr string
			fmt.Scanln(&idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				color.Red("Invalid task ID: %v", err)
				continue
			}

			// Find the task to confirm deletion
			taskFound := false
			for _, task := range taskList.Tasks {
				if task.ID == id {
					taskFound = true
					color.Red("You are about to delete task: %s", task.Title)
					color.Yellow("Are you sure? (y/n): ")

					var confirm string
					fmt.Scanln(&confirm)

					if confirm == "y" || confirm == "Y" {
						err := taskList.deleteTask(id)
						if err == nil {
							err = taskList.saveTasks(filename)
							if err != nil {
								color.Red("Error saving tasks: %v", err)
							}
						}
					} else {
						color.Yellow("Delete operation cancelled")
					}
					break
				}
			}

			if !taskFound {
				color.Red("Task with ID %d not found", id)
			}

		case "5":
			fmt.Print("Enter search term: ")
			scanner.Scan()
			term := scanner.Text()
			found := taskList.searchTask(term)
			if !found {
				color.Yellow("No tasks found matching the search term")
			}

		case "6":
			// Statistics
			taskList.stats()

		case "7":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")

		}
	}
}

func printMenu() {
	color.Cyan("-------------------------")
	color.Cyan("📋 Task Manager")
	color.Cyan("-------------------------")
	color.White("1. Add Task")
	color.White("2. Show Tasks")
	color.White("3. Complete Task")
	color.White("4. Delete Task")
	color.White("5. Search Task")
	color.White("6. Completion Stats!")
	color.White("7. Exit")
	color.Cyan("-------------------------")
	color.Yellow("Enter your choice: ")
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
