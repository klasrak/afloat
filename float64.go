package afloat

import (
	"math"
	"sync/atomic"
)

// Float64 is an atomic wrapper around a float64
type Float64 struct {
	_     noCopy
	value uint64
}

// Add adds value to current value and returns the result.
// It is implemented using atomic.CompareAndSwapUint64
// and is safe for concurrent use by multiple goroutines.
func (f *Float64) Add(value float64) float64 {
	for {
		current := f.Load()
		result := current + value

		if f.CompareAndSwap(current, result) {
			return result
		}
	}
}

// Load returns the current value.
// It is implemented using atomic.LoadUint64
// and is safe for concurrent use by multiple goroutines.
func (f *Float64) Load() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.value))
}

// Store sets the new value.
// It is implemented using atomic.StoreUint64
// and is safe for concurrent use by multiple goroutines.
func (f *Float64) Store(value float64) {
	atomic.StoreUint64(&f.value, math.Float64bits(value))
}

// Swap sets the new value and returns the old value.
// It is implemented using atomic.SwapUint64
// and is safe for concurrent use by multiple goroutines.
func (f *Float64) Swap(value float64) float64 {
	return math.Float64frombits(atomic.SwapUint64(&f.value, math.Float64bits(value)))
}

// CompareAndSwap executes the compare-and-swap operation for the value.
// It is implemented using atomic.CompareAndSwapUint64
// and is safe for concurrent use by multiple goroutines.
func (f *Float64) CompareAndSwap(current, new float64) bool {
	return atomic.CompareAndSwapUint64(&f.value, math.Float64bits(current), math.Float64bits(new))
}
