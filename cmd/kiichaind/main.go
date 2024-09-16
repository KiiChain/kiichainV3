package main

import (
	"os"

	"github.com/KiiChain/kiichainV3/app/params"
	"github.com/KiiChain/kiichainV3/cmd/kiichaind/cmd"

	"github.com/KiiChain/kiichainV3/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
