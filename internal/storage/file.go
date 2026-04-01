package storage

import (
	"encoding/json"
	"os"
	"todo_list/internal/model"
)

const fileName = "task.json"

func LoadTasks() ([]model.Task, error) {
	var tasks []model.Task
	
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}