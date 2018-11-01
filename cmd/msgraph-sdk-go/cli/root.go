package cli

import "github.com/spf13/cobra"

const (
	RootCmdRootText = "msgraph"
)

var RootCmd = &cobra.Command{
	Use: RootCmdRootText,
	Run: func(cmd *cobra.Command, args []string){

	},
}


func NewCmd