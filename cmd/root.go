/*
Copyright © 2022 Eiko Thomas

*/
package cmd

import (
	"bufio"
	"fmt"
	"okieoth/memo/cmd/add"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "memo",
	Short: "A tool to create and manage memos from the terminal",
	Long: `The tool supports the creation, search and list of on-the-fly
textual memos from the terminal.
`,
	// this command is executed if no command was given over the command line
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("command? [add|del|done|find|list|exit]: ")
		stayInLoop := true
		for stayInLoop {
			stayInLoop = false
			scanner.Scan()
			switch scanner.Text() {
			case "add":
				text := add.GetMemoTextFromStdin(os.Stdin)
				targets := add.GetTargetsFromStdin(os.Stdin)
				header := add.GetMemoHeaderFromStdin(os.Stdin)
				addCmd.Flags().Set("text", text)
				addCmd.Flags().Set("target", targets)
				if len(header) > 0 {
					addCmd.Flags().Set("header", header)
				}
				addCmd.Run(addCmd, args)
			case "del":
				delCmd.Run(cmd, args)
			case "done":
				doneCmd.Run(cmd, args)
			case "find":
				findCmd.Run(cmd, args)
			case "list":
				listCmd.Run(cmd, args)
			case "exit":
				fmt.Println("bye.")
				return
			default:
				stayInLoop = true
				fmt.Print("enter a valid command? [add|del|done|find|list|exit]: ")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.memo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
