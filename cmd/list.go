package cmd

import (
	"todo_list/internal/storage"
	"github.com/spf13/cobra"
	"fmt"
)

var lsCmd = &cobra.Command{
	Use: "list",
	Short: "list your todo",
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := storage.LoadTasks()
		for _, v := range data {
			status := "❌"
			if v.Done {
				status = "✅"
			}
			fmt.Printf("%d %s [%s]\n", v.ID, v.Title, status)
		}
	},
}