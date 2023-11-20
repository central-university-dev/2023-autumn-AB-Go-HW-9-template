package main

import (
	"fmt"
	"grpc_workshop/2/types" // Import the generated protobuf code
)

// ProcessTestOneof processes the TestOneof field of ExampleTypes
func ProcessTestOneof(exampleType *types.ExampleTypes) {
	if exampleType == nil {
		fmt.Println("ExampleTypes is nil")
		return
	}

	switch v := exampleType.TestOneof.(type) {
	case *types.ExampleTypes_OneofString:
		fmt.Printf("String field: %s\n", v.OneofString)
	case *types.ExampleTypes_OneofInt:
		fmt.Printf("Int field: %d\n", v.OneofInt)
	default:
		fmt.Println("Unknown type or not set")
	}
}

func main() {
	// Create an instance of ExampleTypes
	exampleTypes := &types.ExampleTypes{}

	// Set the oneof field to a string and process it
	exampleTypes.TestOneof = &types.ExampleTypes_OneofString{OneofString: "Hello, Protobuf!"}
	ProcessTestOneof(exampleTypes)

	// Set the oneof field to an int and process it
	exampleTypes.TestOneof = &types.ExampleTypes_OneofInt{OneofInt: 123}
	ProcessTestOneof(exampleTypes)

	// Process with the oneof field not set
	exampleTypes.TestOneof = nil
	ProcessTestOneof(exampleTypes)
}
