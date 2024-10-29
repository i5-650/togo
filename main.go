package main

import (
	"fmt"
	"log"
	"togo/models"
	"togo/utils"
)

func main() {
	var todo models.Todo

	if err := utils.LoadYAML("./todo.yml", &todo); err != nil {
		log.Fatalf("Error loading todo.yml: %v", err)
	}

	fmt.Println("What are we doing today ?")

	fmt.Println("High Priority Tasks:")
	for _, task := range todo.High {
		fmt.Printf(utils.Red+"- %s\n"+utils.Reset, task.Title)
	}

	fmt.Println("Medium Priority Tasks:")
	for _, task := range todo.Medium {
		fmt.Printf(utils.Yellow+"- %s\n"+utils.Reset, task.Title)
	}

	fmt.Println("Low Priority Tasks:")
	for _, task := range todo.Low {
		fmt.Printf(utils.Green+"- %s\n"+utils.Reset, task.Title)
	}
}

