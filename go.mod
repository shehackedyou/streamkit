module github.com/shehackedyou/streamkit

go 1.19

require (
	github.com/multiverse-os/cli v0.1.0
	github.com/shehackedyou/streamkit/broadcast v0.0.0-20231213124023-a17906fa3e5d
	github.com/shehackedyou/streamkit/xserver v0.0.0-20231213124023-a17906fa3e5d
	golang.org/x/term v0.15.0
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/bannner => github.com/multiverse-os/banner v0.1.0
)
