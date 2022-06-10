package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"okieoth/memo/cmd/add"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new memo entry",
	Long: `Add a new memo to one or more "topics". To add a memo
type in the wanted text and annotate it with one or more hashtags.
The hashtags will be translated in storage locations, typically files.
As result the program responds with one ID per target for the stored memo.
This ID can be used to show later entries or to delete or done them.

For example:
# starts some kind of interactive mode, where you can enter your memo
# and the related targets
memo add

# add a new memo as one liner ...
# the text 'This is my first memo' is stored in target1 and target2
memo add This is my first memo \#target1 \#target2
...
`,
	Run: func(cmd *cobra.Command, args []string) {
		if add.IsInteractiveMode(&args) {
			fmt.Println("add called - interactive mode")
		} else {
			fmt.Println("add called - take the input from the commandline")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
