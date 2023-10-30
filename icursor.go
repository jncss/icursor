package icursor

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type icNode struct {
	key   []any
	index int64
}

type iCursor struct {
	l    []icNode
	keys []string
	desc []bool
	curr int64
}

// New indexed cursor
func NewICursor(data any, keys string) *iCursor {
	// check if data is an array or slice
	if reflect.TypeOf(data).Kind() != reflect.Slice && reflect.TypeOf(data).Kind() != reflect.Array {
		panic("data must be an array or slice")
	}
	// check if elements are structs
	if reflect.TypeOf(data).Elem().Kind() != reflect.Struct {
		panic("data elements must be structs")
	}
	// create a new cursor and fill it with the data
	c := new(iCursor)
	// keys are comma separated field names
	for _, v := range strings.Split(keys, ",") {
		v = strings.TrimSpace(v)
		// if key ends with a minus sign, it is a descending key
		if strings.HasSuffix(v, "-") {
			v = strings.TrimSuffix(v, "-")
			c.keys = append(c.keys, v)
			c.desc = append(c.desc, true)
		} else {
			c.keys = append(c.keys, v)
			c.desc = append(c.desc, false)
		}
	}

	// check if keys are a valid field name
	for _, v := range c.keys {
		if !reflect.ValueOf(data).Index(0).FieldByName(v).IsValid() {
			panic("key " + v + " is not a valid field name")
		}
	}

	c.l = make([]icNode, reflect.ValueOf(data).Len())
	for i := 0; i < reflect.ValueOf(data).Len(); i++ {
		c.l[i].key = make([]any, len(c.keys))
		for j, v := range c.keys {
			c.l[i].key[j] = reflect.ValueOf(data).Index(i).FieldByName(v).Interface()
		}
		c.l[i].index = int64(i)
	}

	// sort the cursor
	sort.Slice(c.l, func(i, j int) bool {
		for k := 0; k < len(c.keys); k++ {
			if c.l[i].key[k] != c.l[j].key[k] {
				switch c.l[i].key[k].(type) {
				case int:
					if c.desc[k] {
						return c.l[i].key[k].(int) > c.l[j].key[k].(int)
					}
					return c.l[i].key[k].(int) < c.l[j].key[k].(int)
				case int8:
					if c.desc[k] {
						return c.l[i].key[k].(int8) > c.l[j].key[k].(int8)
					}
					return c.l[i].key[k].(int8) < c.l[j].key[k].(int8)
				case int16:
					if c.desc[k] {
						return c.l[i].key[k].(int16) > c.l[j].key[k].(int16)
					}
					return c.l[i].key[k].(int16) < c.l[j].key[k].(int16)
				case int32:
					if c.desc[k] {
						return c.l[i].key[k].(int32) > c.l[j].key[k].(int32)
					}
					return c.l[i].key[k].(int32) < c.l[j].key[k].(int32)
				case int64:
					if c.desc[k] {
						return c.l[i].key[k].(int64) > c.l[j].key[k].(int64)
					}
					return c.l[i].key[k].(int64) < c.l[j].key[k].(int64)
				case uint:
					if c.desc[k] {
						return c.l[i].key[k].(uint) > c.l[j].key[k].(uint)
					}
					return c.l[i].key[k].(uint) < c.l[j].key[k].(uint)
				case uint8:
					if c.desc[k] {
						return c.l[i].key[k].(uint8) > c.l[j].key[k].(uint8)
					}
					return c.l[i].key[k].(uint8) < c.l[j].key[k].(uint8)
				case uint16:
					if c.desc[k] {
						return c.l[i].key[k].(uint16) > c.l[j].key[k].(uint16)
					}
					return c.l[i].key[k].(uint16) < c.l[j].key[k].(uint16)
				case uint32:
					if c.desc[k] {
						return c.l[i].key[k].(uint32) > c.l[j].key[k].(uint32)
					}
					return c.l[i].key[k].(uint32) < c.l[j].key[k].(uint32)
				case uint64:
					if c.desc[k] {
						return c.l[i].key[k].(uint64) > c.l[j].key[k].(uint64)
					}
					return c.l[i].key[k].(uint64) < c.l[j].key[k].(uint64)
				case float32:
					if c.desc[k] {
						return c.l[i].key[k].(float32) > c.l[j].key[k].(float32)
					}
					return c.l[i].key[k].(float32) < c.l[j].key[k].(float32)
				case float64:
					if c.desc[k] {
						return c.l[i].key[k].(float64) > c.l[j].key[k].(float64)
					}
					return c.l[i].key[k].(float64) < c.l[j].key[k].(float64)
				case string:
					if c.desc[k] {
						return c.l[i].key[k].(string) > c.l[j].key[k].(string)
					}
					return c.l[i].key[k].(string) < c.l[j].key[k].(string)
				}
			}
		}
		return false
	})

	return c
}

// Len returns the number of elements in the cursor
func (c *iCursor) Len() int {
	return len(c.l)
}

// Println prints the cursor
func (c *iCursor) Println() {
	for _, v := range c.l {
		fmt.Println(v.key, v.index)
	}
}

// Find returns the index of the first element that matches the key
func (c *iCursor) Find(key []any) int64 {
	for _, v := range c.l {
		if reflect.DeepEqual(v.key, key) {
			c.curr = v.index
			return v.index
		}
	}
	return -1
}

// Next returns the next element in the cursor
func (c *iCursor) Next() int64 {
	if c.curr == int64(len(c.l)-1) {
		return -1
	}
	c.curr++
	return c.l[c.curr].index
}

// Prev returns the previous element in the cursor
func (c *iCursor) Prev() int64 {
	if c.curr == 0 {
		return -1
	}
	c.curr--
	return c.l[c.curr].index
}

// First returns the first element in the cursor
func (c *iCursor) First() int64 {
	c.curr = 0
	return c.l[c.curr].index
}

// Last returns the last element in the cursor
func (c *iCursor) Last() int64 {
	c.curr = int64(len(c.l) - 1)
	return c.l[c.curr].index
}

// Get returns the current element in the cursor
func (c *iCursor) Get() int64 {
	return c.l[c.curr].index
}
