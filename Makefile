main:
	go build -o wandsandlasers cmd/game/main.go
run:
	go run cmd/game/main.go
test:
	go test
bench:
	go test -bench=.
clean:
	rm wandsandlasers
