package cmd

import (
	"fmt"
	"project-to-do-list-cli-anas/service"
	"project-to-do-list-cli-anas/utils"

	"github.com/spf13/cobra"
)

// ts for take all flag filter
var filteredKeyword string

// listCmd represent for "list"
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all list task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := service.ListTasks(filteredKeyword)
		if err != nil {
			utils.ErrorMessage(fmt.Sprintf("failed to load data: %v", err))
			return
		}

		utils.DisplayTasks(tasks)
	},
}

func init() {
	// regist command list to root
	rootCmd.AddCommand(listCmd)

	// add flag search according to keyword
	// like go run main.go list --filter "belajar"
	listCmd.Flags().StringVarP(&filteredKeyword, "filter", "f", "", "Search tasks from keyword")
}
