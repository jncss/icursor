package icursor

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
	// // create a new cursor
	// ic := NewICursor(customers, "Name, Id")

	// // Print Len
	// fmt.Println("Len:", ic.Len())

	// fmt.Println("-----")

	// // Print all customers
	// ic.SeekBeforeFirst()
	// for ic.Next() != -1 {
	// 	fmt.Println(customers[ic.Get()])
	// }

	// fmt.Println("-----")

	// // Seek Before "Sue"
	// if ic.SeekBefore([]any{"Sue", 0}) != -1 {
	// 	for ic.Next() != -1 {
	// 		fmt.Println(customers[ic.Get()])
	// 	}
	// }

	// Generate 1000 random customers, show time to sort
	fmt.Println("----- 1000000 random customers -----	")

	now := time.Now()

	var customers []Customer
	for i := 0; i < 1000000; i++ {
		customers = append(customers, Customer{i, RandString(10), RandString(10)})
	}

	// create a new cursor
	ic2 := NewICursor(customers, "Name")

	totalTime := time.Since(now)
	fmt.Println("Total time to sort 1000000 customers:", totalTime)

	// Print Len
	fmt.Println("Len:", ic2.Len())

	// Print 10 last customers
	ic2.SeekAfterLast()
	for i := 0; i < 10; i++ {
		fmt.Println(customers[ic2.Prev()])
	}

	// Create new slice of customers sorted by Name, Id
	var sortedCustomers []Customer

	ic2.SeekBeforeFirst()
	for ic2.Next() != -1 {
		sortedCustomers = append(sortedCustomers, customers[ic2.Get()])
	}

	// Check if sorted correctly
	for i := 0; i < len(sortedCustomers)-1; i++ {
		if sortedCustomers[i].Name > sortedCustomers[i+1].Name {
			t.Error("Customers not sorted correctly")
		}
	}

	// Check if Find works for all customers in sortedCustomers
	for _, v := range sortedCustomers {
		if ic2.Find([]any{v.Name}) == -1 {
			t.Error("Find failed")
		}
	}

	// Check if SeekBefore works for all customers in sortedCustomers
	for _, v := range sortedCustomers {
		ic2.SeekBefore([]any{v.Name})

		if customers[ic2.Next()].Name != v.Name {
			fmt.Println(customers[ic2.Get()].Name, v.Name)
			t.Error("SeekBefore failed")
			break
		}
	}

	// Check if SeekAfter works for all customers in sortedCustomers
	for _, v := range sortedCustomers {
		ic2.SeekAfter([]any{v.Name})

		if customers[ic2.Prev()].Name != v.Name {
			t.Error("SeekAfter failed")
		}
	}

	// Check if SeekBefore "AAAAAAAAAAAAAAAAAAA" must be -1
	if ic2.SeekBefore([]any{"AAAAAAAAAAAAAAAAAAA"}) != -1 {
		t.Error("SeekBefore failed")
	}

	// Check if SeekAfter "zzzzzzzzzzzzzzzzzzzzzz" then get prev must be the last customer
	ic2.SeekAfter([]any{"zzzzzzzzzzzzzzzzzzzzz"})
	if customers[ic2.Prev()].Name != sortedCustomers[len(sortedCustomers)-1].Name {
		fmt.Println(customers[ic2.Prev()].Name, sortedCustomers[len(sortedCustomers)-1].Name)
		t.Error("SeekAfter failed")
	}

	fmt.Println("-----")
}

// RandString returns a random string of length n
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[RandInt(0, len(letterRunes))]
	}
	return string(b)
}

// RandInt returns a random int between min and max
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
