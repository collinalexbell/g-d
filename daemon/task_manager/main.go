package main

// This is a task manager that stores tasks as a tree
// Tasks can be added, deleted, selected, completed

import (
	"fmt"
	"os"
	"bufio"
)

type Task struct {
	Description string
	// A Task can store subtasks	
	Subtasks []*Task
	// A Task can be completed
	Completed bool

	// Pointer to supertask
	Parent *Task
}

func (t *Task) AddSubtask(subtask *Task) {
	t.Subtasks = append(t.Subtasks, subtask)
	subtask.Parent = t
}

func (t *Task) HasCompleteSubtasks() bool {
	for _, subtask := range t.Subtasks {
		if !subtask.Completed {
			return false
		}
	}
	return true
}

func (t *Task) IsComplete() bool {
	return t.Completed && t.HasCompleteSubtasks()
}

func (t *Task) SetCompleted() {
	// Can only SetCompleted if there are no subtasks remaining to complete
	if t.HasCompleteSubtasks() {
		t.Completed = true
	}
}

func (t *Task) SetIncomplete() {
	t.Completed = false
}

func (t *Task) DeleteSubtask(subtask *Task) {
	for i, subtask := range t.Subtasks {
		if subtask == subtask {
			t.Subtasks = append(t.Subtasks[:i], t.Subtasks[i+1:]...)
		}
	}
}

type TaskManager struct {
	Root *Task
	// A TaskManager can be selected
	SelectedTask *Task
}

func (tm *TaskManager) SelectTask(task *Task) {
	tm.SelectedTask = task
}

// Complete selected task, then complete its subtasks
func (tm *TaskManager) CompleteTask() {
	if tm.SelectedTask == nil {
		fmt.Println("No task selected")
	} else {
		if !tm.SelectedTask.HasCompleteSubtasks() {
			fmt.Println("There are uncompleted subtasks")
			fmt.Println("Would you like to complete all the subtasks?")
			fmt.Println("Enter 'y' for yes, 'n' for no")
			var response string
			fmt.Scanln(&response)
			if response == "y" {
				for _, subtask := range tm.SelectedTask.Subtasks {
					subtask.SetCompleted()
				}
			} else {
				return
			}
		}
		tm.SelectedTask.SetCompleted()
		tm.SelectTask(tm.SelectedTask.Parent)
	}
}

// Incomplete selected task, then incomplete its subtasks
func (tm *TaskManager) IncompleteTask() {
	if tm.SelectedTask == nil {
		fmt.Println("No task selected")
	} else {
		tm.SelectedTask.SetIncomplete()
		for _, subtask := range tm.SelectedTask.Subtasks {
			subtask.SetIncomplete()
		}
	}
}

// Delete selected task, then delete its subtasks
func (tm *TaskManager) DeleteTask() {
	if tm.SelectedTask == nil {
		fmt.Println("No task selected")
	} else {
		tm.SelectedTask.DeleteSubtask(tm.SelectedTask)
	}
}

// AddTasks adds one to the selected
func (tm *TaskManager) AddTask(description string) *Task {
	// If there is no selected task, add a new task to the root
	if tm.SelectedTask == nil {
		tm.Root = &Task{Description: description}
		tm.Root.Parent = tm.Root
		tm.SelectedTask = tm.Root
		return tm.Root
	} else {
		// If there is a selected task, add a new task to the selected task
		newTask := &Task{Description: description}
		tm.SelectedTask.AddSubtask(newTask)
		return newTask
	}
}

// This function displays all the possible commands and prompts use to issue a command
func DisplayCommands() {

	fmt.Println("select")
	fmt.Println("deselect")
	fmt.Println("complete")
	fmt.Println("incomplete")
	fmt.Println("delete")
	fmt.Println("add")
	fmt.Println("exit")
}

// This function prompts the user for a command and returns the command
func GetCommand() string {
	fmt.Print("Enter a command: ")
	var command string
	fmt.Scanln(&command)
	return command
}



type TaskManagerCLI interface {
	Select()
	Complete()
	Incomplete()
	Delete()
	Add()
	Display()
}

// Display selected task's description, then display its subtask descriptions indented
func (tm *TaskManager) Display() {
	if tm.SelectedTask == nil {
		fmt.Println("No task selected")
	} else {
		// Complete tasks display as green
		// Incomplete tasks display as red
		if tm.SelectedTask.IsComplete() {
			fmt.Printf("\033[32m")
		} else {
			fmt.Printf("\033[31m")
		}
		fmt.Printf("%s\n", tm.SelectedTask.Description)
		fmt.Printf("\033[0m")
		for i, subtask := range tm.SelectedTask.Subtasks {
			if subtask.IsComplete() {
				fmt.Printf("\033[32m")
			} else {
				fmt.Printf("\033[31m")
			}
			fmt.Printf("%d. %s\n", i, subtask.Description)
			fmt.Printf("\033[0m")
		}
	}
}

func (tm *TaskManager) Select() {
	// Prompt user for task index
	fmt.Print("Enter task index: ")
	var index int
	fmt.Scanln(&index)
	// Select the task
	tm.SelectTask(tm.SelectedTask.Subtasks[index])
}

func (tm *TaskManager) Complete() {
	// Complete the selected task
	tm.CompleteTask()
}

func (tm *TaskManager) Incomplete() {
	// Incomplete the selected task
	tm.IncompleteTask()
}

func (tm *TaskManager) Delete() {
	// Delete the selected task
	tm.DeleteTask()
}

func (tm *TaskManager) Add() {
	// Prompt user for task description
	fmt.Print("Enter task description: ")
	var description string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		description = scanner.Text()

	}
	if len(description) > 0 {
		tm.AddTask(description)
	}
}

func (tm *TaskManager) Deselect() {
	tm.SelectedTask = tm.SelectedTask.Parent
}

func main() {
	// Create a new task manager
	tm := &TaskManager{}

	// Add a task to the root
	tm.AddTask("Clean the apartment")

	// Add a task to the root's subtasks
	tm.AddTask("Wash the dishes")
	tm.AddTask("Vacuum")
	tm.AddTask("Mop the floor")



	// Prompt user for a command
	for {
		fmt.Println("\n\n")
		tm.Display()
		fmt.Println("\n\n")
		DisplayCommands()
		command := GetCommand()

		switch command {
		case "select":
			tm.Select()
		case "deselect":
			tm.Deselect()
		case "complete":
			tm.Complete()
		case "incomplete":
			tm.Incomplete()
		case "delete":
			tm.Delete()
		case "add":
			tm.Add()
		case "exit":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
