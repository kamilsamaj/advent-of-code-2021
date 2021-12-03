package main

import "fmt"

func main() {
	fmt.Println(`
All tasks are stored as a separate command in the '/cmd' directory.

You can run them with:
go run cmd/01/main.go

And test them with:
go test -v cmd/01/*.go`)
}
