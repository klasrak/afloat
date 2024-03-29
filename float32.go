// Package afloat provides atomic operations for float types
package afloat

import (
	"math"
	"sync/atomic"
)

// Float32 is an atomic wrapper around a float32
type Float32 struct {
	_     noCopy
	value uint32
}

// Add adds value to current value and returns the result.
// It is implemented using atomic.CompareAndSwapUint32
// and is safe for concurrent use by multiple goroutines.
func (f *Float32) Add(value float32) float32 {
	for {
		current := f.Load()
		result := current + value

		if f.CompareAndSwap(current, result) {
			return result
		}
	}
}

// Load returns the current value.
// It is implemented using atomic.LoadUint32
// and is safe for concurrent use by multiple goroutines.
func (f *Float32) Load() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.value))
}

// Store sets the new value.
// It is implemented using atomic.StoreUint32
// and is safe for concurrent use by multiple goroutines.
func (f *Float32) Store(value float32) {
	atomic.StoreUint32(&f.value, math.Float32bits(value))
}

// Swap sets the new value and returns the old value.
// It is implemented using atomic.SwapUint32
// and is safe for concurrent use by multiple goroutines.
func (f *Float32) Swap(value float32) float32 {
	return math.Float32frombits(atomic.SwapUint32(&f.value, math.Float32bits(value)))
}

// CompareAndSwap executes the compare-and-swap operation for the value.
// It is implemented using atomic.CompareAndSwapUint32
// and is safe for concurrent use by multiple goroutines.
func (f *Float32) CompareAndSwap(current, new float32) bool {
	return atomic.CompareAndSwapUint32(&f.value, math.Float32bits(current), math.Float32bits(new))
}
