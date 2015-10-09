### Hashmap

A naive implementation of a Hash map in [Golang](https://golang.org/).

### Background
I'm primarily a Python developer by trade and most of my experience is with dynamic languages. So my first attempt at solving this problem was an implementation in Python. However, since the question was a way to make your application stand-out, I thought it would be fun and rewarding to build solve this problem in a (relatively new) statically compiled language. Since Go has been gaining momentum rapidly I decided to build this up in Go. Another key reason why I decided to use a new language was to demonstrate that I'm a quick learner. Since most startups are typically looking for generalists who can learn quickly, I thought it would be a good way to show my enthusiasm and passion for hi-tech startups.  The code that you might see in the subsequent files might not be most idiomatic Go but it works and (hopefully) fulfills the specification.

### About Me
I'm a CS masters student in Columbia University where I'm specializing in Machine Learning. With around 3-4 years of experience as a full stack developer I've worked with varities of technologies e.g. Rails / Django / Flask on the backend & React / Angular / Backbone on the frontend. More importantly, I'm extremely passionate about open-source, writing and doing side-projects.

To know more about me, please visit the following links:

- Github: [http://github.com/prakhar1989](http://github.com/prakhar1989)
- Portfolio: [http://prakhar.me/projects](http://prakhar.me/projects)
- Blog: [http://prakhar.me/articles](http://prakhar.me/articles)

Outside of programming, reading about startups and business is one of my key interests. Having spent inordinate amount of time lurking on [HN](https://news.ycombinator.com/user?id=krat0sprakhar) and reading Paul Graham's essays, I've come to believe that working at a startup right out of school is one of the best way to kickstart your career. Hence, getting an opportuinity to work with one of the portfolio companies of KPCB would be super fun.

### Compile
In order to compile the program for your environment, you need to have the Go compiler installed. The compiler is distributed as a binary and can be downloaded from the [website](https://golang.org/dl/). Once the compiler is ready, just run the command below to download the code as a package.

```
$ go get github.com/prakhar1989/go-hashmap
```
With the dependancies in place, you can now copy/paste the code below and run it with. Alternatively, you can run the example code in the `example` directory.
```
$ go run main.go
```

### Structure

The source code comes in just two files - the main library and an accompanying test suite. The `hashmap.go` file contains the implementation of the hashmap specification and the `hashmap_test.go` has the testcases for the implementation.

```
$ hashmap git:(master) go test
PASS
ok      github.com/prakhar1989/hashmap0.006s
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/prakhar1989/hashmap"
)

func main() {
    h, _ := hashmap.NewHashMap(100) // create the hashmap
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
    fmt.Println(h.Delete("alpha"))
    _, present := h.Get("alpha")
    if present {
        fmt.Println("The key's still there")
    } else {
        fmt.Println("Value associated with alpha deleted")
    }
}
```

Here's the output of running this program

```
$ go run example/main.go
The load factor is: 0.25
Key: alpha -> Value: 5
Key: beta -> Value: 4
Key: charlie -> Value: 7
Key: gamma -> Value: 5
Key: delta -> Value: 5
<nil> false
Value associated with alpha deleted
```

### Disclaimer
Although the implementation tries to stay as close to the spec as it can, in a few places it does behave differently. This is has primarily been done in order to stay idiomatic with the language. For example, instead of returning a `nil` or a value, the `Delete` instead returns a value and a boolean flag indicating the success and failure of the operation. Another example is that the constructor throws up an error when initialized with a non-positive value of size.
