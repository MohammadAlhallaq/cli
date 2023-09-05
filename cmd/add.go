package cmd

import (
	"demo-cli/cmd/todo"
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add",
	Run: addRun,
}

func addRun(_ *cobra.Command, args []string) {
	var items []todo.Item

	for i, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		item.SetLabel(i + 1)
		items = append(items, item)
	}

	err := todo.SaveItems(dataFile, items)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Data Saved Successfully")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
