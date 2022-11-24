package cgo

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#include "../rockxcrypto.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

var (
	emptyUint8 C.uint8_t = 0
)

func AsSliceRefUint8(goBytes []byte) SliceRefUint8 {
	len := len(goBytes)

	if len == 0 {
		// can't take element 0 of an empty slice
		return SliceRefUint8{
			ptr: &emptyUint8,
			len: C.size_t(len),
		}
	}
	return SliceRefUint8{
		ptr: (*C.uint8_t)(unsafe.Pointer(&goBytes[0])),
		len: C.size_t(len),
	}
}
