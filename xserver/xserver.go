package xserver

import (
	"fmt"
	"io/ioutil"
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
	Windows []*Window // TODO: Or save the position in the slice (or linked list that is the active one, or even use linked list to put them in stack order and top is active.

	Window *Window

	CurrentWindowTitle     string
	CurrentWindowChangedAt time.Time
}

func Connect(addr string) *x11.Conn {
	conn, err := x11.NewConnDisplay(addr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("*x11.conn: %v\n", conn)

	return conn
}

func (x *X11) HasActiveWindowChanged() bool {
	return x.Window.Title != x.ActiveWindow().Title
}

func (x *X11) ActiveWindow() *Window {
	activeWindow, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ActiveWindow(): %v\n", activeWindow)

	activeWindowTitle, err := ewmh.GetWMName(
		x.Client,
		activeWindow,
	).Reply(x.Client)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	}

	cachedWindow := &Window{
		Title: activeWindowTitle,
	}

	pid, err := ewmh.GetWMPid(x.Client, activeWindow).Reply(x.Client)
	if err != nil {
		fmt.Printf("error(%v)\n", err)
	} else {
		fmt.Printf("\tPid:%v\n", pid)
		data, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
		fmt.Printf("\t\tCmdline: %s\n", data)
	}

	cachedWindow.PID = pid
	cachedWindow.LastUpdatedAt = time.Now()

	// TODO: Maybe have a cache window data or some such func
	return cachedWindow
}

func (x *X11) InitActiveWindow() *Window {
	activeWindow := x.ActiveWindow()
	x.Window.Title = activeWindow.Title
	x.Window.LastUpdatedAt = time.Now()
	return activeWindow
}

func (x *X11) CacheActiveWindow() *Window {
	x.Window = x.ActiveWindow()
	return x.Window
}
