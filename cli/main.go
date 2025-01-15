package main

import (
	"os"

	"github.com/ixje/neo-go-legacy/cli/server"
	"github.com/ixje/neo-go-legacy/cli/smartcontract"
	"github.com/ixje/neo-go-legacy/cli/vm"
	"github.com/ixje/neo-go-legacy/cli/wallet"
	"github.com/ixje/neo-go-legacy/pkg/config"
	"github.com/urfave/cli"
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-go"
	ctl.Version = config.Version
	ctl.Usage = "Official Go client for Neo"

	ctl.Commands = append(ctl.Commands, server.NewCommands()...)
	ctl.Commands = append(ctl.Commands, smartcontract.NewCommands()...)
	ctl.Commands = append(ctl.Commands, wallet.NewCommands()...)
	ctl.Commands = append(ctl.Commands, vm.NewCommands()...)

	if err := ctl.Run(os.Args); err != nil {
		panic(err)
	}
}
