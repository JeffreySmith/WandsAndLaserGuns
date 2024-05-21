all: linux openbsd freebsd macos
main:
	go build -o wandsandlasers cmd/game/main.go
run:
	go run cmd/game/main.go
test:
	go test
bench:
	go test -bench=.
linux:
	GOOS=linux GOARCH=amd64 go build -o wandsandlasers_linux cmd/game/main.go
openbsd:
	GOOS=openbsd GOARCH=amd64 go build -o wandsandlasers_openbsd cmd/game/main.go
macos:
	GOOS=darwin GOARCH=arm64 go build -o wandsandlasers_macos cmd/game/main.go
freebsd:
	GOOS=freebsd GOARCH=amd64 go build -o wandsandlasers_freebsd cmd/game/main.go
clean:
	rm wandsandlasers*
