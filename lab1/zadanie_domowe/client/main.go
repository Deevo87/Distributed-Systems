package main

import (
	"main/client_back"
)

func main() {
	client := client_back.NewClient("client3", ":8080")
	err := client.Start()
	if err != nil {
		return
	}
}
