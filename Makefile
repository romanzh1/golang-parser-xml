all: build run
build:
	go build -o cmd/parser-xml/main.exe ./cmd/parser-xml
run:
	.\cmd.\parser-xml\main.exe