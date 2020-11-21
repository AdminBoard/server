module github.com/adminboard/server

go 1.15

//replace github.com/eqto/api-server => /home/tuxer/Projects/Tuxer/api-server

require (
	github.com/eqto/api-server v0.5.1-0.20201121064122-e5d09bfc7059
	github.com/eqto/command v0.3.2
	github.com/eqto/config v0.1.1
	github.com/eqto/go-db v0.3.0
	github.com/eqto/go-json v0.2.2
	github.com/eqto/go-logger v0.2.3
	github.com/pkg/errors v0.9.1 // indirect
	github.com/speps/go-hashids v2.0.0+incompatible
)
