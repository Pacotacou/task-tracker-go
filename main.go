package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/models"
)

func main() {
	argsCount := len(os.Args)
	command := os.Args[1]
	taskList, err := models.NewTaskList()
	if err != nil {
		println("Error creating New Task List")
	}
	switch command {
	case "add":
		add(argsCount, &taskList)
	case "update":
		update(argsCount, &taskList)
	case "delete":
		delete(argsCount, &taskList)
	case "list":
		list(argsCount, &taskList)
	case "mark-done":
		markDone(argsCount, &taskList)
	case "mark-in-progress":
		markInProgress(argsCount, &taskList)
	}

}

func markInProgress(argsCount int, taskList *models.TaskList) {
	if argsCount != 3 {
		fmt.Println("Error 1 argument expeted: id")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: id should be a number")
	}
	if err := taskList.MarkTask(models.StatusInProgress, id); err != nil {
		fmt.Println("Error:", err)
	}
}

func markDone(argsCount int, taskList *models.TaskList) {
	if argsCount != 3 {
		fmt.Println("Error 1 argument expeted: id")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: id should be a number")
	}
	if err := taskList.MarkTask(models.StatusDone, id); err != nil {
		fmt.Println("Error:", err)
	}
}

func list(argsCount int, taskList *models.TaskList) {
	if argsCount > 3 {
		fmt.Println("Error: too many arguments")
	} else if argsCount == 3 {
		if err := taskList.ListTasks(os.Args[2]); err != nil {
			fmt.Printf("Error: %v", err)
		}
	} else {
		if err := taskList.ListTasks(""); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}

func delete(argsCount int, taskList *models.TaskList) {
	if argsCount != 3 {
		fmt.Println("Error: 1 argument expected: 'id'")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: ID should be a number")
	}
	err = taskList.RemoveTask(id)
	if err != nil {
		fmt.Printf("Error, %v\n", err)
	}
}

func add(argsCount int, taskList *models.TaskList) {
	if argsCount != 3 {
		fmt.Println("Error: 1 arguments expected: 'description'")
	}
	err := taskList.AddTask(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func update(argsCount int, taskList *models.TaskList) {
	if argsCount != 4 {
		fmt.Println("Error: 2 arguments expected: 'id', 'description'")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: ID should be a number")
	}
	if err = taskList.UpdateTask(os.Args[3], id); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
