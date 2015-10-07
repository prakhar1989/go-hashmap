package hashmap

type HashMapNode struct {
	key   string
	value int // this should be a pointer to any type
}

// has an array of HashMapNodes
type HashMap struct {
	size    int
	count   int
	buckets [][]HashMapNode
}

/* METHODS */

// Constuctor that returns a new HashMap
func NewHashMap(size int) *HashMap {
	h := new(HashMap)
	h.size = size
	h.count = 0
	h.buckets = make([][]HashMapNode, size)
	for i := range h.buckets {
		h.buckets[i] = make([]HashMapNode, 1)
	}
	return h
}

// returns the index at which the key needs to go
// TODO: Do we need to typecast again?
func (h *HashMap) getIndex(key string) int {
	return int(hash(key)) % h.size
}

// gets the value associated with a key in the hashmap
// returns the value to be true / false depending on whether
// the key exists in the hashmap
func (h *HashMap) get(key string) (int, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]
	for _, node := range chain {
		if node.key == key {
			return node.value, true
		}
	}
	return 0, false
}

// sets the value for an associated key in the hashmap
func (h *HashMap) set(key string, value int) bool {
	index := h.getIndex(key)
	chain := h.buckets[index]
	found := false
	for _, node := range chain {
		// if the key already exists
		if node.key == key {
			node.value = value
			found = true
		}
	}
	if found { // hashmap has been updated
		return true
	}
	if h.size == h.count { // hashmap is full
		return false
	}
	// add a new node
	node := HashMapNode{key: key, value: value}
	chain = append(chain, node)
	h.buckets[index] = chain
	h.count += 1
	return true
}

func (h *HashMap) delete(key string) (int, bool) {
	index := h.getIndex(key)
	chain := h.buckets[index]

	// start a search for the key
	found := false
	var location, value int
	for loc, node := range chain {
		if node.key == key {
			found = true
			location = loc
			value = node.value
		}
	}

	// if found delete the elem from the slice
	if found {
		h.count -= 1
		N := len(chain) - 1
		chain[location], chain[N] = chain[N], chain[location]
		chain = chain[:N]
		h.buckets[index] = chain
		return value, true
	}

	// if not found return false
	return 0, false
}

// returns the load factor of the hashmap
func (h *HashMap) load() float32 {
	return float32(h.count) / float32(h.size)
}

// Implements the Jenkins hash function
// TODO: What should the return type be?
func hash(key string) uint32 {
	var h uint32 = 0
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
