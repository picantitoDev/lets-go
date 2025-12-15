package main

import (
	"fmt"
	"os"
)

func printMenu() {
	fmt.Println("===== To-Do List =====")
	fmt.Println("1. View tasks")
	fmt.Println("2. Add a task")
	fmt.Println("3. Remove a task")
	fmt.Println("4. Exit")
	fmt.Print("Choose an option: ")
}

func play() {
	// Initial tasks
	tasks := []string{
		"Buy groceries",
		"Finish Go project",
		"Read a book",
	}

	for {
		printMenu()

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			// View tasks
			if len(tasks) == 0 {
				fmt.Println("No tasks to show.")
			} else {
				fmt.Println("Your Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case 2:
			// Add a task
			fmt.Print("Enter the task description: ")
			var task string
			_, err := fmt.Scan(&task)
			if err != nil {
				fmt.Println("Error reading task.")
				continue
			}
			tasks = append(tasks, task)
			fmt.Println("Task added!")

		case 3:
			// Remove a task
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				continue
			}
			fmt.Println("Which task would you like to remove?")

			// Display current tasks
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

			var taskNum int
			fmt.Print("Enter the number of the task to remove: ")
			_, err := fmt.Scan(&taskNum)
			if err != nil || taskNum < 1 || taskNum > len(tasks) {
				fmt.Println("Invalid input. Please enter a valid task number.")
				continue
			}

			// Remove the task
			tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
			fmt.Println("Task removed!")

		case 4:
			// Exit the program
			fmt.Println("Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Please choose a valid option (1-4).")
		}
	}
}
