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

func addRun(cmd *cobra.Command, args []string) {
	var items []todo.Item

	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}

	err := todo.SaveItems("./data.json", items)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data Saved Successfully")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
