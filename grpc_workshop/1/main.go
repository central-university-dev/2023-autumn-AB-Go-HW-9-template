package main

import (
	"encoding/json"
	"fmt"
	"grpc_workshop/1/example"
	"log"

	"google.golang.org/protobuf/proto"
)

type JSONUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a user instance for JSON
	jsonUser := JSONUser{ID: 123, Name: "John Doe", Email: "john.doe@types.com"}

	// Serialize to JSON
	jsonBytes, err := json.Marshal(jsonUser)
	if err != nil {
		log.Fatal("JSON Marshal error:", err)
	}
	fmt.Printf("JSON size: %d bytes\n", len(jsonBytes))

	// Create a user instance for Protobuf
	protoUser := &example.User{Id: 123, Name: "John Doe", Email: "john.doe@types.com"}

	// Serialize to Protobuf
	protoBytes, err := proto.Marshal(protoUser)
	if err != nil {
		log.Fatal("Protobuf Marshal error:", err)
	}
	fmt.Printf("Protobuf size: %d bytes\n", len(protoBytes))
}
