package cmd

import (
	"todo_list/internal/model"
	"todo_list/internal/storage"
	"github.com/spf13/cobra"
	"fmt"
)

var delCmd = &cobra.Command{
	Use: "del",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter a valid argument")
			return
		}

		var newTasks []model.Task
		data, _ := storage.LoadTasks()

		for _, v := range data {
			if fmt.Sprint(v.ID) != args[0] {
				newTasks = append(newTasks, v)
			}
		}

		storage.SaveTasks(newTasks)
		fmt.Println("Task deleted")

	},
}