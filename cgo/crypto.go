package cgo

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lrockxcrypto
#include "../rockxcrypto.h"
#include <stdlib.h>
*/
import "C"

func Sr25519ExpandFromSeed(seed SliceRefUint8) []byte {
	resp := C.sr25519_expand_from_seed(seed)
	defer resp.destroy()
	return resp.slice()
}
