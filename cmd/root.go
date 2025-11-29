package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// represent the basic command when called without a subcommand
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "To-Do List App",
	Long:  `Tasks management for add, check, and settings daily tasks`,
}

// adding all child commands to root and set flags
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init load
func init() {

}
