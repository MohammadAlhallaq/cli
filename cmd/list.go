package cmd

import (
	"demo-cli/cmd/todo"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: readRun,
}

func readRun(cmd *cobra.Command, args []string) {

	d, err := todo.ReadItems("./data.json")

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("Data Read Successfully")
	fmt.Println(d)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
