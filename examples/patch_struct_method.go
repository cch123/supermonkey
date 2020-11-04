package main

import (
	"context"
	"fmt"

	sm "github.com/cch123/supermonkey"
)

type Bar struct {
	Name string
}

type Foo struct{}

//go:noinline
func (*Foo) MyFunc(ctx context.Context) (*Bar, error) {
	return &Bar{"Bar"}, nil
}

func patchStructMethod() {
	f := &Foo{}
	fmt.Println("original function output:")
	fmt.Println(f.MyFunc(nil))

	patchGuard := sm.Patch((*Foo).MyFunc, func(_ *Foo, ctx context.Context) (*Bar, error) {
		return &Bar{"Not bar"}, nil
	})

	fmt.Println("after patch, function output:")
	fmt.Println(f.MyFunc(nil))

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	fmt.Println(f.MyFunc(nil))

	fmt.Println()
}