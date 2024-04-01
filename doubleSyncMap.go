package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a double nested sync.Map
	nestedMap := sync.Map{}

	// Define keys for the outer map
	outerKeys := []string{"outer1", "outer2", "outer3"}

	// Populate the nested maps
	for _, outerKey := range outerKeys {
		nestedMap.Store(outerKey, &sync.Map{})
	}

	// Define some key-value pairs for the inner maps
	innerKeyValuePairs := map[string][]string{
		"outer1": {"inner1", "inner2", "inner3"},
		"outer2": {"inner4", "inner5"},
		"outer3": {"inner6"},
	}

	// Populate the inner maps
	for outerKey, innerKeys := range innerKeyValuePairs {
		innerMap, _ := nestedMap.Load(outerKey)
		for _, innerKey := range innerKeys {
			innerMap.(*sync.Map).Store(innerKey, "some value")
		}
	}

	// Access and print the values
	nestedMap.Range(func(outerKey, innerMap interface{}) bool {
		fmt.Println("Outer Key:", outerKey)
		innerMap.(*sync.Map).Range(func(innerKey, value interface{}) bool {
			fmt.Println("    Inner Key:", innerKey, "Value:", value)
			return true
		})
		return true
	})
}
