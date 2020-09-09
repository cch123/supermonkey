# SuperMonkey

This lib is inspired by https://github.com/bouk/monkey, and uses some of the code

## Introduction

Patch all functions without limits, including which are unexported

**Warning** : please add -l to your gcflags or add `//go:noinline` to func which you want to patch.


## when running in tests

you should run this lib under a go mod project and provide the full project path

**Warning** : use `go test -ldflags="-s=false" -gcflags="-l"` to enable symbol table and disable inline.

## when running not in tests

### patch private function

#### normal

```go
package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func main() {
	fmt.Println("original function output:")
	heyHey() // fake
	fmt.Println()

	sm.Patch("main", "", "heyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHey() // please be polite
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	heyHey() // fake
}

func heyHey() {
	fmt.Println("fake")
}
```

> go run -gcflags="-l" yourfile.go

#### full symbol name

```go
package main

import (
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func main() {
	fmt.Println("original function output:")
	heyHey()
	fmt.Println()

	sm.PatchByFullSymbolName("main.heyHey", func() {
		fmt.Println("please be polite")
	})
	fmt.Println("after patch, function output:")
	heyHey()
	fmt.Println()

	sm.UnpatchAll()
	fmt.Println("unpatch all, then output:")
	heyHey()
}

//go:noinline
func heyHey() {
	fmt.Println("fake")
}

```

> go run -gcflags="-l" yourfile.go

### patch private instance method

#### normal

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

#### full symbol name

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

	sm.PatchByFullSymbolName("main.(*person).speak", func() {
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
