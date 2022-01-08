package rsize

import (
	"fmt"
	"unsafe"
)

type maptype struct {
	typ    _type
	key    *_type
	elem   *arraytype
	bucket *_type // internal type representing a hash bucket
	// function for hashing keys (ptr to key, seed) -> hash
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8  // size of key slot
	elemsize   uint8  // size of elem slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

func emaps(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {
	a := (*maptype)(typePtr)
	fmt.Println(a)

	subKeyTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	subKeyKind := (*uint8)(unsafe.Add(subKeyTypePtr, 2*word+7))

	subValueTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed+word))
	subValueKind := (*uint8)(unsafe.Add(subValueTypePtr, 2*word+7))

	var count = *(*int)(unsafe.Add(dataPtr, 0))
	var ksize, vsize int

	if ksize, vsize = originKind(*subKeyKind)*count, originKind(*subValueKind)*count; ksize != 0 && vsize != 0 {
		return (ksize + vsize)
	} else {
		var kstep, vstep uintptr = uintptr(*(*uint8)(unsafe.Add(typePtr, typeOffsed+4*word))), uintptr(*(*uint8)(unsafe.Add(typePtr, typeOffsed+4*word+1)))
		var bucketPtrs = *(*unsafe.Pointer)(unsafe.Add(dataPtr, word+8))
		var buckets int = 1 << (*(*uint8)(unsafe.Add(dataPtr, word+1)))

		for i := 0; i < buckets; i++ {
			bucketPtr := unsafe.Add(bucketPtrs, word*uintptr(i))
			if uintptr(bucketPtr) == 0 {
				continue // bucket has no value
			}
			for i := uintptr(0); i < 8; i++ {
				topHash := *(*uint8)(bucketPtr)
				if topHash != 0 { // point has value

					switch *subKeyKind {
					case kindArray:
						vsize = vsize + earray(unsafe.Add(bucketPtr, 8*(kstep+1)+vstep*i), subValueTypePtr)
					case kindChan:
					case kindFunc:
					case kindInterface:
					case kindMap:
					case kindPtr:
					case kindSlice:
						vsize = vsize + eslice(unsafe.Add(bucketPtr, 8*(kstep+1)+vstep*i), subValueTypePtr)
					case kindString:
						vsize = vsize + *(*int)(unsafe.Add(bucketPtr, (8*(kstep+1)+vstep*i)+word))
					case kindStruct:
					case kindUnsafePointer:
					default:
					}

					// ----------------------- //

					switch *subValueKind {
					case kindArray:
						vsize = vsize + earray(unsafe.Add(bucketPtr, 8*(kstep+1)+vstep*i), subValueTypePtr)
					case kindChan:
					case kindFunc:
					case kindInterface:
					case kindMap:
					case kindPtr:
					case kindSlice:
						vsize = vsize + eslice(unsafe.Add(bucketPtr, 8*(kstep+1)+vstep*i), subValueTypePtr)
					case kindString:
						vsize = vsize + *(*int)(unsafe.Add(bucketPtr, (8*(kstep+1)+vstep*i)+word))
					case kindStruct:
					case kindUnsafePointer:
					default:
					}
				}
			}
		}

		return ksize + vsize

		fmt.Println(kstep, vstep)

	}

	return
}
