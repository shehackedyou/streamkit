package xserver

import (
	"fmt"
	"io/ioutil"
	"time"

	x11 "github.com/linuxdeepin/go-x11-client"
	ewmh "github.com/linuxdeepin/go-x11-client/util/wm/ewmh"
)

type X11 struct {
	Client *x11.Conn
	//Desktops

	// TODO: This either needs to go entirely, and we save the Active or focused
	// window; storing a window like this is almost not sensible; but if we do it
	// has to be a pointer and MUST be a pointer to the exact same object in the
	// Windows []*Window otherwise its far less useful (this is do-able but takes
	// a bit more finesse)
	Window *Window

	// TODO: THis is the old way of avoiding unneeded updates but we can get
	// around this by properly subscribing and there are other methods; but for
	// now we can just keep doing these silly things for examples
	CurrentWindowTitle     string
	CurrentWindowChangedAt time.Time

	Windows
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"host": "localhost:10.0",
	}
}

func Connect(address string) *x11.Conn {
	client, err := x11.NewConnDisplay(address)
	if err != nil {
		panic(err)
	}
	return client
}

//func (x *X11) HasActiveWindowChanged() bool {
//	return x.Window.Title != x.ActiveWindow().Title
//}

func (x *X11) ActiveWindow() *Window {
	activeWindow, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("activeWindow: %v\n", activeWindow)

	activeWindowTitle, err := ewmh.GetWMName(
		x.Client,
		activeWindow,
	).Reply(x.Client)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}

	fmt.Printf("ActiveWindowTitle: %v\n", activeWindowTitle)

	pid, err := ewmh.GetWMPid(x.Client, activeWindow).Reply(x.Client)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	} else {
		fmt.Printf("\tPid:%v\n", pid)
		data, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
		fmt.Printf("\t\tCmdline: %s\n", data)
	}

	// TODO: Maybe have a cache window data or some such func
	return &Window{
		Title:         activeWindowTitle,
		PID:           pid,
		LastUpdatedAt: time.Now(),
	}
}

//func (x *X11) InitActiveWindow() *Window {
//	activeWindow := x.ActiveWindow()
//	x.Window.Title = activeWindow.Title
//	x.Window.LastUpdatedAt = time.Now()
//	return activeWindow
//}
//
//func (x *X11) CacheActiveWindow() *Window {
//	x.Window = x.ActiveWindow()
//	return x.Window
//}

// LOL SHORT POLLING AGAINST AN API WITH SUBSCRIPTIONS
//func ShortPoll() {
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
//}
