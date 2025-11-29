package cmd

import (
	"fmt"
	"project-to-do-list-cli-anas/service"
	"project-to-do-list-cli-anas/utils"

	"github.com/spf13/cobra"
)

var title string
var priority string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if title == "" {
			utils.ErrorMessage("please don't be untitled")
			return
		}

		err := service.AddTasks(title, priority)
		if err != nil {
			utils.ErrorMessage(err.Error())
			return
		}

		fmt.Println("New tasks added!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&title, "title", "t", "", "Task title (required)")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "priority (low, medium, high)")

	// ts must to be added
	addCmd.MarkFlagRequired("title")
}
