package main

import (
	"fmt"
)

func main() {
	server := NewServer(":8080")
	err := server.Start()
	if err != nil {
		fmt.Println("Error during start: ", err)
		return
	}
}
