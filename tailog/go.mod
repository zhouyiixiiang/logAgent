module tailog

go 1.14

replace (
	etcd => ../etcd
	kafka => ../kafka
)

require (
	etcd v0.0.0-00010101000000-000000000000
	github.com/hpcloud/tail v1.0.0
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	kafka v0.0.0-00010101000000-000000000000
)
