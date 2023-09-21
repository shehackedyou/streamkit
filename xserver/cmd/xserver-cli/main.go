package main

import (
	"fmt"
	"os"

	xserver "github.com/shehackedyou/streamkit/xserver"

	cli "github.com/multiverse-os/cli"
)

func main() {

	x11 := xserver.X11{
		Client: xserver.Connect(xserver.DefaultConfig()["xserver_host"]),
	}

	cmd, initErrors := cli.New(cli.App{
		Name:        "xserver-cli",
		Description: "xserver command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {

				fmt.Printf("x11.Client: %v\n", x11.Client)
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

	//x11App.X11.InitActiveWindow()

	// TODO: Probably want to load some settings from a YAML config to make things
	// easier

	//fmt.Printf("x11App:\n")

	//tick := time.Tick(x11App.Delay)
	//for {
	//	select {
	//	case <-tick:
	//		if x11App.X11.HasActiveWindowChanged() {
	//			fmt.Printf("HasActiveWindowChanged(): true\n")

	//			activeWindow := x11App.X11.ActiveWindow()
	//			fmt.Printf("  active_window_title: %s\n", activeWindow.Title)

	//			fmt.Printf("  x11.ActiveWindowTitle: %v\n", x11App.X11.ActiveWindowTitle)
	//			// NOTE: This worked to prevent it from repeating
	//			// HasActiveWindowChanged() over and over
	//			x11App.X11.CacheActiveWindow()

	//		} else {
	//			fmt.Printf("tick,...\n")
	//			fmt.Printf("  toolkit.X11.ActiveWindowTitle: %v\n", x11App.X11.ActiveWindowTitle)
	//			fmt.Printf(
	//				"  x11.ActiveWindow().Type.String(): %v\n",
	//				x11App.X11.ActiveWindow().Type.String(),
	//			)
	//		}
	//	}
	//}
}
