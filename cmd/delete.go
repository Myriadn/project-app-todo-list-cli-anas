package cmd

import (
	"fmt"
	"project-to-do-list-cli-anas/service"
	"project-to-do-list-cli-anas/utils"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID <= 0 {
			utils.ErrorMessage("ID must greater than 0")
			return
		}

		err := service.DeleteTasks(deleteID)
		if err != nil {
			utils.ErrorMessage(err.Error())
			return
		}

		fmt.Printf("Task ID %d deleted\n", deleteID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "Task ID will be deleted")
}
