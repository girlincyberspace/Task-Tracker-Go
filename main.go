package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskData struct {
	id          int
	description string
	status      string
	createdAt   string
	updatedAt   string
}

var tasks []TaskData

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		message := strings.TrimSpace(scanner.Text())
		if strings.ToLower(message) == "quit" {
			break
		}
		taskTracker(message)
	}
}

func getCurrentTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func addTask(description string) {
	taskId := len(tasks) + 1
	var newTask = TaskData{
		id:          taskId,
		description: description,
		status:      "todo",
		createdAt:   getCurrentTime(),
	}
	tasks = append(tasks, newTask)
	fmt.Printf("Task added successfully (ID: %d)\n", taskId)
}

func updateTask(newDescription string, taskId int) {
	for i, task := range tasks {
		if taskId == task.id {
			tasks[i].description = newDescription
			tasks[i].updatedAt = getCurrentTime()
			fmt.Printf("You have successfully updated (ID: %d)\n", taskId)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", taskId)
}

func deleteTask(taskId int) {
	var updatedTasks []TaskData
	found := false

	for _, task := range tasks {
		if taskId != task.id {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if found {
		tasks = updatedTasks
		fmt.Printf("Task %d deleted successfully\n", taskId)
	} else {
		fmt.Printf("Task with ID %d not found\n", taskId)
	}
}

func markStatus(taskId int, status string) {
	for i, task := range tasks {
		if taskId == task.id {
			tasks[i].status = status
			fmt.Printf("Task %v marked as %v\n", taskId, status)
			return
		}
	}
	fmt.Printf("Task with ID: %v not found", taskId)
}

func listTasks(status string) {
	if len(tasks) == 0 {
		fmt.Println("No task available")
		return
	}

	found := false
	for _, task := range tasks {
		if status == "" || task.status == status {
			fmt.Printf("%+v\n", task)
			found = true
		}
	}

	if !found {
		fmt.Printf("No tasks with status '%s'\n", status)
	}
}

func taskTracker(message string) {
	commands := strings.Fields(message)

	if len(commands) == 0 {
		fmt.Println("Please enter a valid command")
		return
	}

	action := strings.ToLower(commands[0])
	args := commands[1:]

	switch action {
	case "add":
		if len(args) > 0 {
			addTask(strings.Join(args, " "))
		} else {
			fmt.Printf("Input a task description")
		}
	case "update":
		if len(args) >= 2 {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Task Id")
				return
			}
			updateTask(strings.Join(args[1:], ""), taskId)
		} else {
			fmt.Println("Usage: update <task_id> <task_description>")
		}
	case "delete":
		if len(args) > 0 {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Task Id")
				return
			}
			deleteTask(taskId)
		} else {
			fmt.Println("Usage: delete <task_id>")
		}
	case "mark-in-progress":
		if len(args) > 0 {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Task Id")
				return
			}
			markStatus(taskId, "in-progress")
		} else {
			fmt.Println("Usage: mark-in-progress <task_id>")
		}
	case "mark-done":
		if len(args) > 0 {
			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid Task Id")
				return
			}
			markStatus(taskId, "done")
		} else {
			fmt.Println("Usage: mark-done <task_id>")
		}
	case "list":
		if len(args) == 0 {
			listTasks("")
		} else if args[0] == "done" {
			listTasks("done")
		} else if args[0] == "todo" {
			listTasks("todo")
		} else if args[0] == "in-progress" {
			listTasks("in-progress")
		} else {
			fmt.Println("Invalid list command")
		}
	default:
		fmt.Println("Invalid command")
	}
}
