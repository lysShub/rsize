package rsize

import (
	"fmt"
	"unsafe"
)

func eslice(slicePtr *unsafe.Pointer, elemTypePtr *unsafe.Pointer) (size int) {
	sliceElemKind := (*uint8)(unsafe.Pointer(uintptr(*elemTypePtr) + 2*word + 7))
	sliceCounts := (*uintptr)(unsafe.Pointer(uintptr(*slicePtr) + word*1)) // calculate len rather than cap

	if size = originKind(*sliceElemKind); size != 0 {
		return int(*sliceCounts) * size
	} else {
		// must ergodic all elements
		switch *sliceElemKind {
		case kindArray:

		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:
			subElemTypePtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*elemTypePtr) + 4*word + 16))

			subElemKind := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*subElemTypePtr) + 2*word + 7))
			fmt.Println(*subElemKind)

			sliceDataPtr := (*unsafe.Pointer)(unsafe.Pointer(uintptr(*slicePtr) + word*0))
			for i := uintptr(0); i < *sliceCounts; i = i + word {
				size = size + eslice((*unsafe.Pointer)(unsafe.Add(*sliceDataPtr, word*i)), subElemTypePtr)
			}
			return size
		case kindString:
			for i := uintptr(1); i < *sliceCounts; i = i + 2 {
				size = size + *(*int)(unsafe.Pointer(uintptr(*slicePtr) + i*word))
			}
			return size
		case kindStruct:

		case kindUnsafePointer:
		default:
		}
	}

	return
}

type slicetype struct {
	typ  _type
	elem *slicetype
}
