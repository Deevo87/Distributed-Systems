package main

import (
	"context"
	"fmt"
	"log"
	_ "os"

	"google.golang.org/grpc"

	pb "client/generated/dynamic_executors" // Import your generated proto package
)

const (
	hostname = "localhost"
	port     = 5000
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", hostname, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewExecutionServiceClient(conn)

	// Input loop
	for {
		fmt.Print("> ")
		var executableName string
		if _, err := fmt.Scanln(&executableName); err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		switch executableName {
		case "add":
			execute(client, "JaredProject-1.0-SNAPSHOT.jar", "org.solution.Calculator", "add", "[1, 2, -3, 123]")
		case "average":
			execute(client, "JaredProject-1.0-SNAPSHOT.jar", "org.solution.Calculator", "average", "[10, 10, 10, 10, 10]")
		case "power":
			execute(client, "JaredProject-1.0-SNAPSHOT.jar", "org.solution.Calculator", "power", "[1, 2, 3, 4, 5]")
		case "toUpper":
			execute(client, "JaredProject-1.0-SNAPSHOT.jar", "org.solution.Essa", "toUpper", "only small characters")
		case "hello":
			execute(client, "JaredProject-1.0-SNAPSHOT.jar", "org.solution.Essa", "hello", "Pablo Emilio Escobar Gaviria")
		case "quit":
			fmt.Println("Quitting")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

func execute(client pb.ExecutionServiceClient, jarLocation, className, methodName, data string) {
	// Make gRPC call
	resp, err := client.Execute(context.Background(), &pb.ExecutionRequest{
		JarLocation: jarLocation,
		ClassName:   className,
		MethodName:  methodName,
		Data:        data,
	})
	if err != nil {
		log.Fatalf("RPC Error: %v", err)
	}
	// Process response
	if resp.GetErrCode() != "" {
		fmt.Printf("Server returned an error: '%s'\n", resp.GetErrCode())
	} else {
		fmt.Printf("Server returned a reply: '%s'\n", resp.GetData())
	}
}
