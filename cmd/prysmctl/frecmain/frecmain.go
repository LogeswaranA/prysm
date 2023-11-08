package frecmain

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	{
		Name:  "frecmain",
		Usage: "commands for dealing with Frec Chain frecmain",
		Subcommands: []*cli.Command{
			generateGenesisStateCmd,
		},
	},
}
