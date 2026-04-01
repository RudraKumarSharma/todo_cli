package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"todo_list/internal/model"
	"todo_list/internal/storage"
)

var delCmd = &cobra.Command{
	Use: "del",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter a valid argument")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Task ID must be a number")
			return
		}

		var newTasks []model.Task
		data, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Failed to load tasks:", err)
			return
		}

		found := false

		for _, v := range data {
			if v.ID != taskID {
				newTasks = append(newTasks, v)
			} else {
				found = true
			}
		}
		if !found {
			fmt.Println("Task not found")
			return
		}

		if err := storage.SaveTasks(newTasks); err != nil {
			fmt.Println("Failed to save tasks:", err)
			return
		}
		fmt.Println("Task deleted")

	},
}