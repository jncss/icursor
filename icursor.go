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

type icKey struct {
	name string
	kind reflect.Kind
	desc bool
}
type iCursor struct {
	l    []icNode
	keys []icKey
	curr int64
}

// New indexed cursor
func NewICursor(data any, keys string) *iCursor {
	// check if data is an array or slice
	if reflect.TypeOf(data).Kind() != reflect.Slice && reflect.TypeOf(data).Kind() != reflect.Array {
		panic("Data must be an array or slice")
	}

	// check if elements are structs
	if reflect.TypeOf(data).Elem().Kind() != reflect.Struct {
		panic("Data elements must be structs")
	}

	// create a new cursor and fill it with the data
	c := new(iCursor)
	c.curr = -1
	// keys are comma separated field names
	for _, v := range strings.Split(keys, ",") {
		var desc bool
		v = strings.TrimSpace(v)
		// if key ends with a minus sign, it is a descending key
		if strings.HasSuffix(v, "-") {
			v = strings.TrimSuffix(v, "-")
			desc = true
		}
		// check if key is a valid field name
		if !reflect.ValueOf(data).Index(0).FieldByName(v).IsValid() {
			panic("Key " + v + " is not a valid field name")
		}
		c.keys = append(c.keys,
			icKey{
				name: v,
				kind: reflect.ValueOf(data).Index(0).FieldByName(v).Kind(),
				desc: desc,
			})

	}

	// fill the cursor
	c.l = make([]icNode, reflect.ValueOf(data).Len())
	for i := 0; i < reflect.ValueOf(data).Len(); i++ {
		c.l[i].key = make([]any, len(c.keys))
		for j, v := range c.keys {
			c.l[i].key[j] = reflect.ValueOf(data).Index(i).FieldByName(v.name).Interface()
		}
		c.l[i].index = int64(i)
	}

	// sort the cursor
	sort.Slice(c.l, func(i, j int) bool {
		return c.less(c.l[i].key, c.l[j].key)
	})

	return c
}

// Compare two keys and return -1, 0, 1
func (c *iCursor) compare(key1, key2 []any) int {
	for k := 0; k < len(c.keys); k++ {
		if key1[k] != key2[k] {
			switch key1[k].(type) {
			case int:
				if c.keys[k].desc {
					if key1[k].(int) > key2[k].(int) {
						return -1
					}
					if key1[k].(int) < key2[k].(int) {
						return 1
					}
				} else {
					if key1[k].(int) < key2[k].(int) {
						return -1
					}
					if key1[k].(int) > key2[k].(int) {
						return 1
					}
				}
			case int8:
				if c.keys[k].desc {
					if key1[k].(int8) > key2[k].(int8) {
						return -1
					}
					if key1[k].(int8) < key2[k].(int8) {
						return 1
					}
				} else {
					if key1[k].(int8) < key2[k].(int8) {
						return -1
					}
					if key1[k].(int8) > key2[k].(int8) {
						return 1
					}
				}
			case int16:
				if c.keys[k].desc {
					if key1[k].(int16) > key2[k].(int16) {
						return -1
					}
					if key1[k].(int16) < key2[k].(int16) {
						return 1
					}
				} else {
					if key1[k].(int16) < key2[k].(int16) {
						return -1
					}
					if key1[k].(int16) > key2[k].(int16) {
						return 1
					}
				}
			case int32:
				if c.keys[k].desc {
					if key1[k].(int32) > key2[k].(int32) {
						return -1
					}
					if key1[k].(int32) < key2[k].(int32) {
						return 1
					}
				} else {
					if key1[k].(int32) < key2[k].(int32) {
						return -1
					}
					if key1[k].(int32) > key2[k].(int32) {
						return 1
					}
				}
			case int64:
				if c.keys[k].desc {
					if key1[k].(int64) > key2[k].(int64) {
						return -1
					}
					if key1[k].(int64) < key2[k].(int64) {
						return 1
					}
				} else {
					if key1[k].(int64) < key2[k].(int64) {
						return -1
					}
					if key1[k].(int64) > key2[k].(int64) {
						return 1
					}
				}
			case uint:
				if c.keys[k].desc {
					if key1[k].(uint) > key2[k].(uint) {
						return -1
					}
					if key1[k].(uint) < key2[k].(uint) {
						return 1
					}
				} else {
					if key1[k].(uint) < key2[k].(uint) {
						return -1
					}
					if key1[k].(uint) > key2[k].(uint) {
						return 1
					}
				}
			case uint8:
				if c.keys[k].desc {
					if key1[k].(uint8) > key2[k].(uint8) {
						return -1
					}
					if key1[k].(uint8) < key2[k].(uint8) {
						return 1
					}
				} else {
					if key1[k].(uint8) < key2[k].(uint8) {
						return -1
					}
					if key1[k].(uint8) > key2[k].(uint8) {
						return 1
					}
				}
			case uint16:
				if c.keys[k].desc {
					if key1[k].(uint16) > key2[k].(uint16) {
						return -1
					}
					if key1[k].(uint16) < key2[k].(uint16) {
						return 1
					}
				} else {
					if key1[k].(uint16) < key2[k].(uint16) {
						return -1
					}
					if key1[k].(uint16) > key2[k].(uint16) {
						return 1
					}
				}
			case uint32:
				if c.keys[k].desc {
					if key1[k].(uint32) > key2[k].(uint32) {
						return -1
					}
					if key1[k].(uint32) < key2[k].(uint32) {
						return 1
					}
				} else {
					if key1[k].(uint32) < key2[k].(uint32) {
						return -1
					}
					if key1[k].(uint32) > key2[k].(uint32) {
						return 1
					}
				}
			case uint64:
				if c.keys[k].desc {
					if key1[k].(uint64) > key2[k].(uint64) {
						return -1
					}
					if key1[k].(uint64) < key2[k].(uint64) {
						return 1
					}
				} else {
					if key1[k].(uint64) < key2[k].(uint64) {
						return -1
					}
					if key1[k].(uint64) > key2[k].(uint64) {
						return 1
					}
				}
			case float32:
				if c.keys[k].desc {
					if key1[k].(float32) > key2[k].(float32) {
						return -1
					}
					if key1[k].(float32) < key2[k].(float32) {
						return 1
					}
				} else {
					if key1[k].(float32) < key2[k].(float32) {
						return -1
					}
					if key1[k].(float32) > key2[k].(float32) {
						return 1
					}
				}
			case float64:
				if c.keys[k].desc {
					if key1[k].(float64) > key2[k].(float64) {
						return -1
					}
					if key1[k].(float64) < key2[k].(float64) {
						return 1
					}
				} else {
					if key1[k].(float64) < key2[k].(float64) {
						return -1
					}
					if key1[k].(float64) > key2[k].(float64) {
						return 1
					}
				}
			case string:
				if c.keys[k].desc {
					if key1[k].(string) > key2[k].(string) {
						return -1
					}
					if key1[k].(string) < key2[k].(string) {
						return 1
					}
				} else {
					if key1[k].(string) < key2[k].(string) {
						return -1
					}
					if key1[k].(string) > key2[k].(string) {
						return 1
					}
				}
			}
		}
	}
	return 0
}

// less returns true if key1 < key2
func (c *iCursor) less(key1, key2 []any) bool {
	return c.compare(key1, key2) == -1
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
	// binary search
	lo := 0
	hi := len(c.l) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if c.compare(c.l[mid].key, key) == -1 {
			lo = mid + 1
		} else if c.compare(c.l[mid].key, key) == 1 {
			hi = mid - 1
		} else {
			c.curr = int64(mid)
			return c.l[mid].index
		}
	}
	c.curr = -1
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
	if c.curr == -1 {
		return -1
	}
	return c.l[c.curr].index
}

// SeekBefore returns the index of the first element that is less than the key
func (c *iCursor) SeekBefore(key []any) int64 {
	// binary search
	lo := 0
	hi := len(c.l) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if c.compare(c.l[mid].key, key) == -1 {
			lo = mid + 1
		} else if c.compare(c.l[mid].key, key) == 1 {
			hi = mid - 1
		} else {
			c.curr = int64(mid)
			return c.l[mid].index
		}
	}
	if lo == 0 {
		c.curr = -1
		return -1
	}
	c.curr = int64(lo - 1)
	return c.l[lo-1].index
}

// SeekAfter returns the index of the first element that is greater than the key
func (c *iCursor) SeekAfter(key []any) int64 {
	// binary search
	lo := 0
	hi := len(c.l) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if c.compare(c.l[mid].key, key) == -1 {
			lo = mid + 1
		} else if c.compare(c.l[mid].key, key) == 1 {
			hi = mid - 1
		} else {
			c.curr = int64(mid)
			return c.l[mid].index
		}
	}
	if lo == len(c.l) {
		c.curr = -1
		return -1
	}
	c.curr = int64(lo)
	return c.l[lo].index
}

// SeekBeforeFirst sets the cursor before the first element
func (c *iCursor) SeekBeforeFirst() int64 {
	c.curr = -1
	return -1
}

// SeekAfterLast sets the cursor after the last element
func (c *iCursor) SeekAfterLast() int64 {
	c.curr = int64(len(c.l))
	return -1
}

// Seek returns the index of the first element that matches the key
func (c *iCursor) Seek(key []any) int64 {
	return c.Find(key)
}
