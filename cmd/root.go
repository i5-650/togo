package cmd

import (
	"log"
	"os"
	"togo/models"
	"togo/utils"

	"github.com/spf13/cobra"
)

var todo models.Todo

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "A CLI & TUI app to handle your todo !",
	Long: `Togo is a terminal based application that aim to be a replacement for the todo apps you use. It relies on YAML file so you will always be able to print the content to switch to another app or do what you want.

    The idea is that you can use the app as a CLI and run some commands or use it like an application as a TUI. To write a few task in one shot or mark them done.
    `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := utils.LoadYAML("./todo.yml", &todo); err != nil {
			log.Fatalf("Error loading todo.yml: %v", err)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.togo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
