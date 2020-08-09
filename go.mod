module logAgent

go 1.14

replace (
	common => ./common
	etcd => ./etcd
	kafka => ./kafka
	tailog => ./tailog
	util => ./util
)

require (
	common v0.0.0-00010101000000-000000000000
	etcd v0.0.0-00010101000000-000000000000
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	go.etcd.io/etcd v3.3.22+incompatible // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	kafka v0.0.0-00010101000000-000000000000
	tailog v0.0.0-00010101000000-000000000000
	util v0.0.0-00010101000000-000000000000
)
