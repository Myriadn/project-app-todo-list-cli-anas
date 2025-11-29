package service

import (
	"errors"
	"project-to-do-list-cli-anas/model"
	"strings"
)

// add new tasks with validation
func AddTasks(title, priority string) error {
	// check title is not untitled
	if title == "" {
		return errors.New("is untitled so we can't")
	}

	// load ts data
	tasks, err := LoadTasks()
	if err != nil {
		return nil
	}

	// duplication validation with loop slice
	// and equalfold for case-insensitive
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return errors.New("this title is already on ts")
		}
	}

	// generate ID (auto increment)
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// create new object task
	newTask := model.Task{
		ID:       newID,
		Title:    title,
		Status:   "new",
		Priority: priority,
	}

	// append to tasks
	tasks = append(tasks, newTask)

	return SaveTasks(tasks)
}

// show all list tasks or filter with search
func ListTasks(filter string) ([]model.Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	// if not exists just return it
	if filter == "" {
		return tasks, nil
	}

	// if it exists do search
	var filteredTasks []model.Task
	for _, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(filter)) {
			filteredTasks = append(filteredTasks, t)
		}
	}

	return filteredTasks, nil
}

// update status tasks from ID
func UpdateStatus(id int, status string) error {
	// input validation
	validStatuses := map[string]bool{
		"new":       true,
		"progress":  true,
		"completed": true,
	}

	if !validStatuses[status] {
		return errors.New("Status not valid it's only: new, progress, completed")
	}

	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// searching tasks with ID then update
	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			found = true
			break
		}
	}

	if !found {
		return errors.New("ts tasks ID not exists")
	}

	return SaveTasks(tasks)
}

// delete tasks from id
func DeleteTasks(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	var newTasks []model.Task
	found := false

	// slice for not deleting the element
	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue // then skip it
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return errors.New("id is not existsted")
	}

	return SaveTasks(newTasks)
}
