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
		config := config.Get()
		run(args, config)
	},
}

func run(args []string, config config.Config) {
	var memo add.Memo
	var err error
	if add.IsInteractiveMode(&args) {
		memo, err = add.GetMemoFromStdin(os.Stdin)
	} else {
		memo, err = add.ParseInput(&args)
	}
	if err != nil {
		fmt.Println("Error while parsing input: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Received memo: ", memo.Text)
		output := fmt.Sprintf("  included targets: %v,\n (number of targets: %d)", memo.Targets, len(memo.Targets))
		fmt.Println(output)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
