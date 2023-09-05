package cmd

import (
	"demo-cli/cmd/todo"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: listRun,
}

func listRun(_ *cobra.Command, _ []string) {

	d, err := todo.ReadItems(dataFile)

	if err != nil {
		fmt.Printf("%v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range d {
		_, err := fmt.Fprintln(w, i.Label()+" "+i.PrettyPri()+"\t"+i.Text+"\t")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	err = w.Flush()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
