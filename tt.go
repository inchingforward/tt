package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"
)

type task struct {
	Name    string `json:"name"`
	Started int64  `json:"started"`
	Ended   int64  `json:"ended"`
}

type taskCollection struct {
	Collection []task
}

func getTTFileName() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", usr.HomeDir, "tt.json")
}

func loadTasks() []task {
	file, e := ioutil.ReadFile(getTTFileName())
	if e != nil {
		return []task{}
	}

	tasks := make([]task, 0)
	json.Unmarshal(file, &tasks)

	return tasks
}

func saveTasks(tasks []task) {
	tasksJSON, _ := json.Marshal(tasks)
	ioutil.WriteFile(getTTFileName(), tasksJSON, 0644)
}

func showTasks(tasks []task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}

	format := "%-40s %-16s %-16s %-8s\n"
	fmt.Printf(format, "Name", "Started", "Ended", "Time")

	count := 1
	for _, task := range tasks {
		fmt.Printf(format, task.Name, "", "", "")
		count++
	}
}

func stopTask(tasks []task) {
	for idx, task := range tasks {
		if task.Ended == -1 {
			tasks[idx].Ended = time.Now().Unix()
			fmt.Printf("Task \"%s\" ended\n", task.Name)
		}
	}

	saveTasks(tasks)
}

func startTask(tasks []task, name string) {
	stopTask(tasks)

	fmt.Printf("Starting task \"%s\"\n", name)

	task := task{name, time.Now().Unix(), -1}
	tasks = append(tasks, task)

	saveTasks(tasks)
}

func printUsage() {
	fmt.Println("Usage: tt [start|stop]")
}

func main() {
	tasks := loadTasks()

	if len(os.Args) == 1 {
		showTasks(tasks)
	} else if len(os.Args) == 2 && os.Args[1] == "stop" {
		stopTask(tasks)
	} else if len(os.Args) == 3 && os.Args[1] == "start" {
		startTask(tasks, os.Args[2])
	} else {
		printUsage()
	}
}
