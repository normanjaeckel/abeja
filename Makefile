all: build run

build:
	go build ./cmd/abeja/

run:
	./abeja
