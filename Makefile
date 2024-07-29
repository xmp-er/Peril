build:
	go build -gcflags=all=-w -o main.go
run:
	go run main.go 2>/dev/null