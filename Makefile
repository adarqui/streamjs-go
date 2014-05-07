all:
	go build

test:
	go test -bench=".*" -test.v
