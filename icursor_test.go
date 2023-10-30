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
	cr := NewICursor(customers, "Name, Id")

	// Print Len
	fmt.Println("Len:", cr.Len())

	fmt.Println("-----")

	// Print all customers
	cr.SeekBeforeFirst()
	for cr.Next() != -1 {
		fmt.Println(customers[cr.Get()])
	}

	fmt.Println("-----")

	// Seek Before "Sue"
	if cr.SeekBefore([]any{"Sue", 0}) != -1 {
		for cr.Next() != -1 {
			fmt.Println(customers[cr.Get()])
		}
	}
}
