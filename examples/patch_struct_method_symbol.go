package main

import (
	"context"
	"fmt"

	sm "github.com/cch123/supermonkey"
)

func patchStructMethodSymbol() {
	f := &Foo{}
	fmt.Println("original function output:")
	fmt.Println(f.MyFunc(nil))

	//patchGuard := sm.PatchByFullSymbolName("main.(*Foo).MyFunc", func(_ *Foo, ctx context.Context) (*Bar, error) {
	patchGuard := sm.PatchByFullSymbolName("main.(*Foo).MyFunc", func(_ uintptr, ctx context.Context) (*Bar, error) {
		return &Bar{"Not bar"}, nil
	})

	fmt.Println("after patch, function output:")
	fmt.Println(f.MyFunc(nil))

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	fmt.Println(f.MyFunc(nil))

	patchGuard.Restore()
	fmt.Println("restore, then output:")
	fmt.Println(f.MyFunc(nil))

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	fmt.Println(f.MyFunc(nil))

	fmt.Println()
}
