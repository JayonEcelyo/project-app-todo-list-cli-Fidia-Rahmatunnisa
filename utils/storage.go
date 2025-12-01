package utils

import (
	"encoding/json"
	"os"
	"mini-project-2/model"
)

const DataFile = "data/tasks.json"

// Load data
func LoadTasks() ([]model.Task, error) {
	file, err := os.ReadFile(DataFile)
	if err != nil {
		return []model.Task{}, nil // tidak error kalau file belum ada
	}

	var tasks []model.Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// Save data
func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DataFile, data, 0644)
}
