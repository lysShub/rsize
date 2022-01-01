package rsize

import "unsafe"

const word = unsafe.Sizeof(uintptr(0)) // 位长的字节数
const typeOffsed = word*4 + 16

// copy form: src/runtime/typekind.go
const (
	kindBool = 1 + iota
	kindInt
	kindInt8
	kindInt16
	kindInt32
	kindInt64
	kindUint
	kindUint8
	kindUint16
	kindUint32
	kindUint64
	kindUintptr
	kindFloat32
	kindFloat64
	kindComplex64
	kindComplex128
	kindArray
	kindChan
	kindFunc
	kindInterface
	kindMap
	kindPtr
	kindSlice
	kindString
	kindStruct
	kindUnsafePointer

	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6
	kindMask        = (1 << 5) - 1
)

// copy form src/runtime/type.go
//  8 8 4 1 1 1 1 8 8 4 4
//
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

type arraytype struct {
	typ  _type
	elem *_type
	// slice *_type
	// len   uintptr
}

type tflag uint8
type nameOff int32
type typeOff int32

type myarraytype struct {
	typ  _type
	elem *_type
}
