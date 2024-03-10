package main

import (
	"main/client_test"
)

func main() {
	client := client_test.NewClient("client1", ":12340", ":8080")
	err := client.Start()
	if err != nil {
		return
	}
}
