package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/cmd/cmdperm"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/ui"
	"github.com/spf13/cobra"
)

var wrapupCmd = &cobra.Command{
	Use:   "wrapup",
	Short: "Add changes to index and auto commit",
	Long: fmt.Sprintf(`
%s adds the files in the workin directory to the index and makes a commit. The commit message is autogenerated and contains 'close issue', e.g. "close #42 fetch meaning of life". The issue number is fetched from the branch name, if the workon command have been used. When force is used on the message, the message will replace, the otherwice automatically generated commit message.
`, ui.Format.Bold("wrapup")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.WrapUp()
	},
}

func init() {
	RootCmd.AddCommand(wrapupCmd)

	wrapupCmd.Flags().StringVar(&options.GlobalFlagForceMessage, "force", "", "use a custom commit message instead")
}
