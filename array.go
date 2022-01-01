package rsize

import "unsafe"

func earray(arrPtr *unsafe.Pointer, elemTypePtr *unsafe.Pointer) (size int) {
	arrayElemKind := (*uint8)(unsafe.Pointer(uintptr(*elemTypePtr) + (2*word + 7)))
	arrayCounts := (*uintptr)(unsafe.Pointer(uintptr(*elemTypePtr) + 0))
	if size = originKind(*arrayElemKind); size != 0 {
		return int(*arrayCounts)
	} else {
		// must ergodic all elements
		switch *arrayElemKind {
		case kindArray:
			// [2][2]int equal [1][4]int equal [4]int
			subElemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*elemTypePtr) + word*4 + 16))
			return earray(arrPtr, subElemTypePtr)
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:

		case kindString:
			for i := uintptr(1); i < *arrayCounts; i = i + 2 {
				size = size + *(*int)(unsafe.Pointer(uintptr(*arrPtr) + i*word))
			}
			return size
		case kindStruct:

		case kindUnsafePointer:
		default:
		}

	}

	return
}

func earray2(dataPtr *unsafe.Pointer, typePtr *unsafe.Pointer) (size int) {

	subElemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*typePtr) + typeOffsed))
	subElemKind := (*uint8)(unsafe.Pointer(uintptr(*subElemTypePtr) + (2*word + 7)))

	arrayCounts := (*uintptr)(unsafe.Pointer(uintptr(*typePtr) + 0))

	if size = originKind(*subElemKind); size != 0 {
		return int(*arrayCounts)
	} else {
		// must ergodic all elements
		switch *subElemKind {
		case kindArray:
			// [2][2]int equal [1][4]int equal [4]int
			subElemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*subElemTypePtr) + word*4 + 16))
			return earray(dataPtr, subElemTypePtr)
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:

		case kindString:
			for i := uintptr(1); i < *arrayCounts; i = i + 2 {
				size = size + *(*int)(unsafe.Pointer(uintptr(*dataPtr) + i*word))
			}
			return size
		case kindStruct:

		case kindUnsafePointer:
		default:
		}

	}

	return
}
