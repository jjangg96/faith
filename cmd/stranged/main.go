package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/strangelove-ventures/strange-chain/app"
	"github.com/strangelove-ventures/strange-chain/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"strange-chain",
		"cosmos",
		app.DefaultNodeHome,
		"strange-chain-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
