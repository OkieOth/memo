/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"okieoth/memo/internal/pkg/config"
	"okieoth/memo/internal/pkg/utils"
)

var regexpContentLine *regexp.Regexp

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
		targets, _ := cmd.Flags().GetStringArray("target")
		number, _ := cmd.Flags().GetInt32("number")
		config := config.Get()
		if len(targets) == 0 {
			targets = make([]string, 1)
			targets[0] = config.DefaultTarget
		}
		runPrint(&targets, number, config)
	},
}

func runPrint(targets *[]string, number int32, config config.Config) {
	prepareTargets(targets, config)
	for _, target := range *targets {
		b, e := utils.DoesFileExist(target)
		if (!b) || (e != nil) {
			fmt.Printf("-> target seems not to exist: %s, error: %v\n", target, e)
		} else {
			printTargetNow(target, number)
		}
	}
}

func getTargetShortName(target string) string {
	r := regexp.MustCompile(`.*/([a-zA-Z0-9]*)\.md$`)
	matches := r.FindAllStringSubmatch(target, -1)
	fmt.Printf("matches: %v\n", matches)
	return matches[0][1]
}

func isContentLine(line string) bool {
	if regexpContentLine == nil {
		// Example enty ...
		// * Das ist ein kleiner Test [202206226101035]
		regexpContentLine = regexp.MustCompile(`\s*\* [^\[]*\[\d{15}\]`)
	}
	return regexpContentLine.MatchString(line)
}

func getLastFileEntries(fileScanner *bufio.Scanner, entryCount int32) *[]string {
	ringBuffer := make([]string, entryCount)
	fileScanner.Split(bufio.ScanLines)

	currentTmp := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if isContentLine(line) {
			ringBuffer[currentTmp%int(entryCount)] = line
			currentTmp++
		}
	}

	// This is needed in cases where less content is provided than the desired number of elems
	if currentTmp < int(entryCount) {
		currentTmp = 0
	}
	returnBuffer := make([]string, entryCount)
	for i := 0; i < int(entryCount); i++ {
		returnBuffer[i] = ringBuffer[(currentTmp+i)%int(entryCount)]
	}
	return &returnBuffer
}

func printTargetNow(target string, number int32) {
	targetName := getTargetShortName(target)
	fmt.Printf("\nTarget: %s\n==================================\n", targetName)
	file, err := os.Open(target)
	if err != nil {
		fmt.Printf("ERROR while reading '%s': %v", target, err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	getLastFileEntries(fileScanner, number)
	// TODO print lastEntries
}

func prepareTargets(targets *[]string, config config.Config) {
	targetExt := ".md"
	for i, target := range *targets {
		var strBuilder strings.Builder
		strBuilder.WriteString(config.TargetDir)
		strBuilder.WriteRune('/')
		strBuilder.WriteString(target)
		strBuilder.WriteString(targetExt)
		(*targets)[i] = strBuilder.String()
	}
}

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.Flags().StringArrayP("target", "t", make([]string, 0), "targets to print. If not target is given, then the default target is printed")
	printCmd.Flags().Int32P("number", "n", 10, "number of lines printed from the target")
}
