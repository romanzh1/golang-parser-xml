FROM golang:latest

WORKDIR /app

COPY ./ ./

EXPOSE 4000

RUN go mod download
RUN go build -o parser-xml ./cmd/parser-xml/main.go

CMD ["./parser-xml"]