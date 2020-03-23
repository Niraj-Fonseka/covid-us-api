
all: build run

build:
	go build -o covid

run:
	./covid


generate-new-data:
	curl htto://localhost:8080/newdata

rendertest:
	curl http://localhost:8080/render

renderquery:
	curl http://localhost:8080/render?page=
draw:
	curl http://localhost:8080/draw

test:
	curl http://localhost:8080/test

testrefactor:
	curl http://localhost:8080/testrefactor


generate-all: 
	curl http://localhost:8080/generate-daily
	curl http://localhost:8080/generate-summary

upload-mainpage:
	curl http://localhost:8080/upload-mainpage