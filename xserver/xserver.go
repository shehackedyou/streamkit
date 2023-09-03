package xserver

import (
	"fmt"
	"time"

	x11 "github.com/linuxdeepin/go-x11-client"
	ewmh "github.com/linuxdeepin/go-x11-client/util/wm/ewmh"
)

type X11 struct {
	Client *x11.Conn // 	xdisplay       *x.Conn

	//Desktops

	// TODO
	// When needed bother to store the history of active windows but that
	// isn't needed quite yet, so there is about ZERO point in implementing
	// it.

	// TODO: Maybe just cache the active window name so we do simple name
	// comparison, but this leads to a bug where two windows with the same name
	// are considered the name window
	Windows []*Window

	CurrentWindowTitle     string
	CurrentWindowChangedAt time.Time
}

func ConnectTo(addr string) *x11.Conn {
	conn, err := x11.NewConnDisplay(addr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("conn: %v\n", conn)

	return conn
}

//func (x *) HasActiveWindowChanged() bool {
//	return x.ActiveWindowTitle != x.ActiveWindow().Title
//}

func (x *X11) CurrentWindow() *Window {
	//	var err error
	//	// TODO: This returns an x.Window object which can get all sorts of
	//	// information beyond just the name, like the PID. We shouldn't need a second
	fmt.Printf("trying active window\n")
	//	// call at all to get the title of the window, thats absurdist.
	activeWindow, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("active_window: %v\n", activeWindow)

	// TODO: Do we actually need to do GetWMName? Shouldn't we actually do the
	// GetWindowInfo thing so we get it and much more information we could cache
	//activeWindowTitle, err := ewmh.GetWMName(
	//	x.Client,
	//	activeWindow,
	//).Reply(x.Client)
	//if err != nil {
	//	fmt.Printf("error(%v)\n", err)
	//}

	//pid, err := ewmh.GetWMPid(x.Client, activeWindow).Reply(x.Client)
	//if err != nil {
	//	fmt.Printf("error(%v)\n", err)
	//} else {
	//	fmt.Printf("\tPid:%v\n", pid)
	//	data, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	//	fmt.Printf("\t\tCmdline: %s\n", data)
	//}

	// TODO: Maybe have a cache window data or some such func
	//cachedWindow := &Window{
	//	//PID: uint32(4444),
	//}

	//return cachedWindow
	return nil
}

//func (x *X11) InitActiveWindow() *Window {
//	activeWindow := x.ActiveWindow()
//	x.ActiveWindowTitle = activeWindow.Title
//	x.ActiveWindowChangedAt = time.Now()
//	return activeWindow
//}
//
//func (x *X11) CacheActiveWindow() *Window {
//	activeWindow := x.ActiveWindow()
//	x.ActiveWindowTitle = x.ActiveWindow().Title
//	x.ActiveWindowChangedAt = time.Now()
//	return activeWindow
//}
