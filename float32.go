// Package afloat provides atomic operations for float types
package afloat

import (
	"math"
	"sync/atomic"
)

// noCopy is used to ensure that Float32 cannot be copied.
// If a copy of Float32 is made, the copy will not be usable.
// See https://github.com/golang/go/issues/8005#issuecomment-190753527 for more information.
type noCopy struct{}

// Lock is a no-op used to ensure that noCopy cannot be copied.
func (*noCopy) Lock() {}

// Float32 is an atomic wrapper around a float32
type Float32 struct {
	_     noCopy
	value uint32
}

// Add adds value to current value and returns the result.
// It is implemented using atomic.CompareAndSwapUint32.
// It is safe for concurrent use by multiple goroutines.
func (f *Float32) Add(value float32) float32 {
	for {
		current := math.Float32frombits(atomic.LoadUint32(&f.value))
		result := current + value

		if atomic.CompareAndSwapUint32(&f.value, math.Float32bits(current), math.Float32bits(result)) {
			return result
		}
	}
}

// Load returns the current value.
// It is safe for concurrent use by multiple goroutines.
func (f *Float32) Load() float32 {
	return math.Float32frombits(atomic.LoadUint32(&f.value))
}
