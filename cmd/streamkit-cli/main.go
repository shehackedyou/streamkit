package main

import (
	"fmt"
	"os"

	streamkit "github.com/shehackedyou/streamkit"

	cli "github.com/multiverse-os/cli"
)

func main() {

	cmd, initErrors := cli.New(cli.App{
		Name:        "streamkit-cli",
		Description: "streamkit command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		GlobalFlags: cli.Flags(
			cli.Flag{
				Name:        "test",
				Alias:       "t",
				Description: "Test function because I'm curious how this is interacting",
			}),
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("[onStart] preforming action...")

				toolkit, err := streamkit.New()
				if err != nil {
					fmt.Printf("error loading streamkit: %v\n", err)
				} else {
					fmt.Printf("toolkit: %v\n", toolkit)
				}
				//toolkit.HandleWindowEvents()
				// aDD all the listening and event driven stuff
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			//OnExit: func(c *cli.Context) error {
			//	c.CLI.Log("on exit action")
			//	return nil
			//},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
