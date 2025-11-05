package main

import (
	"fmt"
)

type kv struct {
	key   string
	value interface{}
}

type HashTable struct {
	data [][]kv
	size int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([][]kv, size),
		size: size,
	}
}

func (h *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % h.size
	}
	// fmt.Println("hash ->", hash)
	// fmt.Println("key ->", key)
	// fmt.Println("size ->", h.size)
	return hash
}

func (h *HashTable) Set(key string, value interface{}) {
	addr := h.hash(key)
	if h.data[addr] == nil {
		h.data[addr] = []kv{}
	}
	h.data[addr] = append(h.data[addr], kv{key: key, value: value})
}

func (h *HashTable) Get(key string) (interface{}, bool) {
	addr := h.hash(key)
	bucket := h.data[addr]
	for i := range bucket {
		if bucket[i].key == key {
			return bucket[i].value, true
		}
	}
	return nil, false
}

func main() {
	myHashTable := NewHashTable(50)

	myHashTable.Set("grapes", 10000)
	if v, ok := myHashTable.Get("grapes"); ok {
		fmt.Println("grapes ->", v)
	} else {
		fmt.Println("grapes not found")
	}

	myHashTable.Set("apples", 9)
	if v, ok := myHashTable.Get("apples"); ok {
		fmt.Println("apples ->", v)
	} else {
		fmt.Println("apples not found")
	}
}
