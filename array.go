package rsize

import "unsafe"

func earray(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {
	subElemTypePtr := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(typePtr) + typeOffsed))
	subElemKind := (*uint8)(unsafe.Pointer(uintptr(subElemTypePtr) + (2*word + 7)))
	arraySizes := (*uintptr)(unsafe.Pointer(uintptr(typePtr) + 0))

	if size = originKind(*subElemKind); size != 0 {
		return int(*arraySizes)
	} else {
		// must ergodic all elements
		switch *subElemKind {
		case kindArray:
			// [2][2]int equal [1][4]int equal [4]int
			return earray(dataPtr, subElemTypePtr)
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:

		case kindString:
			for i := uintptr(1); i < *arraySizes; i = i + 2 {
				size = size + *(*int)(unsafe.Pointer(uintptr(dataPtr) + i*word))
			}
			return size
		case kindStruct:

		case kindUnsafePointer:
		default:
		}

	}

	return
}
