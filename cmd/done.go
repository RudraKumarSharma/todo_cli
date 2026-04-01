package cmd

import (
	"fmt"
	"strconv"
	"todo_list/internal/storage"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use: "done",
	Short: "Mark tasks done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter valid argument")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Task ID must be a number")
			return
		}

		data, err := storage.LoadTasks()
		if err != nil {
			fmt.Println("Failed to load tasks:", err)
			return
		}

		found := false
		for i := range data {
			if data[i].ID == taskID {
				data[i].Done = true
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Task not found")
			return
		}

		if err := storage.SaveTasks(data); err != nil {
			fmt.Println("Failed to save tasks:", err)
			return
		}
		fmt.Println("Task marked done")
	},
}