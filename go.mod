module github.com/shehackedyou/streamkit

go 1.19

require (
	github.com/multiverse-os/cli v0.1.0
	github.com/shehackedyou/streamkit/broadcast v0.0.0-20231218005902-1751dbc6d9fd
	github.com/shehackedyou/streamkit/xserver v0.0.0-20231218001123-f8f9196ec064
	golang.org/x/term v0.15.0
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

require (
	github.com/andreykaipov/goobs v0.12.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/linuxdeepin/go-x11-client v0.0.0-20230710064023-230ea415af17 // indirect
	github.com/multiverse-os/banner v0.0.0-20231006133835-80f8c892b073 // indirect
	github.com/multiverse-os/cli/data v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/ansi v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/loading v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
