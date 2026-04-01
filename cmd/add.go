package cmd

import (
	"todo_list/internal/model"
	"todo_list/internal/storage"
	"github.com/spf13/cobra"
	"fmt"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add new tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter a valid name for the task")
			return
		}

		tasks, _ := storage.LoadTasks()

		newTasks := model.Task {
			ID: len(args)+1,
			Title: args[0],
			Done: false,
		}

		tasks = append(tasks, newTasks)
		storage.SaveTasks(tasks)

		fmt.Println("Task added")
	},
}