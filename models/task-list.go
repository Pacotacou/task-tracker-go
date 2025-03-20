package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
	DataFile         = "tasks.json"
)

func NewTaskList() (TaskList, error) {
	var taskList TaskList

	if _, err := os.Stat(DataFile); os.IsNotExist(err) {
		return TaskList{Tasks: []Task{}}, nil
	}

	data, err := os.ReadFile(DataFile)
	if err != nil {
		return taskList, fmt.Errorf("error %v", err)
	}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return taskList, fmt.Errorf("error: %v", err)
	}

	return taskList, nil
}

func (taskList *TaskList) SaveTasks() error {
	json, err := json.MarshalIndent(taskList, "", "	")
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	err = os.WriteFile(DataFile, json, 0644)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func nextID(taskList *TaskList) int {
	if len(taskList.Tasks) == 0 {
		return 1
	}
	maxID := 1
	for _, task := range taskList.Tasks {
		maxID = task.ID
	}
	return maxID + 1
}

func (taskList *TaskList) AddTask(description string) error {
	newTask := Task{
		ID:          nextID(taskList),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	taskList.Tasks = append(taskList.Tasks, newTask)

	err := taskList.SaveTasks()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Task saved successfully with ID: %v\n", newTask.ID)
	return nil
}

func (taskList *TaskList) UpdateTask(description string, id int) error {
	found := false
	for index, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[index].Description = description
			taskList.Tasks[index].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("error: %v", errors.New("task id not found"))
	}
	err := taskList.SaveTasks()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("task updated successfully\n")
	return nil
}

func (taskList *TaskList) RemoveTask(id int) error {
	found := false
	for index, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("id not found")
	}

	if err := taskList.SaveTasks(); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Task deleted successfully")

	return nil
}

func (taskList *TaskList) ListTasks(filter string) error {
	tasksPrinted := false

	//print header
	fmt.Printf("%-5s %-15s %-30s %-20s\n", "ID", "Status", "Description", "Updated At")
	fmt.Println(strings.Repeat("-", 100))

	switch filter {
	case "":
		for _, task := range taskList.Tasks {
			tasksPrinted = true
			fmt.Printf("%-5d %-15s %-30s %-20s\n", task.ID, task.Status, task.Description, task.UpdatedAt)
		}
	default:
		for _, task := range taskList.Tasks {
			if task.Status == filter {
				tasksPrinted = true
				fmt.Printf("%-5d %-15s %-30s %-20s\n", task.ID, task.Status, task.Description, task.UpdatedAt)
			}
		}
	}

	if !tasksPrinted {
		fmt.Println("No tasks to print!")
	}

	return nil
}

func (taskList *TaskList) MarkTask(status string, id int) error {
	found := false
	for index, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[index].Status = status
			found = true
		}
	}
	if !found {
		return fmt.Errorf("error, id not found")
	}
	if err := taskList.SaveTasks(); err != nil {
		return fmt.Errorf("error saving tasks, %v", err)
	}
	fmt.Println("Task marked", status, "successfully")

	return nil
}
