package service

import (
	"encoding/json"
	"os"
	"path/filepath"
	"project-to-do-list-cli-anas/model"
)

// json file location
const fileName = "data/todo.json"

// init tasks or load from json then sliced it
func LoadTasks() ([]model.Task, error) {

	// check file if it exist. if it's not, do nothing (not error)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	// Read file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// if list nothing then nothing
	if len(data) == 0 {
		return []model.Task{}, nil
	}

	// decode json to struct
	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// sliced for saving tasks to file json
func SaveTasks(tasks []model.Task) error {
	// encode struct to json
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil
	}

	// make sure folder data is exists before writefile
	// for not error "discovery not found"
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}

	// then write it
	// 0644 code is for read/write
	return os.WriteFile(fileName, data, 0644)
}
