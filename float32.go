// Package afloat provides atomic operations for float types
package afloat

import (
	"math"
	"sync/atomic"
	"unsafe"
)

// Add32 adds value to *addr and returns the result.
// It is implemented using atomic.CompareAndSwapUint32.
// It is safe for concurrent use by multiple goroutines.
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

// Load32 returns the value at *addr.
// It is implemented using atomic.LoadUint32.
// It is safe for concurrent use by multiple goroutines.
func Load32(addr *float32) (result float32) {
	return math.Float32frombits(atomic.LoadUint32((*uint32)(unsafe.Pointer(addr))))
}
