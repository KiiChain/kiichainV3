package tools

import (
	"github.com/KiiChain/kiichainV3/tools/tx-scanner/cmd"
	"github.com/spf13/cobra"
)

func ToolCmd() *cobra.Command {
	toolsCmd := &cobra.Command{
		Use:   "tools",
		Short: "A set of useful tools for sei chain",
	}
	toolsCmd.AddCommand(cmd.ScanCmd())
	return toolsCmd
}
