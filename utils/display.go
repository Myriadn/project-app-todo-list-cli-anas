package utils

import (
	"fmt"
	"os"
	"project-to-do-list-cli-anas/model"
	"text/tabwriter"
)

// displaying all list tasks to CLI
func DisplayTasks(tasks []model.Task) {
	// if it nothing do ts
	if len(tasks) == 0 {
		fmt.Println("unemployed")
		return
	}

	// tabwriter init
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// header table
	fmt.Fprintln(w, "No\tTask\tStatus\tPriority\t")
	fmt.Fprintln(w, "--\t----\t------\t--------\t")

	// loop data
	for _, t := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", t.ID, t.Title, t.Status, t.Priority)
	}

	// then flush it for printing output to terminal / cli
	w.Flush()
}

// error message
func ErrorMessage(msg string) {
	fmt.Printf("Error: %s\n", msg)
}
