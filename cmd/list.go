package cmd

import (
	"fmt"
	"os"
	"task7/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.Alltasks()
		if err != nil {
			fmt.Println("Something is wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("All your tasks are completed!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d, %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
