package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/jjangg96/faith/app"
	"github.com/jjangg96/faith/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"faith",
		"cosmos",
		app.DefaultNodeHome,
		"faith-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
