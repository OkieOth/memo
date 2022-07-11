/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"okieoth/memo/internal/pkg/config"
	"okieoth/memo/internal/pkg/utils"
	"os"
	"regexp"
	"time"
)

var flagList bool
var flagWide bool

// targetsCmd represents the targets command
var targetsCmd = &cobra.Command{
	Use:   "targets",
	Short: "Maintains the targets stored on this machine",
	Long:  `If you want to know what targets are available use this command`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Get()
		if flagWide {
			printWideTargetInfos(config.TargetDir)
		} else {
			if flagList {
				printTargets(config.TargetDir)
			} else {
				printShortTargetInfos(config.TargetDir)
			}
		}
	},
}

func printWideTargetInfos(targetDir string) {
	printTargetsBase(targetDir, longTargetPrint)
}

func longTargetPrint(targetName string, fileInfo os.FileInfo) {
	now := time.Now()
	diff := now.Sub(fileInfo.ModTime())
	hrs := int(diff.Hours())
	mins := int(diff.Minutes())
	second := int(diff.Seconds())
	days := int(diff.Hours() / 24)
	var lastChangedStr string
	if days > 0 {
		if days == 1 {
			lastChangedStr = "1 day ago"
		} else {
			lastChangedStr = fmt.Sprintf("%d days ago", days)
		}
	} else {
		if hrs == 1 {
			lastChangedStr = "1 h ago"
		} else {
			if hrs > 1 {
				lastChangedStr = fmt.Sprintf("%d hours ago", hrs)
			} else {
				if mins > 0 {
					lastChangedStr = fmt.Sprintf("%d' ago", mins)
				} else {
					lastChangedStr = fmt.Sprintf("%d\" ago", second)
				}
			}
		}
	}
	fmt.Printf("  %s\t\t(last changed: %s, size: %d bytes)\n", targetName, lastChangedStr, fileInfo.Size())
}

func simpleTargetPrint(targetName string, fileInfo os.FileInfo) {
	fmt.Printf("  %s\n", targetName)
}

func printTargets(targetDir string) {
	printTargetsBase(targetDir, simpleTargetPrint)
}

type printFunc func(targetName string, fileInfo os.FileInfo)

func printTargetsBase(targetDir string, pf printFunc) {
	td, _ := utils.ReplaceEnvVars(targetDir)
	fileInfoSlice, err := ioutil.ReadDir(td)
	if err != nil {
		panic(err)
	}
	r, _ := regexp.Compile(`\.md$`)
	fmt.Println("\nThe current targets are:")
	for _, fileInfo := range fileInfoSlice {
		fileName := fileInfo.Name()
		if (!fileInfo.IsDir()) && r.MatchString(fileName) {
			targetName := r.ReplaceAllString(fileName, "")
			pf(targetName, fileInfo)
		}
	}
}

func printShortTargetInfos(targetDir string) {
	fmt.Printf("Targets are stored in: %s", targetDir)
}

func init() {
	rootCmd.AddCommand(targetsCmd)
	targetsCmd.Flags().BoolVarP(&flagList, "list", "l", false, "print current used targets")
	targetsCmd.Flags().BoolVarP(&flagWide, "wide", "w", false, "print current used targets and some extra infos")
}
