package rsize

import (
	"fmt"
	"unsafe"
)

func init() {
	if err := Check(); err != nil {
		panic(err)
	}
}

// build on go version go1.17.3

type eface1 struct {
	_type *_type
	data  unsafe.Pointer
}

func Size(e interface{}) (size int) {

	a := (*eface1)(unsafe.Pointer(&e))
	fmt.Println(a)

	// efaceTypePtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*0))
	efaceTypePtr := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&e), word*0))
	// efaceDataPtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*1))
	efaceDataPtr := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&e), word*1))

	return eface(efaceDataPtr, efaceTypePtr)
}

func Check() (err error) {

	return
}

type efacetype struct {
	_type *_type
	data  unsafe.Pointer
}

/*

switch kind {
case kindArray:
case kindChan:
case kindFunc:
case kindInterface:
case kindMap:
case kindPtr:
case kindSlice:
case kindString:
case kindStruct:
case kindUnsafePointer:
default:
}

*/

// originKind
// 	如果类型是纯纯的“值类型”, 就返回此类型的大小, 否则返回0
func originKind(kind uint8) (size int) {
	switch kind {
	case kindBool, kindInt8, kindUint8:
		return 1
	case kindInt16, kindUint16:
		return 2
	case kindInt32, kindUint32, kindFloat32:
		return 4
	case kindInt64, kindUint64, kindFloat64, kindComplex64, kindFunc:
		return 8
	case kindComplex128:
		return 16
	case kindInt, kindUint:
		return int(word)
	default:
		return 0
	}
}
