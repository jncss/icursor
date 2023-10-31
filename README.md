# icursor
Indexed cursor to search data or iterate, in an ordered manner, over slice/array of golang struct's. It allows
 multiple keys, indexed ascending or descending.


## Install:
```
go get github.com/jncss/icursor
```
## Example:
```golang
package main

import (
	"fmt"

	"github.com/jncss/icursor"
)

type Customer struct {
	ID    int
	Name  string
	Age   int
	Phone string
}

var customers = []Customer{
	{ID: 1, Name: "John", Age: 20, Phone: "12345678"},
	{ID: 2, Name: "Peter", Age: 30, Phone: "12345679"},
	{ID: 3, Name: "Mark", Age: 25, Phone: "12345670"},
	{ID: 4, Name: "Paul", Age: 40, Phone: "12345671"},
	{ID: 5, Name: "Simon", Age: 50, Phone: "12345672"},
	{ID: 6, Name: "James", Age: 60, Phone: "12345673"},
	{ID: 7, Name: "Luke", Age: 70, Phone: "12345674"},
	{ID: 8, Name: "Judas", Age: 80, Phone: "12345675"},
	{ID: 9, Name: "Thomas", Age: 90, Phone: "12345676"},
	{ID: 10, Name: "Matthew", Age: 100, Phone: "12345677"},
	{ID: 11, Name: "Andrew", Age: 11, Phone: "12345678"},
	{ID: 12, Name: "Philip", Age: 12, Phone: "12345679"},
	{ID: 13, Name: "Bartholomew", Age: 13, Phone: "12345670"},
	{ID: 14, Name: "James", Age: 14, Phone: "12345671"},
	{ID: 15, Name: "Simon", Age: 15, Phone: "12345672"},
	{ID: 16, Name: "Thaddaeus", Age: 16, Phone: "12345673"},
}

func main() {
	// Create a new icursor
	// Sort by Name (ascending) and Age (descending)
	ic := icursor.New(customers, "Name, Age-")

	// Print all customers
	ic.SeekBeforeFirst()
	for ic.Next() != -1 {
		c := customers[ic.Get()]
		fmt.Printf("%d\t%-20s\t%d\t%s\n", c.ID, c.Name, c.Age, c.Phone)
	}

	fmt.Println("----------")

	// Search for key "James" age 14
	if ic.Find([]any{"James", 14}) != -1 {
		c := customers[ic.Get()]
		fmt.Printf("%d\t%-20s\t%d\t%s\n", c.ID, c.Name, c.Age, c.Phone)
	} else {
		fmt.Println("Not found")
	}
}
```
### Example output:
```
11      Andrew                  11      12345678
13      Bartholomew             13      12345670
6       James                   60      12345673
14      James                   14      12345671
1       John                    20      12345678
8       Judas                   80      12345675
7       Luke                    70      12345674
3       Mark                    25      12345670
10      Matthew                 100     12345677
4       Paul                    40      12345671
2       Peter                   30      12345679
12      Philip                  12      12345679
5       Simon                   50      12345672
15      Simon                   15      12345672
16      Thaddaeus               16      12345673
9       Thomas                  90      12345676
----------
14      James                   14      12345671
```

### API:
```golang
// New indexed cursor. "keys" are comma separated field names. 
// If a key ends with a minus sign, it is a descending key.
New(data any, keys string) *iCursor

// Len returns the number of elements in the cursor
(c *iCursor) Len() int

// Find returns the index of the first element that matches the key
(c *iCursor) Find(key []any) int64

// Next returns the next element in the cursor
(c *iCursor) Next() int64

// Prev returns the previous element in the cursor
(c *iCursor) Prev() int64

// First returns the first element in the cursor
(c *iCursor) First() int64

// Last returns the last element in the cursor
(c *iCursor) Last() int64

// Get returns the current element in the cursor
(c *iCursor) Get() int64

// SeekBefore returns the possition of the first element that is less than the key
func (c *iCursor) SeekBefore(key []any) int64

// SeekAfter returns the index of the first element that is greater than the key
func (c *iCursor) SeekAfter(key []any) int64

// SeekBeforeFirst sets the cursor before the first element
func (c *iCursor) SeekBeforeFirst() int64

// SeekAfterLast sets the cursor after the last element
func (c *iCursor) SeekAfterLast() int64

// Seek returns the index of the first element that matches the key
func (c *iCursor) Seek(key []any) int64
```