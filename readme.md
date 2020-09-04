# SuperMonkey

This lib is not production ready, please don't use

## Introduction

Patch all functions including which are unexported

**Warning** : please add -l to your gcflags, or it will panic.

### patch private function

```go
package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func main() {
	fmt.Println("original function output:")
	heyHey() // fuck
	fmt.Println()

	sm.Patch("main", "", "heyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHey() // please be polite
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	heyHey() // fuck
}

func heyHey() {
	fmt.Println("fuck")
}
```

> go run -gcflags="-l" yourfile.go

### patch private instance method

```go
package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

type person struct{ name string }

func (p *person) speak() {
	fmt.Println("my name is ", p.name)
}

func main() {
	var p = person{"Xargin"}
	fmt.Println("original function output:")
	p.speak()
	fmt.Println()

	sm.Patch("main", "*person", "speak", func() {
		fmt.Println("we are all the same")
	})
	fmt.Println("after patch, function output:")
	p.speak()
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	p.speak()
}

```

> go run -gcflags="-l" yourfile.go
