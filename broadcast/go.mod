module github.com/shehackedyou/streamkit/broadcast

go 1.15

replace (
	github.com/multiverse-os/cli => /home/user/go/src/github.com/multiverse-os/cli
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d
)

require (
	github.com/andreykaipov/goobs v0.12.1
	github.com/multiverse-os/cli v0.0.0-00010101000000-000000000000
)
