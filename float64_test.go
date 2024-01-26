package afloat

import (
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	MIN_F64       = math.SmallestNonzeroFloat64
	MAX_F64       = 10.0
	TOLERANCE_F64 = 1e-5 // 10^-5
)

func TestAddF64(t *testing.T) {
	type testCase struct {
		name         string
		value        float64
		delta        float64
		maxAdditions int
	}

	testCases := []testCase{
		{
			name:         "Add 1.0 to 0.0",
			value:        0.0,
			delta:        1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add 1.0 to 1.0",
			value:        1.0,
			delta:        1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add random value to random value",
			value:        getRandomFloat64(),
			delta:        getRandomFloat64(),
			maxAdditions: 1,
		},
		{
			name:         "Add random value to random value 100 times",
			value:        getRandomFloat64(),
			delta:        getRandomFloat64(),
			maxAdditions: 100,
		},
		{
			name:         "Add -1.0 to 1.0",
			value:        1.0,
			delta:        -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add -1.0 to -1.0",
			value:        -1.0,
			delta:        -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add 1.0 to -1.0",
			value:        -1.0,
			delta:        1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add negative random value to positive random value",
			value:        getRandomFloat64(),
			delta:        getRandomFloat64() * -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add negative random value to positive random value 100 times",
			value:        getRandomFloat64(),
			delta:        getRandomFloat64() * -1.0,
			maxAdditions: 100,
		},
		{
			name:         "Add positive random value to negative random value",
			value:        getRandomFloat64() * -1.0,
			delta:        getRandomFloat64(),
			maxAdditions: 1,
		},
		{
			name:         "Add positive random value to negative random value 100 times",
			value:        getRandomFloat64() * -1.0,
			delta:        getRandomFloat64(),
			maxAdditions: 100,
		},
		{
			name:         "Add negative random value to negative random value",
			value:        getRandomFloat64() * -1.0,
			delta:        getRandomFloat64() * -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add negative random value to negative random value 100 times",
			value:        getRandomFloat64() * -1.0,
			delta:        getRandomFloat64() * -1.0,
			maxAdditions: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				v  Float64
				wg sync.WaitGroup
			)

			v.Store(tc.value)

			expected := tc.value + (tc.delta * float64(tc.maxAdditions))

			for i := 0; i < tc.maxAdditions; i++ {
				wg.Add(1)

				go func() {
					defer wg.Done()
					v.Add(tc.delta)
				}()
			}

			wg.Wait()

			if !float64Equals(expected, v.Load()) {
				t.Errorf("Expected %.18f, got %.18f", expected, v.Load())
			}
		})
	}
}

func TestLoad64(t *testing.T) {
	type testCase struct {
		name     string
		maxLoads int
	}

	testCases := []testCase{
		{
			name:     "Load correct value",
			maxLoads: 1,
		},
		{
			name:     "Load correct value 100 concurrent accesses",
			maxLoads: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				v               Float64
				wg              sync.WaitGroup
				currentExpected chan float64 = make(chan float64, 1)
			)

			v.Store(getRandomFloat64())

			go func() {
				for {
					time.Sleep(1 * time.Millisecond)
					currentExpected <- v.Add(getRandomFloat64())
				}
			}()

			testLoad := func() {
				// Simulate concurrent access
				expected := <-currentExpected // Get the current expected value
				result := v.Load()            // Get the current value

				// Check if the result is the expected value
				if expected != result {
					t.Errorf("Expected %.18f, got %.18f", expected, result)
				}
			}

			for i := 0; i < tc.maxLoads; i++ {
				wg.Add(1)

				go func() {
					defer wg.Done()
					testLoad()
				}()
			}

			wg.Wait()
		})

	}
}

func TestStoreF64(t *testing.T) {
	type testCase struct {
		name     string
		maxLoads int
	}

	testCases := []testCase{
		{
			name:     "Store correct value",
			maxLoads: 1,
		},
		{
			name:     "Store correct value 100 concurrent accesses",
			maxLoads: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var (
				v               Float64
				wg              sync.WaitGroup
				currentExpected chan float64 = make(chan float64, 1)
			)

			go func() {
				for {
					time.Sleep(1 * time.Millisecond)
					currentExpected <- getRandomFloat64()
				}
			}()

			testStore := func() {
				// Simulate concurrent access
				expected := <-currentExpected // Get the current expected value
				v.Store(expected)             // Set the current value

				// Check if the result is the expected value
				if expected != v.Load() {
					t.Errorf("Expected %.18f, got %.18f", expected, v.Load())
				}
			}

			for i := 0; i < tc.maxLoads; i++ {
				wg.Add(1)

				go func() {
					defer wg.Done()
					testStore()
				}()
			}

			wg.Wait()
		})
	}
}

func TestSwapF64(t *testing.T) {
	var v Float64
	v.Store(getRandomFloat64())

	expected := v.Load()
	got := v.Swap(getRandomFloat64())

	if expected != got {
		t.Errorf("Expected %.18f, got %.18f", expected, got)
	}
}

func TestCompareAndSwapF64(t *testing.T) {
	type testCase struct {
		name     string
		current  float64
		new      float64
		expected bool
	}

	testCases := []testCase{
		{
			name:     "should be true",
			expected: true,
		},
		{
			name:     "should be false",
			current:  1,
			new:      getRandomFloat64(),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v Float64

			// v.Load() is zero value for float64
			// if tc.current is not zero, should be false
			result := v.CompareAndSwap(tc.current, tc.new)

			if result != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, result)
			}
		})
	}
}

func getRandomFloat64() float64 {
	return MIN_F64 + rand.Float64()*(MAX_F64-MIN_F64)
}

func float64Equals(a, b float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < TOLERANCE_F64
	}

	return (d / math.Abs(b)) < TOLERANCE_F64
}
