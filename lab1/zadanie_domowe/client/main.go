package main

func main() {
	client := NewClient("client1", ":12340", ":8080")
	err := client.Start("cze")
	if err != nil {
		return
	}
}
