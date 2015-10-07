### Hashmap

A naive implementation of a Hash map in Golang.

Nothing interesting to see here... move along. :runner:

### Usage

```go
package main

import (
    "fmt"
    "github.com/prakhar1989/hashmap"
)

func main() {
    h := hashmap.NewHashMap(100) // create the hashmap
    keys := []string{"alpha", "beta", "charlie", "gamma", "delta"}

    // add the keys
    for _, key := range keys {
        h.Set(key, len(key))
    }

    fmt.Println("The load factor is:", h.Load())
}
```
