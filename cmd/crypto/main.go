package main

// #cgo linux LDFLAGS: ${SRCDIR}/../../librockxcrypto.a -Wl,-unresolved-symbols=ignore-all
// #cgo darwin LDFLAGS: ${SRCDIR}/../../librockxcrypto.a -Wl,-undefined,dynamic_lookup
import "C"

import (
	"fmt"
	"go-ffi/cgo"

	"github.com/ChainSafe/go-schnorrkel"
)

func main() {
	println("hello word!")
	sdr25519expandfromseed()
}

func sdr25519expandfromseed() {
	mnemonicSeed := "person dry champion eye bridge envelope copy off isolate biology method banner"
	msc, err := schnorrkel.MiniSecretFromMnemonic(mnemonicSeed, "")
	if err != nil {
		println("schnorrkel MiniSecretFromMnemonic err", err)
		return
	}
	seedFixed := msc.Encode()
	seed := seedFixed[:]
	res := cgo.Sr25519ExpandFromSeed(cgo.AsSliceRefUint8(seed))

	fmt.Printf("r25519ExpandFromSeed res%v \n", res)
}
