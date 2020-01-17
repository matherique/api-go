build:
	go build -o bin/api .
	
run:
	./bin/api

all: build run