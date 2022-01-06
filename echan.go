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

	/*
		if qcount >0 {
			if recvx > sendx {
				next = recvx + 1 // 满的buf读取了一部分，还没有读取完成时
			} else if recvx == sendx {
				next = 0 // buff刚好写满, 此时两者值为0
			} else if recvx < sendx {
				next = 0 // e没有写满buff
			}
		}
	*/

	elemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	elemKind := *(*uint8)(unsafe.Add(elemTypePtr, 2*word+7))

	counts := *(*int)(unsafe.Add(dataPtr, 0))

	// elemTypePtr := *(*unsafe.Pointer)(unsafe.Add(typePtr, typeOffsed))
	fmt.Println(elemKind, counts)

	if size = originKind(elemKind); size != 0 {
		return size * counts
	} else {

	}

	return
}
