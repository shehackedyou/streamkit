package xserver

import (
	"fmt"
	"io/ioutil"
	"time"

	x11 "github.com/linuxdeepin/go-x11-client"
	ewmh "github.com/linuxdeepin/go-x11-client/util/wm/ewmh"
)

type xWindow uint32

type WindowName uint8

type ChangeType int

const (
	Undefined ChangeType = iota
	OnActiveWindow
)

type X11 struct {
	Client   *x11.Conn
	Window   *Window
	OnChange map[ChangeType]func()
	//Desktops
	//Windows
}

func NewWindowManager(host string) *X11 {
	x := &X11{
		Client:   Connect(host),
		OnChange: make(map[ChangeType]func()),
	}
	x.Window = x.ActiveWindow()
	return x
}

func DefaultConfig() map[string]string {
	return map[string]string{
		"host": "localhost:11.0",
	}
}

func Connect(host string) (client *x11.Conn) {
	// TODO: Default is going to check localhost:10.0 so we try 11.0
	client, err := x11.NewConn()
	if err != nil {
		client, err = x11.NewConnDisplay(host)
		if err != nil {
			panic(err)
		}
	}
	return client
}

func (x *X11) ActiveWindowMonitor() (err error) {
	root := x.Client.GetDefaultScreen().Root
	fmt.Printf("active window monitor\n")
	err = x11.ChangeWindowAttributesChecked(
		x.Client,
		root,
		x11.CWEventMask,
		[]uint32{x11.EventMaskPropertyChange},
	).Check(x.Client)
	if err != nil {
		fmt.Printf("err(%v)\n", err)
		return err
	}

	activeWindowAtom, err := x.Client.GetAtom("_NET_ACTIVE_WINDOW")
	if err != nil {
		panic(fmt.Errorf("failed to get _NET_ACTIVE_WINDOW atom:%v\n", err))
	}

	events := make(chan x11.GenericEvent, 10)
	x.Client.AddEventChan(events)

	for event := range events {
		switch event.GetEventCode() {
		case x11.PropertyNotifyEventCode:
			activeWindowEvent, _ := x11.NewPropertyNotifyEvent(event)
			if activeWindowEvent.Atom == activeWindowAtom && activeWindowEvent.Window == root {
				x.Window = x.ActiveWindow()
				x.OnChange[OnActiveWindow]()
			}
		}
	}
	return nil
}

func (x *X11) ActiveWindow() *Window {
	active, err := ewmh.GetActiveWindow(x.Client).Reply(x.Client)
	if err != nil {
		fmt.Printf("error: no window returning undefined\n")
		return UndefinedWindow()
	}

	if active != 0 {
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
			X11:     x,
			Title:   title,
			PID:     pid,
			Command: fmt.Sprintf("%s", byteBuffer),
		}

		window.Type = window.WindowType()
		window.LastUpdatedAt = time.Now()

		return window
	} else {
		return nil
	}
}
