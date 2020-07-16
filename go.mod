module logAgent

go 1.14

replace (
	common => ./common
	kafka => ./kafka
	tailog => ./tailog
)

require (
	common v0.0.0-00010101000000-000000000000
	github.com/hpcloud/tail v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	kafka v0.0.0-00010101000000-000000000000
	tailog v0.0.0-00010101000000-000000000000
)
