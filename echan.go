package rsize

import (
	"fmt"
	"unsafe"
)

func echan(dataPtr unsafe.Pointer, typePtr unsafe.Pointer) (size int) {
	a := (*hchan)(dataPtr)
	fmt.Println(a)
	return

	elemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	elemKind := (*(*uint8)(unsafe.Add(elemTypePtr, 2*word+7))) & kindMask

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
			for i := uint(0); i < counts; i++ {
				size = size + eptrs(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
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
			for i := uint(0); i < counts; i++ {
				size = size + eunptr(unsafe.Add(subDataPtr, (start+i)*step), elemTypePtr)
			}
		default:
		}

	}

	return size
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

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

type waitq struct {
	first uint64
	last  uint64
}

type mutex struct {
	key uintptr
}
