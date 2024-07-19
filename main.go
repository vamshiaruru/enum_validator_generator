package main

import (
	"log"

	"github.com/vamshiaruru/enum_validator_generator/generator"
)

func main() {
	err := generator.Run()
	if err != nil {
		log.Fatalf("Error running generator: %v", err)
	}
}
