package xserver

import (
	"strings"
	"time"
)

//var skipTaskBarWindowTypes = []string{
//	"_NET_WM_WINDOW_TYPE_UTILITY",
//	"_NET_WM_WINDOW_TYPE_COMBO",
//	"_NET_WM_WINDOW_TYPE_DESKTOP",
//	"_NET_WM_WINDOW_TYPE_DND",
//	"_NET_WM_WINDOW_TYPE_DOCK",
//	"_NET_WM_WINDOW_TYPE_DROPDOWN_MENU",
//	"_NET_WM_WINDOW_TYPE_MENU",
//	"_NET_WM_WINDOW_TYPE_NOTIFICATION",
//	"_NET_WM_WINDOW_TYPE_POPUP_MENU",
//	"_NET_WM_WINDOW_TYPE_SPLASH",
//	"_NET_WM_WINDOW_TYPE_TOOLBAR",
//	"_NET_WM_WINDOW_TYPE_TOOLTIP",
//}

// func SetActiveWindow(c *x.Conn, val x.Window)

// TODO: Remember in X11 the term "client" is the concept of the Window
// So for example this function

//     func SetClientList(c *x.Conn, vals []x.Window)

// This gets a list of windows (aka ClientList)

type Windows []*Window

// TODO: We need to switch to developing around Windows as a collection too
// which will make the xserver client more sensible; be able to easily pull
// specific windows, have a .Focus()

// TODO The point of this is to give us a collection object we can create
// methods from like

//func (windows Windows) Window(id string) *Window {
//	for _, window := range windows {
//		if window.ID == id {
//			return window
//		}
//	}
//	return nil
//}

// TODO: We could hash the title with the PID to get a more unique identifier to
// check with so we avoid the windows with the same title being the same. (Why
// are these not uint?)
type Position struct {
	X int16
	Y int16
}

// lol didnt use point
// TODO: This will need layer number, not just is in focus but guess that makes
// more sense within window than rectangle even though its the rectangle itself
// that is layered on the desktop.
type Rectangle struct {
	X, Y          int16
	Width, Height uint16
}

type Window struct {
	X11           *X11
	Title         string
	Command       string
	PID           uint32
	Type          WindowType
	LastUpdatedAt time.Time
}

func UndefinedWindow() *Window {
	return &Window{
		X11:           nil,
		Title:         "",
		Command:       "",
		PID:           0,
		Type:          UndefinedType,
		LastUpdatedAt: time.Now(),
	}
}

func (w *Window) IsUndefined() bool    { return w.Type == UndefinedType }
func IsWindowUndefined(w *Window) bool { return w.IsUndefined() }

type WindowType uint8 // 0..255

const (
	UndefinedType WindowType = iota
	Terminal
	Browser
	Other
)

func (wt WindowType) String() string {
	switch wt {
	case Terminal:
		return "terminal"
	case Browser:
		return "browser"
	case Other:
		return "other"
	default: // UndefinedType
		return "undefined"
	}
}

func MarshalWindowType(wt string) WindowType {
	switch strings.ToLower(wt) {
	case Terminal.String():
		return Terminal
	case Browser.String():
		return Browser
	case Other.String():
		return Other
	default:
		return UndefinedType
	}
}

func (w *Window) WindowType() WindowType {
	switch {
	case w.CommandContains("gnome-terminal-server"):
		return Terminal
	case w.HasTitleSuffix("chromium"):
		return Browser
	default:
		return UndefinedType
	}
}

func (w *Window) IsType(wt WindowType) bool {
	return w.Type == wt
}

func (w *Window) HasTitleSuffix(suffix string) bool {
	return strings.HasSuffix(strings.ToLower(w.Title), suffix)
}

func (w *Window) CommandContains(search string) bool {
	return strings.Contains(w.Command, search)
}

func (w *Window) HasTitle(title string) bool {
	return strings.ToLower(w.Title) == title
}
