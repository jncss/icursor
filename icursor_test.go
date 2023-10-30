package icursor

import (
	"testing"
)

type Customer struct {
	Id    int
	Name  string
	Phone string
}

var customers = []Customer{
	{Id: 1, Name: "John", Phone: "555-1212"},
	{Id: 2, Name: "Mary", Phone: "555-1213"},
	{Id: 3, Name: "Bill", Phone: "555-1214"},
	{Id: 4, Name: "Jill", Phone: "555-1215"},
	{Id: 5, Name: "Jack", Phone: "555-1216"},
	{Id: 6, Name: "Jen", Phone: "555-1217"},
	{Id: 7, Name: "Jone", Phone: "555-1218"},
	{Id: 8, Name: "Jill1", Phone: "555-1219"},
	{Id: 9, Name: "Jill2", Phone: "555-1220"},
	{Id: 10, Name: "Jill", Phone: "555-1221"},
	{Id: 11, Name: "Jill", Phone: "555-1222"},
	{Id: 12, Name: "Jill6", Phone: "555-1223"},
	{Id: 13, Name: "Jill7", Phone: "555-1224"},
	{Id: 14, Name: "Jill8", Phone: "555-1225"},
}

// test
func TestICursor(t *testing.T) {
	// create a new cursor
	cursor := NewICursor(customers, "Name")

	// print the list
	cursor.Println()
}
