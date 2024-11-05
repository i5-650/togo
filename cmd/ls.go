package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"togo/models"
	"togo/utils"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all task at all levels",
	Long:  `List all task at all levels`,
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

func List() {
	var todo models.Todo
	crossed := color.New(color.CrossedOut)

	if err := utils.LoadYAML("./todo.yml", &todo); err != nil {
		log.Fatalf("Error loading todo.yml: %v", err)
	}

	fmt.Println("What are we doing today ?")

	fmt.Println(utils.Red + "High Priority Tasks:" + utils.Reset)
	RecursivePrint(todo.High, crossed)

	fmt.Println(utils.Yellow + "Medium Priority Tasks:" + utils.Reset)
	RecursivePrint(todo.Medium, crossed)

	fmt.Println(utils.Green + "Low Priority Tasks:" + utils.Reset)
	RecursivePrint(todo.Low, crossed)

}

func RecursivePrint(tasks []*models.Task, crossed *color.Color, level ...int) {
	indent := 0
	if len(level) > 0 {
		indent = level[0]
	}

	for idx, task := range tasks {
		if task.Done {
			fmt.Printf("%s", strings.Repeat("    ", indent))
			crossed.Printf("%d. %s\n", idx+1, task.Title)
		} else {
			fmt.Printf("%s%d. %s\n", strings.Repeat("    ", indent), idx+1, task.Title)
		}
		RecursivePrint(task.Subtasks, crossed, indent+1)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
