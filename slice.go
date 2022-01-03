package rsize

import (
	"unsafe"
)

func eslice(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {

	subElemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	subElemKind := (*uint8)(unsafe.Add(subElemTypePtr, 2*word+7))

	sliceCounts := *(*int)(unsafe.Add(dataPtr, word*1)) // calculate len rather than cap
	if sliceCounts <= 0 {
		return 0
	}

	if size = originKind(*subElemKind); size != 0 {
		return int(sliceCounts) * size
	} else {
		// must ergodic all elements
		switch *subElemKind {
		case kindArray:
			for i := 0; i < sliceCounts; i++ {
				size = size + earray(dataPtr, subElemTypePtr)
			}
			return size
		case kindChan:
		case kindFunc:
		case kindInterface:
		case kindMap:
		case kindPtr:
		case kindSlice:
			subDataPtrs := *(*unsafe.Pointer)(dataPtr)
			for i := 0; i < sliceCounts; i++ {
				size = size + eslice(unsafe.Add(subDataPtrs, uintptr(i)*word*3), subElemTypePtr)
			}
			return size

		case kindString:
			sdataPtr := *(*unsafe.Pointer)(dataPtr) // slice having structure
			l := sliceCounts * 2
			for i := 0; i < l; i = i + 2 {
				size = size + *(*int)(unsafe.Add(sdataPtr, uintptr(i+1)*word))

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
