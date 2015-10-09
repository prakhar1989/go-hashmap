package main

import (
	//	"fmt"
	"github.com/prakhar1989/hashmap"
)

func main() {
	h, _ := hashmap.NewHashMap(100) // create the hashmap
	h.Set("alpha", 100)
	h.Set("alpha", 200)

	//node, _ := h.Get("alpha")
	//	fmt.Println(node)
	//	fmt.Println(node.Value.(int))

	/*
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
		for _, key := range keys {
			val, present := h.Get(key)
			if present {
				fmt.Println("Key:", key, "->", "Value:", val.Value.(int))
			} else {
				fmt.Println(key, "is not present")
			}
		}
	*/
}
