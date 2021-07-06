package main

import (
	"fmt"
	"unsafe"

	sm "github.com/cch123/supermonkey"
)

func patchInstanceFuncSymbol() {
	p := &person{"Linda"}
	fmt.Println("original function output:")
	p.speak()

	patchGuard := sm.PatchByFullSymbolName("main.(*person).speak", func(ptr uintptr) {
		p = (*person)(unsafe.Pointer(ptr))
		fmt.Println(p.name, ", we are all the same")
	})
	fmt.Println("after patch, function output:")
	p.speak()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	p.speak()

	patchGuard.Restore()
	fmt.Println("restore, then output:")
	p.speak()

	patchGuard.Unpatch()
	fmt.Println("unpatch, then output:")
	p.speak()

	fmt.Println()
}
