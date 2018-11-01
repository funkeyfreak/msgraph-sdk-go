package auth

import (
	"github.com/spf13/cobra"

	"../../../../internal/cli-runtime/pkg/genericcli" //TODO: Replace with package import post build
	cmdUtil "../../util"
)

const (
	AuthCmdUse = "auth"
	AuthCmdShort = "Configure Authorization"
	AuthCmdLong = "Configure Authorization for msgraph-sdk-go"
)

func NewCmdAuth(streams genericcli.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use: AuthCmdUse,
		Short: AuthCmdShort,
		Long: AuthCmdLong,
		Run: cmdUtil.DefaultSubCommandRun(streams.ErrOut),
	}

	return cmd
}