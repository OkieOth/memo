/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"okieoth/memo/internal/pkg/config"
)

var flagPrint bool
var flagTargetDir string
var flagDefaultTarget string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "print or change the configuration",
	Long: `The configuration for the program is stored in ~/.memo/config.json. With
	this command you can print or change the current configuration.

	For example:
	# print the current configuration
	memo config -p
	
	# set the new target dir to store the memos
	memo config --targetDir $HOME/myTargets

	# set the new default target for cases when no target else is given
	memo config --defaultTarget $HOME/myTargets
	...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Get()
		if flagPrint {
			printConfig(&config)
		} else {
			changed := false
			if flagTargetDir != "" {
				config.TargetDir = flagTargetDir
				changed = true
			}
			if flagDefaultTarget != "" {
				config.DefaultTarget = flagDefaultTarget
				changed = true
			}
			if changed {
				err := config.Write()
				if err != nil {
					fmt.Printf("Error while change configuration: %v", err)

				} else {
					printConfig(&config)
				}
			}
		}

	},
}

func printConfig(config *config.Config) {
	output, _ := config.AsJson()
	_, _ = fmt.Printf("Current memo configuration:\n\n%s\n\n", output)
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&flagPrint, "print", "p", false, "print current configuration")
	configCmd.Flags().StringVarP(&flagTargetDir, "targetDir", "t", "", "set target dir to store the memos")
	configCmd.Flags().StringVarP(&flagDefaultTarget, "defaultTarget", "d", "", "set default target")
}
