package main

import (
	"fmt"
	"time"

	"github.com/google/btree"
)

// Descriptions struct represents the data structure to be sorted.
type Descriptions struct {
	Description string
	Time        time.Time
}

// LessThan defines the less than function for Descriptions.
func LessThan(a, b Descriptions) bool {
	return a.Time.Before(b.Time)
}

func main() {
	// Create a new B-tree with degree 4 and custom less than function.
	tree := btree.NewG(4, LessThan)

	// Example data
	data := []Descriptions{
		{Description: "Third", Time: time.Now().Add(3 * time.Hour)},
		{Description: "Second", Time: time.Now().Add(2 * time.Hour)},
		{Description: "First", Time: time.Now().Add(1 * time.Hour)},
	}

	// Insert data into the tree
	for _, d := range data {
		tree.ReplaceOrInsert(d)
	}

	// Traverse the tree in ascending order (sorted order)
	tree.Ascend(func(item Descriptions) bool {
		fmt.Printf("Description: %s, Time: %s\n", item.Description, item.Time.Format(time.RFC3339))
		return true
	})
}
