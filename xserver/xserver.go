package xserver

import (
	"fmt"
	"io/ioutil"
	"time"

	x11 "github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-x11-client/util/wm/ewmh"
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
		"host": "localhost:11.0",
	}
}

func Connect(address string) (client *x11.Conn) {
	// TODO: Default is going to check localhost:10.0 so we try 11.0
	client, err := x11.NewConn()
	if err != nil {
		client, err = x11.NewConnDisplay(address)
		if err != nil {
			panic(err)
		}
	}
	return client
}

func (x *X11) HasActiveWindowChanged() bool {
	return x.Window.PID != x.ActiveWindow().PID
}

func (x *X11) ActiveWindow() *Window {
	active, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		fmt.Printf("error: no window returning undefined\n")
		return UndefinedWindow()
	}

	title, err := ewmh.GetWMName(x.Client, active).Reply(x.Client)
	if err != nil {
		fmt.Printf("error: no window title returning undefined\n")
		return UndefinedWindow()
	}

	var byteBuffer []byte
	pid, err := ewmh.GetWMPid(x.Client, active).Reply(x.Client)
	if err != nil {
		fmt.Printf("error: no window pid returning undefined\n")
		return UndefinedWindow()
	} else {
		byteBuffer, err = ioutil.ReadFile(
			fmt.Sprintf("/proc/%d/cmdline", pid),
		)
		if err != nil {
			fmt.Printf("error: no window cmdline returning undefined\n")
			return UndefinedWindow()
		}
	}

	window := &Window{
		Title:   title,
		PID:     pid,
		Command: fmt.Sprintf("%s", byteBuffer),
	}

	window.Type = window.WindowType()
	window.LastUpdatedAt = time.Now()

	return window
}

func (x *X11) CacheActiveWindow() *Window {
	x.Window = x.ActiveWindow()
	return x.Window
}
