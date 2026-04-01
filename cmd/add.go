package cmd

import (
	"fmt"
	"todo_list/internal/model"
	"todo_list/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add new tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter a valid name for the task")
			return
		}

		tasks, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Failed to load tasks:", err)
			return
		}

		nextID := 1
		for _, t := range tasks {
			if t.ID >= nextID {
				nextID = t.ID + 1
			}
		}

		newTasks := model.Task{
			ID:    nextID,
			Title: args[0],
			Done:  false,
		}

		tasks = append(tasks, newTasks)
		if err := storage.SaveTasks(tasks); err != nil {
			fmt.Println("Failed to save task:", err)
			return
		}

		fmt.Println("Task added")
	},
}