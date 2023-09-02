/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	items := []todo.Item{}
	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}
	fmt.Println(items)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
