package afloat

import (
	"math"
	"sync/atomic"
	"unsafe"
)

func Add32(addr *float32, value float32) (result float32) {
	for {
		old := *addr
		result = old + value

		if atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(addr)), math.Float32bits(old), math.Float32bits(result)) {
			break
		}
	}

	return
}
