package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"okieoth/memo/cmd/add"

	"okieoth/memo/internal/pkg/config"
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
		header, _ := cmd.Flags().GetString("header")
		targets, _ := cmd.Flags().GetStringArray("target")
		text, _ := cmd.Flags().GetString("text")

		var memo add.Memo
		memo.Header = header
		memo.Targets = targets
		memo.Text = text
		config := config.Get()
		run(args, config, memo)
	},
}

func run(args []string, config config.Config, memo add.Memo) {
	if len(memo.Text) == 0 {
		var stdin add.InitStdin
		stdin.Stdin = os.Stdin
		add.InitMemoFromStdin(stdin, &memo)
	}
	err := add.StoreMemo(memo, config)
	if err != nil {
		fmt.Printf("Error while store memo: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringArrayP("target", "t", make([]string, 0), "targets to store this memo")
	addCmd.Flags().StringP("header", "H", "", "header used for this memo")
	addCmd.Flags().StringP("text", "T", "", "text to store")
}
