package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type theTasks struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsItDone bool   `json:"Completed"`
}

var Taskslist []theTasks
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
func JasonWriting() {
	data, err := json.MarshalIndent(Taskslist, "", "")

	if err != nil {
		panic(err)
	}
	fmt.Print(data)
	err1 := os.WriteFile(Pathname, data, 0664)
	if err1 != nil {
		panic(err1)
	}
}

func AddNEwTask(nameOfthetask string) {

	taskId := len(Taskslist) + 1
	task := theTasks{Id: taskId, Name: nameOfthetask, IsItDone: false}
	Taskslist = append(Taskslist, task)
	JasonWriting()
	fmt.Print("task added", nameOfthetask)
}

func showAllTask(mmd bool) {
	if len(Taskslist) == 0 {
		fmt.Print("no task avalible")
		return
	}
	//true uncompele
	for _, task1 := range Taskslist {
		maybyitsdone := "pending"

		if task1.IsItDone {
			maybyitsdone = "task done"
			if mmd {
				continue
			}
		}

		fmt.Printf("\n%d. %s - %s\n", task1.Id, task1.Name, maybyitsdone)

	}

}

func main() {
	gettingAllTasks()
	fmt.Print("Welcome")
	for {
		fmt.Print("\n1.Add Task \n2.Show All Tasks\n3.Mark Task as Done\n4.Delete Task\n5.Exite\n")

		var choices int
		fmt.Scan(&choices)

		switch choices {
		case 1:
			userinput := bufio.NewReader(os.Stdin)
			fmt.Print("EnterTaskName :")
			userinput.ReadLine()
			text, _ := userinput.ReadString('\n')
			AddNEwTask(text)
		case 2:
			userinput1 := bufio.NewReader(os.Stdin)
			fmt.Print("1.Show only uncompleted task\n2.show all of them")
			userinput1.ReadLine()
			text, _ := userinput1.ReadString('\n')
			tuly := false
			if text == "1" {
				tuly = true
			}
			showAllTask(tuly)
		case 3:
		case 4:
		case 5:
			os.Exit(0)
		}
	}
}
