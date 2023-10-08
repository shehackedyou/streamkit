module github.com/shehackedyou/streamkit

go 1.15

replace (
	github.com/shehackedyou/streamkit/broadcast => ./broadcast
	github.com/shehackedyou/streamkit/terminal => ./terminal
	github.com/shehackedyou/streamkit/xserver => ./xserver
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text => github.com/multiverse-os/text v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

require (
	github.com/linuxdeepin/go-x11-client v0.0.0-20230710064023-230ea415af17
	github.com/multiverse-os/cli v0.1.0
	golang.org/x/term v0.8.0

)

require github.com/kr/pretty v0.3.1 // indirect
