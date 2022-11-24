package cgo

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#include "../rockxcrypto.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ByteArray64 = C.uint8_64_array_t
type SliceRefUint8 = C.slice_ref_uint8_t

func (ptr *ByteArray64) destroy() {
	if ptr != nil {
		C.destroy_sr25519_expand_private_key(ptr)
		ptr = nil
	}
}

func (ptr ByteArray64) slice() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&ptr.idx[0])), 64)
}
