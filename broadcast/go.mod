module github.com/shehackedyou/streamkit/broadcast

go 1.15

require github.com/multiverse-os/cli v0.0.0-20230212101701-e7017a44551d

//replace (
//	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.0.0
//	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d
//	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
//	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
//)

require github.com/andreykaipov/goobs v0.12.1

replace github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
