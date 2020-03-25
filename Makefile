
all: build run

build:
	go build -o covid

run:
	./covid



render:
	curl http://localhost:8080/render

generate-daily:
	curl http://localhost:8080/generate-daily

generate-summary:
	curl http://localhost:8080/generate-summary

generate-county:
	curl http://localhost:8080/generate-county


generate-all: 
	curl http://localhost:8080/generate-daily
	curl http://localhost:8080/generate-summary

upload-mainpage:
	curl http://localhost:8080/upload-mainpage

upload-statepages:
	curl http://localhost:8080/upload-statespages

upload-all: upload-mainpage upload-statepages

run-all: generate-all render upload-all
