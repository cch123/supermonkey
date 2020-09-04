package supermonkey

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/cch123/supermonkey/nm"
)

var (
	patchRecord = map[uintptr][]byte{}
	symbolTable = map[string]uintptr{}
)

// Patch patches a function
func Patch(pkgName, typeName, methodName string, patchFunc interface{}) {
	// find addr of the func
	symbolName := getSymbolName(pkgName, typeName, methodName)
	PatchByFullSymbolName(symbolName, patchFunc)
}

// PatchByFullSymbolName needs user to provide the full symbol path
func PatchByFullSymbolName(symbolName string, patchFunc interface{}) {
	addr := symbolTable[symbolName]
	if addr == 0 {
		fmt.Printf("The symbol is %v, and the patch target addr is 0, there may be 2 possible reasons\n", symbolName)
		fmt.Println("	1. the function is inlined, please add //go:noinline to function comment or add -l to gcflags")
		fmt.Println("	2. your input for symbolName or pkg/obj/method is wrong, check by using go tool nm {your_bin_file}")
		panic("")
	}
	originalBytes := replaceFunction(addr, (uintptr)(getPtr(reflect.ValueOf(patchFunc))))
	patchRecord[addr] = originalBytes
}

// UnpatchAll unpatches all functions
func UnpatchAll() {
	for funcAddr, funcBytes := range patchRecord {
		copyToLocation(funcAddr, funcBytes)
		delete(patchRecord, funcAddr)
	}
}

// return a arch dependent full symbol string
func getSymbolName(pkgName, typeName, methodName string) string {
	if typeName != "" {
		return pkgName + "." + "(" + typeName + ")" + "." + methodName
	}

	return pkgName + "." + methodName
}

func init() {
	content, _ := nm.Parse(os.Args[0])

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		arr := strings.Split(line, " ")
		if len(arr) < 3 {
			continue
		}

		funcSymbol, addr := arr[2], arr[0]
		addrUint, _ := strconv.ParseUint(addr, 16, 64)
		symbolTable[funcSymbol] = uintptr(addrUint)
	}
}

type value struct {
	_   uintptr
	ptr unsafe.Pointer
}

func getPtr(v reflect.Value) unsafe.Pointer {
	return (*value)(unsafe.Pointer(&v)).ptr
}
