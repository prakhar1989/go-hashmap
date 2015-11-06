package hashmap

import "errors"

// Node which is stored at each level
type Node struct {
	key   string
	Value interface{}
}

// HashMap implemented with a fixed size.
// Uses chaining to resolve collisions.
type HashMap struct {
	size    int
	count   int
	buckets [][]Node
}

/** PRIVATE METHODS **/

// returns the index at which the key needs to go
func (h *HashMap) getIndex(key string) int {
	return int(hash(key)) % h.size
}

// Implements the Jenkins hash function
func hash(key string) uint32 {
	var h uint32
	for _, c := range key {
		h += uint32(c)
		h += (h << 10)
		h ^= (h >> 6)
	}
	h += (h << 3)
	h ^= (h >> 11)
	h += (h << 15)
	return h
}

/** PUBLIC METHODS **/

// Len returns the count of the elements in the hashmap
func (h *HashMap) Len() int {
	return h.count
}

// Size returns the size of the hashamp
func (h *HashMap) Size() int {
	return h.size
}

// NewHashMap is the constuctor that returns a new HashMap of a fixed size
// returns an error when a size of 0 is provided
func NewHashMap(size int) (*HashMap, error) {
	h := new(HashMap)
	if size < 1 {
		return h, errors.New("size of hashmap has to be > 1")
	}
	h.size = size
	h.count = 0
	h.buckets = make([][]Node, size)
	for i := range h.buckets {
		h.buckets[i] = make([]Node, 0)
	}
	return h, nil
}

// Get returns the value associated with a key in the hashmap,
// and an error indicating whether the value exists
func (h *HashMap) Get(key string) (*Node, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]
	for _, node := range chain {
		if node.key == key {
			return &node, true
		}
	}
	return nil, false
}

// Set the value for an associated key in the hashmap
func (h *HashMap) Set(key string, value interface{}) bool {
	index := h.getIndex(key)
	chain := h.buckets[index]
	found := false

	// first see if the key already exists
	for i := range chain {
		// if found, update the value
		node := &chain[i]
		if node.key == key {
			node.Value = value
			found = true
		}
	}
	if found { // hashmap has been updated
		return true
	}

	// if key doesn't exist, add it to the hashmap
	// first check whether space exists
	if h.size == h.count {
		return false
	}

	// yup there's space, let's add a new node
	node := Node{key: key, Value: value}
	chain = append(chain, node)
	h.buckets[index] = chain
	h.count++

	return true
}

// Delete the value associated with key in the hashmap
func (h *HashMap) Delete(key string) (*Node, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]

	found := false
	var location int
	var mapNode *Node

	// start a search for the key
	for loc, node := range chain {
		if node.key == key {
			found = true
			location = loc
			mapNode = &node
		}
	}

	// if found delete the elem from the slice
	if found {
		h.count--
		N := len(chain) - 1
		chain[location], chain[N] = chain[N], chain[location]
		chain = chain[:N]
		h.buckets[index] = chain
		return mapNode, true
	}

	// if not found return false
	return nil, false
}

// Load returns the load factor of the hashmap
func (h *HashMap) Load() float32 {
	return float32(h.count) / float32(h.size)
}
