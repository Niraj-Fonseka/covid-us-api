
all: build run

build:
	go build -o covid

run:
	./covid


generate-new-data:
	curl htto://localhost:8080/newdata
draw:
	curl http://localhost:8080/draw

test:
	curl http://localhost:8080/test

testrefactor:
	curl http://localhost:8080/testrefactor
