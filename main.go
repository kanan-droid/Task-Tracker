package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsDone      bool      `json:"completed"`
	AdedDate    time.Time `json:"addedtime"`
	Description string    `json:"description"`
	Started     bool      `json:"taskstartedornot"`
}

var Taskslist []Task
var Pathname string = "database1.json"

// getting data from db
func gettingAllTasks() {
	newdata, err := os.ReadFile(Pathname)
	if err != nil {
		panic(err)
	}
	err1 := json.Unmarshal(newdata, &Taskslist)
	if err1 != nil {
		panic(err1)
	}
}

// this is for writing data to the db
func SavingThetasks() {
	data, err := json.MarshalIndent(Taskslist, "", "")

	if err != nil {
		panic(err)
	}
	err1 := os.WriteFile(Pathname, data, 0664)
	if err1 != nil {
		panic(err1)
	}
}

// add task here
func AddNEwTask(nameOfthetask string, desc string) {
	taskId := len(Taskslist) + 1
	task := Task{ID: taskId, Name: nameOfthetask, IsDone: false, AdedDate: time.Now(), Description: desc, Started: false}
	Taskslist = append(Taskslist, task)
	SavingThetasks()
	fmt.Print("task added ", nameOfthetask)
}

// this one shows Tasks
func showAllTask(onlyuncompe bool, onlyInprogress bool) {
	if len(Taskslist) == 0 {
		fmt.Print("no task avalible")
		return
	}
	var counting bool = true
	for _, task1 := range Taskslist {
		maybyitsdone := "pending"
		if onlyInprogress {
			//this if make sure if task has not been started dont show it or if task is finished dont show it show only the one that are in progerssion
			if task1.IsDone {
				continue
			}
		}

		if task1.IsDone {
			//this is to show only the unvompleted tasks and skip the for loop
			if onlyuncompe {
				continue
			}
			maybyitsdone = "task done"
			//for uncomplete task
		}
		if task1.Started {
			maybyitsdone = "in progress"

		}
		//time management
		Timebeforcreation := time.Since(task1.AdedDate)
		var ageago string = ""
		if Timebeforcreation > 24 {
			ageago = fmt.Sprintf("added %f hours ago", Timebeforcreation.Hours()/24)
		}
		if Timebeforcreation > 1 {
			ageago = fmt.Sprintf("added %f minutes ago", Timebeforcreation.Minutes())
		} else {
			ageago = fmt.Sprintf("added %f days ago", Timebeforcreation.Hours())
		}

		fmt.Printf("\n%d.%s - %s \ndescerption:%s -Started:%s \n", task1.ID, task1.Name, maybyitsdone, task1.Description, ageago)
		counting = false
	}
	if counting {
		fmt.Print("\n No uncompleted task available")
	}
}

// remove tasks
func Taskremover(ID string) {
	if len(Taskslist) == 0 {
		fmt.Print("no task available")
		return
	}

	taskID, err := strconv.Atoi(strings.TrimSpace(ID))
	if err != nil {
		fmt.Print("invalid task ID")
		return
	}
	var updatTasks []Task

	for _, task := range Taskslist {
		if task.ID != taskID {
			updatTasks = append(updatTasks, task)
		}
	}

	Taskslist = updatTasks
	// Update task IDs sequentially
	for i := range Taskslist {
		Taskslist[i].ID = i + 1
	}
	SavingThetasks()
	fmt.Printf("\ntask %d removed", taskID)
}

// this func mark tasks as done
func MarkingTaskDone(Id string, markingitdone bool, startingIt bool) {
	if len(Taskslist) == 0 {
		fmt.Print("no task available")
		return
	}
	//input from user to intiger
	taskID, err := strconv.Atoi(strings.TrimSpace(Id))
	if err != nil {
		fmt.Print("invalid task ID")
		return
	}
	var updatTasks []Task

	if markingitdone {
		for _, Taski1 := range Taskslist {
			if Taski1.ID == taskID {
				Taski1.IsDone = true
				Taski1.Started = false
				updatTasks = append(updatTasks, Taski1)
			}
		}
	}
	if startingIt {
		for _, Taski2 := range Taskslist {
			if Taski2.ID == taskID {
				Taski2.Started = true
				Taski2.IsDone = false
				updatTasks = append(updatTasks, Taski2)
			}
		}
	}

	fmt.Printf("\nTask %d has been updated", taskID)
	Taskslist = updatTasks
	SavingThetasks()

}

// main loop with everything
func main() {
	gettingAllTasks()
	fmt.Print("Welcome\n")
	for {
		fmt.Print("\n1. Add Task \n2. Show Tasks\n3.Update task\n4. Delete Task\n5.Exit\n")

		var choices int
		fmt.Scan(&choices)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		switch choices {

		case 1:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("\nEnter Task Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("\nEnter Task Description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			AddNEwTask(name, description)

		case 2:
			userinput1 := bufio.NewReader(os.Stdin)
			fmt.Print("1. Show only uncompleted tasks\n2. Show all tasks\n ")
			text, _ := userinput1.ReadString('\n')
			text = strings.TrimSpace(text)

			uncomp := false
			if text == "1" {
				uncomp = true
			}
			showAllTask(uncomp, false)

			// Update Task
		case 3:
			userinput1 := bufio.NewReader(os.Stdin)
			fmt.Print("\n1. Start a task\n2. Mark a task done\n")
			text, _ := userinput1.ReadString('\n')
			text = strings.TrimSpace(text)

			// Check if the user wants to start a task or mark it done
			if text == "1" {
				fmt.Print("\nWhich task would you like to start?\n")
				showAllTask(true, false)
			} else if text == "2" {
				fmt.Print("\nWhich task would you like to mark as done?\n")
				showAllTask(false, true)
			} else {
				fmt.Println("Invalid option. Please select a valid option.")
				continue
			}

			// Capture the task ID input from the user
			userinput4 := bufio.NewReader(os.Stdin)
			text423, _ := userinput4.ReadString('\n')
			text423 = strings.TrimSpace(text423)

			if text == "1" {
				// Start the task
				MarkingTaskDone(text423, false, true)
			} else if text == "2" {
				// Mark the task as done
				MarkingTaskDone(text423, true, false)
			}

		// Delete Task
		case 4:
			showAllTask(false, false)
			fmt.Print("\nEnter Task ID: ")
			userinput2 := bufio.NewReader(os.Stdin)
			text, _ := userinput2.ReadString('\n')
			text = strings.TrimSpace(text)
			Taskremover(text)

		// Exit
		case 5:
			os.Exit(0)
		}
	}
}
