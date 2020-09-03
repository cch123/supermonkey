# SuperMonkey

This lib is not production ready, please don't use

## Introduction

Patch all functions including which are unexported

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

//go:noinline
func heyHey() {
	fmt.Println("fuck")
}
```