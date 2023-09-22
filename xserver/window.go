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

// TODO: A function to move window would be great for setting up development
// environments or at the very least setting up streaming automatically

// TODO: InnerID is a md5 hashed value of a few things to get a unique thing,
// so while not the hash algo I would have used it is what we want.
type Window struct {
	ID      string // TODO: Maybe store innerID and see if its something we can use
	Title   string
	Command string
	PID     uint32
	Type    WindowType
	Focused bool // aka Active
	//X11     x11.Window // The base Window object from our library
	// eventually we should just load all this data into our window object and
	// then be able to do like .XWindow() => x11.Window type
	// There is also tons of window info data that may just be better to save
	// in the form x11.WindowInfo, and that stores X11.Window inside it

	LastUpdatedAt time.Time

	Rectangle
}

func EmptyWindow() *Window {
	return &Window{
		Type: UndefinedType,
	}
}

func (w *Window) IsEmpty() bool    { return w.Type == UndefinedType }
func IsWindowEmpty(w *Window) bool { return w.IsEmpty() }

// TODO: Can generate x11.WindowInfo from x11.Window
//func (x *X11) ParseWindow(xwin Window) (*Window, error) {
//	name, err := ewmh.GetWMName(x.Client, xwin).Reply(x.Client)
//	if err != nil {
//		return nil, err
//	} else {
//		fmt.Printf("\tName: ")
//		fmt.Printf("%s\n", name)
//	}
//
//	pidString, err := ewmh.GetWMPid(x.Client, xwin).Reply(x.Client)
//	if err != nil {
//		return nil, err
//	} else {
//		fmt.Printf("\tPid:%v\n", pidString)
//		data, _ := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pidString))
//		fmt.Printf("\t\tCmdline: %v\n", data)
//	}
//
//	pid, err := strconv.Atoi(pidString)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Window{
//		Title: name,
//		PID:   pid,
//	}, nil
//}

var UndefinedWindow = Window{Title: "", Type: UndefinedType}

type WindowType uint8 // 0..255

const (
	UndefinedType WindowType = iota
	Terminal
	Browser
	Other
)

//func (wt WindowType) String() string {
//	switch wt {
//	case Terminal:
//		return "terminal"
//	case Browser:
//		return "browser"
//	case Other:
//		return "other"
//	default: // UndefinedType
//		return "undefined"
//	}
//}
//
//func MarshalWindowType(wt string) WindowType {
//	switch strings.ToLower(wt) {
//	case Terminal.String():
//		return Terminal
//	case Browser.String():
//		return Browser
//	case Other.String():
//		return Other
//	default:
//		return UndefinedType
//	}
//}

//primaryWindow.TitleSuffixIs("chromium")

func (w *Window) TitleSuffixIs(searchText string) bool {
	return strings.HasSuffix(strings.ToLower(w.Title), searchText)
}

func (w *Window) WindowTitleIs(title string) bool {
	return strings.ToLower(w.Title) == title
}

func (w *Window) IsWindowType(windowType WindowType) bool {
	return w.Type == windowType
}
