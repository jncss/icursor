package icursor

import (
	"fmt"
	"reflect"
)

type icNode struct {
	key   any
	index int64
}

type iCursor struct {
	l []icNode
}

// New indexed cursor
func NewICursor(data any, key string) *iCursor {
	// check if data is a slice
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		panic("data must be a slice")
	}
	// check if key is a valid field name
	if reflect.ValueOf(data).Index(0).FieldByName(key).IsValid() == false {
		panic("key must be a valid field name")
	}
	// create a new cursor iterating over data
	c := new(iCursor)
	c.l = make([]icNode, reflect.ValueOf(data).Len())
	for i := 0; i < reflect.ValueOf(data).Len(); i++ {
		c.l[i].key = reflect.ValueOf(data).Index(i).FieldByName(key).Interface()
		c.l[i].index = int64(i)
	}

	return c
}

// Println
func (c *iCursor) Println() {
	for _, v := range c.l {
		fmt.Println(v.key, v.index)
	}
}
