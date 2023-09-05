/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "demo-cli",
}
var dataFile string
var priority int

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	CurrentUser, err := os.Getwd()

	if err != nil {
		fmt.Println(err.Error())
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", CurrentUser+string(os.PathSeparator)+"data.json", "data file to store items")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority:1,2,3")
}
