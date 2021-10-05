all: build run
build:
	go build -o cmd/parser-xml/main.exe ./cmd/parser-xml
run:
	.\cmd.\parser-xml\main.exe
d:
	docker build -t parser-xml .
	docker run -p 4000:4000 parser-xml
