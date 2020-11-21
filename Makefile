all: build run

build:
	go build cmd/abeja/main.go

run:
	./abeja
