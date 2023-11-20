package main

import (
	"fmt"
	"grpc_workshop/2/types" // Import the generated protobuf code
)

func main() {
	// Create an instance of ExampleTypes and populate it
	exampleTypes := &types.ExampleTypes{
		IntField:         42,
		FloatField:       3.14,
		DoubleField:      2.71828,
		StringField:      "Hello, Protobuf!",
		BoolField:        true,
		BytesField:       []byte("Raw bytes"),
		EnumField:        types.ExampleTypes_OPTION_ONE,
		RepeatedIntField: []int32{1, 2, 3},
		MapField:         map[string]int32{"key1": 100, "key2": 200},
		NestedMessage: &types.ExampleTypes_NestedMessage{
			NestedField: "Nested content",
		},
	}

	// Using oneof field
	// Only one of these will be set at a time
	exampleTypes.TestOneof = &types.ExampleTypes_OneofString{OneofString: "I am a string"}
	// exampleTypes.TestOneof = &example.ExampleTypes_OneofInt{OneofInt: 1234}

	// Print the content
	fmt.Printf("ExampleTypes: %+v\n", exampleTypes)
}
