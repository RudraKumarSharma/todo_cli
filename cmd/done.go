package cmd

import (
	"todo_list/internal/storage"
	"github.com/spf13/cobra"
	"fmt"
)

var doneCmd = &cobra.Command{
	Use: "done",
	Short: "Mark tasks done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter valid argument")
			return
		}

		data, _ := storage.LoadTasks()

		for  _, v := range data {
			if fmt.Sprint(v.ID) == args[0] {
				v.Done = true
			}
		}

		storage.SaveTasks(data)
		fmt.Println("Task marked done")
	},
}