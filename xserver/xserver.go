package xserver

import (
	"fmt"
	"time"

	x11 "github.com/linuxdeepin/go-x11-client"
	ewmh "github.com/linuxdeepin/go-x11-client/util/wm/ewmh"
)

type xWindow uint32

type WindowName uint8

type X11 struct {
	Client *x11.Conn
	//Desktops

	// TODO: This either needs to go entirely, and we save the Active or focused
	// window; storing a window like this is almost not sensible; but if we do it
	// has to be a pointer and MUST be a pointer to the exact same object in the
	// Windows []*Window otherwise its far less useful (this is do-able but takes
	// a bit more finesse)
	//Window *Window

	// TODO: THis is the old way of avoiding unneeded updates but we can get
	// around this by properly subscribing and there are other methods; but for
	// now we can just keep doing these silly things for examples
	CurrentWindowTitle     string
	CurrentWindowChangedAt time.Time

	Windows
}

func (x *X11) ActiveWindow() string {
	win, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		fmt.Println("error(ewmh.GetActiveWindow(x.Client)...):", err)
		return ""
	}

	activeWindowName, err := ewmh.GetWMName(x.Client, win).Reply(x.Client)
	if err != nil {
		fmt.Println("error(ewmh.GetWMName(x.Client, win)...):", err)
		return ""
	}

	return activeWindowName
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"host": "localhost:10.0",
	}
}

func Connect(address string) (client *x11.Conn) {
	// TODO: For now, we are doing simple then complex
	client, err := x11.NewConn()
	if err != nil {
		client, err = x11.NewConnDisplay(address)
		if err != nil {
			panic(err)
		}
	}
	return client
}

//func (x *X11) HasActiveWindowChanged() bool {
//	return x.Window.Title != x.ActiveWindow().Title
//}

//func (x *X11) ActiveWindow() *Window {
//	fmt.Printf("x or *X11 (%v)\n", x)
//	fmt.Printf("x.Client (%v)\n", x.Client)
//
//	active, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
//	if err != nil {
//		panic(err)
//	} else if active == 0 {
//		active = x.Client.GetDefaultScreen().Root
//	}
//
//	fmt.Printf("active window returned from GetActiveWindow().Reply(): %v\n", active)
//
//	activeWindowTitle, err := ewmh.GetWMName(x.Client, active).Reply(x.Client)
//	if err != nil {
//		fmt.Println("error(ewmh.GetWMName(x.Client, windowName).Reply(x.Client):", err)
//		//} else if len(activeWindowTitle) == 0 {
//		//	return EmptyWindow()
//	}
//
//	fmt.Printf("ActiveWindowTitle: %v\n", activeWindowTitle)
//
//	//pid, err := ewmh.GetWMPid(x.Client, active).Reply(x.Client)
//	//if err != nil {
//	//	fmt.Printf("get wm pid error(%v)\n", err)
//	//} else {
//	//	fmt.Printf("\tPid:%v\n", pid)
//	//	data, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
//	//	fmt.Printf("\t\tCmdline: %s\n", data)
//	//}
//
//	// TODO: Should only be returning a window if its actually not full of errors,
//	// then we also want a function that returns a empty window we can use for
//	// cases like this.
//
//	// TODO: Maybe have a cache window data or some such func
//	return &Window{
//		//Title: activeWindowTitle,
//		Type: Terminal,
//		//PID:   pid,
//		//LastUpdatedAt: time.Now(),
//	}
//}

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
