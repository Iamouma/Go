package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"io/ioutil"
)

const todoFile = "task.txt"

// Function to read tasks from the file
func readTasks() ([]string, error) {
	// Check if the file exists, if not create it
	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		return []string{}, nil
	}

	content, err := ioutil.ReadFile(todoFile)
	if err != nil {
		return nil, err
	}

	tasks := strings.Split(string(content), "\n")
	return tasks[:len(tasks)-1], nil //Remove the last empty line
}

// Function to save tasks to  the file
func saveTasks(tasks []string) error {
	return ioutil.WriteFile(todoFile, []byte(strings.Join(tasks, "\n") +"\n"), 0644)
}

// Function to display all tasks
func viewTasks() {
	tasks, err := readTasks()
	if err != nil {
		log.Fatalf("Error reding tasks: %s\n", err)
	}

	fmt.Println("\nYour To-Do List:")
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
	} else {
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
	fmt.Println()
}

// Function to add a task
func addTask(task string) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatalf("Error reading tasks: %s\n", err)
	}

	tasks = append(tasks, task)
	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Error saving tasks: %s\n", err)
	}

	fmt.Printf("Task added: %s\n", task)
}

// Function to delete a task
func deleteTask(index int) {
	tasks, err := readTasks()
	if err != nil {
		log.Fatalf("Error reading tasks: %s\n", err)
	}

	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	err = saveTasks(tasks)
	if err != nil {
		log.Fatalf("Error saving tasks: %s\n", err)
	}

	fmt.Println("Task deleted.")
}

// CLI Menu
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("To-Do List CLI")
        fmt.Println("1. View tasks")
        fmt.Println("2. Add a task")
        fmt.Println("3. Delete a task")
        fmt.Println("4. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            viewTasks()
        case "2":
            fmt.Print("Enter the task: ")
            scanner.Scan()
            task := scanner.Text()
            addTask(task)
        case "3":
            viewTasks()
            fmt.Print("Enter the task number to delete: ")
            scanner.Scan()
            var taskNum int
            fmt.Sscanf(scanner.Text(), "%d", &taskNum)
            deleteTask(taskNum)
        case "4":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please enter a valid option.")
        }
    }
}