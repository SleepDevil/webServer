package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func test() {
	// Creating UUID Version 4
	u1 := uuid.NewV4()

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
