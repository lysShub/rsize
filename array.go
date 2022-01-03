package rsize

import (
	"unsafe"
)

func earray(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {
	subElemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	subElemKind := (*uint8)(unsafe.Add(subElemTypePtr, 2*word+7))

	arrayLens := *(*int)(unsafe.Add(typePtr, typeOffsed+word*2))
	if arrayLens <= 0 {
		return 0
	}
	if size = originKind(*subElemKind); size != 0 {
		return int(arrayLens)
	} else {
		// must ergodic all elements
		switch *subElemKind {
		case kindArray:
			// any array nest, element has equal length
			step := (*(*int)(unsafe.Pointer(typePtr))) / arrayLens
			for i := 0; i < arrayLens; i++ {
				size = size + earray(unsafe.Add(dataPtr, uintptr(i*step)), subElemTypePtr)
			}
			return size
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:

		case kindString:
			for i := 0; i < arrayLens; i++ {
				size = size + *(*int)(unsafe.Add(dataPtr, word+uintptr(i)*word*2))
			}
			return size
		case kindStruct:

		case kindUnsafePointer:
		default:
		}

	}

	return
}

type arraytype struct {
	typ   _type
	elem  *arraytype1
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
