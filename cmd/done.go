package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"togo/models"
	"togo/utils"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Long: `Mark a task as done. 

    To mark a task as done, you nedd to provide the task category (high, medium, low) followed by the task id. 
    If the task you want to mark as done is a sub task of a sub sub task, you need to specify the parent task and the sub task separated by a dot.
    Example: togo done high 1.2 # this will mark the sub task 2 of the task 1 of the high group done
             togo done low 1.1.2.3 # this will work for all the sub tasks`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("You need to provide a category and a task id")
			return
		}

		categ := args[0]

		if categ == "high" {
			CompleteTask(todo.High, args[1])

		} else if categ == "medium" {
			CompleteTask(todo.Medium, args[1])

		} else if categ == "low" {
			CompleteTask(todo.Low, args[1])

		} else {
			fmt.Println("You monkey")
			return
		}
		saveTasksToFile(todo)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func CompleteTask(tasks []*models.Task, ids string) {
	idStack, err := utils.ParseIds(ids)
	if err != nil {
		fmt.Println("Error parsing ids")
		return
	}

	if len(idStack) == 0 {
		fmt.Println("You need to provide a task id")
		return
	}

	for len(idStack) > 0 {
		currentId, err := idStack.Pop()

		if err != nil || currentId >= len(tasks) {
			fmt.Println("Task not found")
			return
		}

		if len(idStack) == 0 {
			tasks[currentId].Done = !tasks[currentId].Done
			fmt.Printf("%s - is done", tasks[currentId].Title)
		} else {
			tasks = tasks[currentId].Subtasks
		}
	}
}

func saveTasksToFile(todo models.Todo) error {
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	return fmt.Errorf("could not find user home directory: %v", err)
	// }

	filePath := filepath.Join("./", "todo.yml")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file at %s: %v", filePath, err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(todo); err != nil {
		return fmt.Errorf("could not encode tasks to YAML: %v", err)
	}

	return nil
}
