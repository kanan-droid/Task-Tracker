package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"completed"`
}

var Taskslist []Task
var Pathname string = "database1.json"

// getting data from db
func gettingAllTasks() {
	newdata, err := os.ReadFile()
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
func AddNEwTask(nameOfthetask string) {
	taskId := len(Taskslist) + 1
	task := Task{ID: taskId, Name: nameOfthetask, IsDone: false}
	Taskslist = append(Taskslist, task)
	SavingThetasks()
	fmt.Print("task added ", nameOfthetask)
}

// this one shows Tasks
func showAllTask(mmd bool) {
	if len(Taskslist) == 0 {
		fmt.Print("no task avalible")
		return
	}
	//true uncompele
	for _, task1 := range Taskslist {
		maybyitsdone := "pending"

		if task1.IsDone {
			maybyitsdone = "task done"
			if mmd {
				continue
			}
		}

		fmt.Printf("\n%d. %s - %s - %s \n", task1.ID, task1.Name, maybyitsdone)

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
func MarkingTaskDone(Id string) {
	if len(Taskslist) == 0 {
		fmt.Print("no task available")
		return
	}

	taskID, err := strconv.Atoi(strings.TrimSpace(Id))
	if err != nil {
		fmt.Print("invalid task ID")
		return
	}
	var updatTasks []Task
	for _, task := range Taskslist {
		if task.ID == taskID {
			task.IsDone = true
		}
		updatTasks = append(updatTasks, task)
	}
	Taskslist = updatTasks
	SavingThetasks()
	fmt.Printf("\nTask %d Marked as Done\n", taskID)
}

// main loop with everything
func main() {
	gettingAllTasks()
	fmt.Print("Welcome")
	for {
		fmt.Print("\n1.Add Task \n2.Show Tasks\n3.Mark Task as Done\n4.Delete Task\n5.Exite\n")

		var choices int
		fmt.Scan(&choices)

		switch choices {
		//add task
		case 1:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("\nEnterTaskName :")
			reader.ReadLine()
			name, _ := reader.ReadString('\n')
			AddNEwTask(name)

			//show task
		case 2:

			userinput1 := bufio.NewReader(os.Stdin)
			fmt.Print("1.Show only uncompleted task\n2.show all of the tasks\n==:")
			userinput1.ReadLine()
			text, _ := userinput1.ReadString('\n')
			onlycomp := false
			if text == "1" {
				onlycomp = true
			}
			showAllTask(onlycomp)

			//mark task as done
		case 3:
			showAllTask(false)
			fmt.Print("\nEnter taskID :")
			userinput3 := bufio.NewReader(os.Stdin)
			userinput3.ReadLine()
			text, _ := userinput3.ReadString('\n')
			MarkingTaskDone(text)
			//delete task
		case 4:
			showAllTask(false)
			fmt.Print("\nEnter taskID :")
			userinput2 := bufio.NewReader(os.Stdin)
			userinput2.ReadLine()
			text, _ := userinput2.ReadString('\n')
			Taskremover(text)
		case 5:
			os.Exit(0)
		}
	}
}
