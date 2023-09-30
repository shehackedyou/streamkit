package streamkit

import (
	"fmt"
	"os"
	"time"

	broadcast "github.com/shehackedyou/streamkit/broadcast"
	xserver "github.com/shehackedyou/streamkit/xserver"
)

//list, _ := client.Inputs.GetInputList()
//import typedefs "github.com/andreykaipov/goobs/api/typedefs"
//// Represents the request body for the GetSceneItemList request.
//type GetSceneItemListParams struct {

type PathType uint8

const (
	Config PathType = iota
	Data
)

type Path string

type Toolkit struct {
	// NOTE: Short-poll rate [we will rewrite without short polling after]
	Delay time.Duration
	Show  *broadcast.Show
	X11   *xserver.X11
	// TODO: Our local copy of the show is entirely separate from obs.Client so we
	// can change that out while maintaining logic and a data object
	Config map[string]string
	Paths  map[PathType]Path
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"name":    "she hacked you",
		"obs":     broadcast.DefaultConfig()["host"],
		"xserver": xserver.DefaultConfig()["host"],
	}
}

// TODO: Could pass the host for OBS and the host for X11 as variadic strings so
// it can be empty, or provide position 1 for obs position 2 for x11 (though x11
// should assumingly always be 127.0.0.1 whereas obs reasonably could be
// different

// TODO: Obvio we need to be passing in the fucking config not just have it be
// hardcoded like shitty law
func New() (toolkit *Toolkit) {
	// TODO: Show should be from config, and obs and x11 information. Logically
	// stored in ~/.config/$APP_NAME and the local data should be
	// ~/.local/share/$APP_NAME

	// TODO: This would be defined the CLI and passed in or at the VERY least this
	// would be set to a function that returns this as DefaultConfig(); use
	// variadic input and when that variadic input is empty then we resort to
	// using this
	toolkitConfig := DefaultConfig()

	appName := "streamkit"
	userHome, _ := os.UserHomeDir()

	toolkit = &Toolkit{
		Config: toolkitConfig,
		Show: broadcast.OpenShow(
			toolkitConfig["name"],
			toolkitConfig["obs"],
		),
		X11:   xserver.NewWindowManager(toolkitConfig["xserver"]),
		Delay: 1500 * time.Millisecond,
		Paths: map[PathType]Path{
			Config: Path(fmt.Sprintf("%v/.config/%v", userHome, appName)),
			Data:   Path(fmt.Sprintf("%v/.local/share/%v", userHome, appName)),
		},
	}

	primaryScene := toolkit.Show.Scene("Primary")
	browserItem := primaryScene.ItemByName("Research Browser")
	toolkit.X11.OnChange[xserver.OnActiveWindow] = func() {
		// Setup changing OBS stuff based on window changes
		fmt.Printf("toolkit.X11.Window(%v)\n", toolkit.X11.Window)
		switch toolkit.X11.Window.Type {
		case xserver.Terminal:
			fmt.Printf("we got TERMINAL type, so wait 3 seconds...\n")
			time.Sleep(3 * time.Second)
			fmt.Printf("  1. then we transition to Primary\n")
			primaryScene.Transition()
			fmt.Printf("  2. then we HIDE research browser\n")
			browserItem.Hide()
		case xserver.Browser:
			fmt.Printf("we got BROWSER type, so wait 3 seconds...\n")
			time.Sleep(3 * time.Second)
			fmt.Printf("  1. then we transition to Primary\n")
			primaryScene.Transition()
			fmt.Printf("  2. then we UNHIDE research browser\n")
			browserItem.Unhide()
		default:
			fmt.Printf("we got a undefined type, thats cool, probably desktop...\n")
		}
	}

	fmt.Printf("toolkit.X11: %v\n", toolkit.X11)

	activeWindow := toolkit.X11.ActiveWindow()
	fmt.Printf("activeWindow reply is (%v)\n", activeWindow)

	toolkit.Show.YAML(0)

	return toolkit
}
