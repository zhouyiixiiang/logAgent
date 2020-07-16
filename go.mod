module logAgent

go 1.14

replace (
	common => ./common
	kafka => ./kafka
	tailog => ./tailog
)

require kafka v0.0.0-00010101000000-000000000000
