
all: build run

build:
	go build -o covid

run:
	./covid

draw:
	curl http://localhost:8080/draw