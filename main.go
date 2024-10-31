package main

import (
	"fmt"
	"log"
	"strings"
	"togo/models"
	"togo/utils"
)

func main() {
	var todo models.Todo

	if err := utils.LoadYAML("./todo.yml", &todo); err != nil {
		log.Fatalf("Error loading todo.yml: %v", err)
	}

	fmt.Println("What are we doing today ?")

	fmt.Println(utils.Red + "High Priority Tasks:" + utils.Reset)
	for _, task := range todo.High {
		fmt.Printf("- [%s] %s\n", IsDone(task), task.Title)
		RecursivePrint(task)
	}

	fmt.Println(utils.Yellow + "Medium Priority Tasks:" + utils.Reset)
	for _, task := range todo.Medium {
		fmt.Printf("- [%s] %s\n", IsDone(task), task.Title)
		RecursivePrint(task)
	}

	fmt.Println(utils.Green + "Low Priority Tasks:" + utils.Reset)
	for _, task := range todo.Low {
		fmt.Printf("- [%s] %s\n", IsDone(task), task.Title)
		RecursivePrint(task)
	}
}

func IsDone(t models.Task) string {
	if t.Done {
		return "X"
	}
	return " "
}

func RecursivePrint(task models.Task, level ...int) {
	indent := 1
	if len(level) > 0 {
		indent = level[0]
	}
	for _, sub := range task.Subtasks {
		fmt.Printf("%s- [%s] %s\n", strings.Repeat("    ", indent), IsDone(*sub), sub.Title)
		RecursivePrint(*sub, indent+1)
	}
}
