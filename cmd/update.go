package cmd

import (
	"fmt"
	"project-to-do-list-cli-anas/service"
	"project-to-do-list-cli-anas/utils"

	"github.com/spf13/cobra"
)

var updateID int
var updateStatus string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update status task",
	Run: func(cmd *cobra.Command, args []string) {
		if updateID <= 0 {
			utils.ErrorMessage("index must be greater than 0")
			return
		}

		if updateStatus == "" {
			fmt.Println("--status is required (new, progress, completed)")
			return
		}

		err := service.UpdateStatus(updateID, updateStatus)
		if err != nil {
			utils.ErrorMessage(err.Error())
			return
		}

		fmt.Println("tasks status updated")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateID, "index", "d", 0, "Task ID to update")
	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "", "Task status (new, progress, completed)")
}
