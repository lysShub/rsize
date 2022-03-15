package rsize

import (
	"unsafe"
)

func earray(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {

	subElemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	subElemKind := (*(*uint8)(unsafe.Add(subElemTypePtr, 2*word+7))) & kindMask

	arrayLens := *(*int)(unsafe.Add(typePtr, typeOffsed+word*2))
	if arrayLens <= 0 {
		return 0
	}
	if size = originKind(subElemKind); size != 0 {
		size = int(arrayLens) * size
	} else {
		// any type array, element has equal length
		step := (*(*int)(unsafe.Pointer(typePtr))) / arrayLens

		switch subElemKind {
		case kindArray:
			for i := 0; i < arrayLens; i++ {
				size = size + earray(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindChan:
			for i := 0; i < arrayLens; i++ {
				size = size + echan(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindFunc:
		case kindInterface:
			for i := 0; i < arrayLens; i++ {
				size = size + eslice(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindMap:
			size = emaps(dataPtr, typePtr)
		case kindPtr:
			for i := 0; i < arrayLens; i++ {
				size = size + eptrs(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindSlice:
			for i := 0; i < arrayLens; i++ {
				size = size + eslice(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindString:
			for i := 0; i < arrayLens; i++ {
				size = size + *(*int)(unsafe.Add(dataPtr, word+uintptr(i)*word*2))
			}
		case kindStruct:
			for i := 0; i < arrayLens; i++ {
				size = size + estruct(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		case kindUnsafePointer:
			for i := 0; i < arrayLens; i++ {
				size = size + eunptr(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
		default:
		}
	}

	return size
}

type arraytype struct {
	typ   _type
	elem  *_type
	slice *_type
	len   uintptr
}

type arraytype1 struct {
	typ   _type
	elem  *_type
	slice *_type
	len   uintptr
}

type stringStruct struct {
	str unsafe.Pointer
	len int
}
