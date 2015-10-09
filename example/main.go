package main

import (
	"fmt"
	"github.com/prakhar1989/hashmap"
)

func main() {
	h, _ := hashmap.NewHashMap(20)

	keys := []string{"alpha", "beta", "charlie", "gamma", "delta"}

	// add the keys
	for _, key := range keys {
		h.Set(key, len(key))
	}

	fmt.Println("The load factor is:", h.Load())

	// retrieve the keys
	for _, key := range keys {
		val, present := h.Get(key)
		if present {
			fmt.Println("Key:", key, "->", "Value:", val.Value.(int))
		} else {
			fmt.Println(key, "is not present")
		}
	}

	// delete a key
	node, _ := h.Delete("alpha")
	fmt.Println(node.Value.(int), "deleted")
	val, present := h.Get("alpha")
	if present {
		fmt.Println("alpha ->", val)
	} else {
		fmt.Println("alpha is not present")
	}

	// delete a key
	fmt.Println(h.Delete("alpha"))
	_, present = h.Get("alpha")
	if present {
		fmt.Println("The key's still there")
	} else {
		fmt.Println("Value associated with alpha deleted")
	}
}
