package rsize

import (
	"unsafe"
)

// build on go version go1.17.3

func Size(efacePtr interface{}) (size int) {

	// efaceTypePtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*0))
	efaceTypePtr := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&efacePtr), word*0))
	// efaceDataPtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(efacePtr)) + word*1))
	efaceDataPtr := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&efacePtr), word*1))

	return eface(efaceDataPtr, efaceTypePtr)
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
