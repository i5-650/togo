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
		fmt.Println("-", task.Title)
	}

	fmt.Println("Medium Priority Tasks:")
	for _, task := range todo.Medium {
		fmt.Println("-", task.Title)
	}

	fmt.Println("Low Priority Tasks:")
	for _, task := range todo.Low {
		fmt.Println("-", task.Title)
	}
}

