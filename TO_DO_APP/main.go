package main

import (
	"fmt"
	"net/http"
)

var tasks_predefined = []string{
	"watch go crash course",
	"explore go lang",
	"research about go frameworks",
	"research about go in cloud technologies",
}

func main() {
	fmt.Println("Welcome to Go Task Manager")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/get-task", getTasks)
	http.HandleFunc("/add-task", addTaskHandler)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello User, Welcome to Go Task Manager")
}

func getTasks(writer http.ResponseWriter, request *http.Request) {
	for i, task := range tasks_predefined {
		fmt.Fprintf(writer, "%d. %s\n", i+1, task)
	}
}

func addTaskHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	newTask := request.URL.Query().Get("task")
	if newTask == "" {
		http.Error(writer, "Missing task parameter", http.StatusBadRequest)
		return
	}
	tasks_predefined = append(tasks_predefined, newTask)
	fmt.Fprintf(writer, "Task added: %s\n", newTask)
}
