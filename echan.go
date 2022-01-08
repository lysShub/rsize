package rsize

import (
	"fmt"
	"unsafe"
)

type chantype struct {
	typ  _type
	elem *_type
	dir  uintptr
}

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	lock mutex
}
type waitq struct {
	first uintptr
	last  uintptr
}
type mutex struct {
	lockRankStruct
	key uintptr
}
type lockRankStruct struct {
}

func echan(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {

	a := *(*hchan)(dataPtr)

	fmt.Println(a)

	elemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	elemKind := *(*uint8)(unsafe.Add(elemTypePtr, 2*word+7))

	counts := *(*uint)(unsafe.Add(dataPtr, 0))
	if size = originKind(elemKind); size != 0 {
		size = size * int(counts)
	} else {
		start := *(*uint)(unsafe.Add(dataPtr, word*5+8))      // recvx
		step := uint(*(*uint16)(unsafe.Add(dataPtr, word*3))) // elemsize

		subDataPtr := *(*unsafe.Pointer)(unsafe.Add(dataPtr, word*2))
		switch elemKind {
		case kindArray:
			for i := uint(0); i < counts; i++ {
				size = size + earray(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindChan:
			for i := uint(0); i < counts; i++ {
				size = size + echan(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindFunc:
		case kindInterface:
			for i := uint(0); i < counts; i++ {
				size = size + eface(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindMap:
			for i := uint(0); i < counts; i++ {
				size = size + emaps(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindPtr:
		case kindSlice:
			for i := uint(0); i < counts; i++ {
				size = size + eslice(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindString:
			for i := uint(0); i < counts; i++ {
				size = size + *(*int)(unsafe.Add(subDataPtr, uintptr((start+i)*step)+word))
			}
		case kindStruct:
			for i := uint(0); i < counts; i++ {
				size = size + estruct(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		case kindUnsafePointer:
		default:
		}

	}

	return size
}
