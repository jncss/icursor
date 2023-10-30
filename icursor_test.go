package icursor

import (
	"fmt"
	"testing"
)

type Customer struct {
	Id    int
	Name  string
	Phone string
}

// 20 Random customers
var customers = []Customer{
	{1, "John", "555-1234"},
	{2, "Sue", "555-2345"},
	{3, "Bob", "555-3456"},
	{4, "Mel", "555-4567"},
	{5, "Jen", "555-5678"},
	{6, "Sue", "555-6789"},
	{7, "Ken", "555-7890"},
	{8, "Dave", "555-8901"},
	{9, "Beth", "555-9012"},
	{10, "Ray", "555-0001"},
	{11, "Sam", "555-0002"},
	{12, "Dan", "555-0003"},
	{13, "Sue", "555-0004"},
	{14, "Mike", "555-0005"},
	{15, "Bob", "555-0006"},
	{16, "Andy", "555-0007"},
	{17, "Joe", "555-0008"},
	{18, "Sue", "555-0009"},
	{19, "Phil", "555-0010"},
	{20, "Mary", "555-0011"},
}

// test
func TestICursor(t *testing.T) {
	// create a new cursor
	cr := NewICursor(customers, "Name-, Id")

	// print the list
	cr.Println()

	// Find key "Sue, 13"
	idx := cr.Find([]any{"Sue", 13})
	if idx == -1 {
		t.Error("Find failed")
	} else {
		fmt.Println("Found Sue, 13 at index", idx)
	}

	// First element
	idx = cr.First()
	if idx == -1 {
		t.Error("First failed")
	} else {
		fmt.Println("First element is", customers[idx])
	}

	// Iterate through the list
	for idx = cr.First(); idx != -1; idx = cr.Next() {
		fmt.Println(customers[idx])
	}

	// Last element
	idx = cr.Last()
	if idx == -1 {
		t.Error("Last failed")
	} else {
		fmt.Println("Last element is", customers[idx])
	}

	// Iterate through the list backwards
	for idx = cr.Last(); idx != -1; idx = cr.Prev() {
		fmt.Println(customers[idx])
	}
}
