/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print the last content of one or more targets",
	Long: `Prints the last entries of a given target.
	If no target is specified, then the last entries of the default target are printed.
For example:
# prints the last 30 entries of the default target
memo print -n 30

# print the last 10 entries of the 'schnulli' target
memo print -t schnulli

# print the last 10 entries of the 'schnulli' and the 'bulli' target
memo print -t schnulli -t bulli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.Flags().StringArrayP("target", "t", make([]string, 0), "targets to print. If not target is given, then the default target is printed")
	printCmd.Flags().Int32P("number", "n", 10, "number of lines printed from the target")
}
