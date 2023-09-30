module github.com/shehackedyou/streamkit

go 1.15

replace (
	github.com/multiverse-os/cli => /home/user/go/src/github.com/multiverse-os/cli
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d
	github.com/shehackedyou/streamkit/broadcast => ./broadcast
	github.com/shehackedyou/streamkit/terminal => ./terminal
	github.com/shehackedyou/streamkit/xserver => ./xserver
)

require (
	github.com/multiverse-os/cli v0.0.0-20230212101701-e7017a44551d
	golang.org/x/term v0.8.0
)

require github.com/shehackedyou/streamkit/xserver v0.0.0-00010101000000-000000000000

require (
	github.com/kr/pretty v0.3.1 // indirect
	github.com/linuxdeepin/go-x11-client v0.0.0-20230710064023-230ea415af17
	github.com/shehackedyou/streamkit/broadcast v0.0.0-00010101000000-000000000000
)
