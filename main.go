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
			if !task1.Started || task1.IsDone {
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
		var ageago string
		if Timebeforcreation > 24 {
			ageago = fmt.Sprintf("sinc %f hours ago", Timebeforcreation.Hours())
		} else {
			ageago = fmt.Sprintf("sinc %f days ago", Timebeforcreation.Hours()/24)
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

	for i := range Taskslist {
		if Taskslist[i].ID == taskID {
			if startingIt {
				Taskslist[i].Started = true
			}
			if markingitdone {
				Taskslist[i].IsDone = true
			}
			break
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
			fmt.Print("\n1.start a task \n2.mark task done\n")
			text, _ := userinput1.ReadString('\n')

			if text == "1" {
				showAllTask(true, false)
				fmt.Print("\nwhich task would you like to Start?\n")
			}
			if text == "2" {
				showAllTask(false, true)
				fmt.Print("\nwhich task would you like to Mark Done?\n")
			}
			userinput4 := bufio.NewReader(os.Stdin)
			text423, _ := userinput4.ReadString('\n')
			if text == "1" {
				MarkingTaskDone(text423, false, true)
			} else {
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
