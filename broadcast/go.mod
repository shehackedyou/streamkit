module github.com/shehackedyou/streamkit/broadcast

go 1.15

replace (
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d
	github.com/shehackedyou/streamkit/broadcast/obs => ./obs
	github.com/shehackedyou/streamkit/broadcast/show => ./show
	github.com/shehackedyou/streamkit/broadcast/show/scene => ./show/scene
)

require (
	github.com/andreykaipov/goobs v0.12.1
	github.com/multiverse-os/cli v0.0.0-20230212101701-e7017a44551d
)
