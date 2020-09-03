package supermonkey

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

var (
	patchRecord = map[uintptr][]byte{}
	symbolTable = map[string]uintptr{}
)

// Patch patches a function
func Patch(pkgName, typeName, methodName string, patchFunc interface{}) {
	// find addr of the func
	symbolName := getSymbolName(pkgName, typeName, methodName)
	addr := symbolTable[symbolName]
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
		return "_" + pkgName + "." + "(" + typeName + ")" + methodName
	}

	return "_" + pkgName + "." + methodName
}

func init() {
	cmd := exec.Command("nm", os.Args[0])
	contentBytes, err := cmd.Output()
	if err != nil {
		fmt.Println("this lib depend on nm cmd, please install nm(binutils) first")
		os.Exit(1)
	}

	content := string(contentBytes)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
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

// from is a pointer to the actual function
// to is a pointer to a go funcvalue
func replaceFunction(from, to uintptr) (original []byte) {
	jumpData := jmpToFunctionValue(to)
	f := rawMemoryAccess(from, len(jumpData))
	original = make([]byte, len(f))
	copy(original, f)

	copyToLocation(from, jumpData)
	return
}

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0x48, 0xBA,
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24),
		byte(to >> 32),
		byte(to >> 40),
		byte(to >> 48),
		byte(to >> 56), // movabs rdx,to
		0xFF, 0x22,     // jmp QWORD PTR [rdx]
	}
}

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func mprotectCrossPage(addr uintptr, length int, prot int) {
	pageSize := syscall.Getpagesize()
	for p := pageStart(addr); p < addr+uintptr(length); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, pageSize)
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

// this function is super unsafe
// aww yeah
// It copies a slice to a raw memory location, disabling all memory protection before doing so.
func copyToLocation(location uintptr, data []byte) {
	f := rawMemoryAccess(location, len(data))

	mprotectCrossPage(location, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(f, data[:])
	mprotectCrossPage(location, len(data), syscall.PROT_READ|syscall.PROT_EXEC)
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
