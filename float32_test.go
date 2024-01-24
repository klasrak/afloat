package afloat

import (
	"math"
	"math/rand"
	"sync"
	"testing"
)

const (
	MIN_F32       = math.SmallestNonzeroFloat32
	MAX_F32       = math.MaxFloat32 / 10
	TOLERANCE_F32 = 1e-12
)

func TestAdd32(t *testing.T) {

	type testCase struct {
		name         string
		value        float32
		delta        float32
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
			value:        getRandomFloat32(),
			delta:        getRandomFloat32(),
			maxAdditions: 1,
		},
		{
			name:         "Add random value to random value 100 times",
			value:        getRandomFloat32(),
			delta:        getRandomFloat32(),
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
			name:         "Add negative random value to positive random value",
			value:        getRandomFloat32(),
			delta:        getRandomFloat32() * -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add negative random value to positive random value 100 times",
			value:        getRandomFloat32(),
			delta:        getRandomFloat32() * -1.0,
			maxAdditions: 100,
		},
		{
			name:         "Add positive random value to negative random value",
			value:        getRandomFloat32() * -1.0,
			delta:        getRandomFloat32(),
			maxAdditions: 1,
		},
		{
			name:         "Add positive random value to negative random value 100 times",
			value:        getRandomFloat32() * -1.0,
			delta:        getRandomFloat32(),
			maxAdditions: 100,
		},
		{
			name:         "Add negative random value to negative random value",
			value:        getRandomFloat32() * -1.0,
			delta:        getRandomFloat32() * -1.0,
			maxAdditions: 1,
		},
		{
			name:         "Add negative random value to negative random value 100 times",
			value:        getRandomFloat32() * -1.0,
			delta:        getRandomFloat32() * -1.0,
			maxAdditions: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wg := sync.WaitGroup{}

			expected := tc.value + (tc.delta * float32(tc.maxAdditions))

			for i := 0; i < tc.maxAdditions; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					Add32(&tc.value, tc.delta)
				}()
			}

			wg.Wait()

			if !float32Equals(tc.value, expected, TOLERANCE_F32) {
				t.Errorf("Expected %.18f, got %.18f", expected, tc.value)
			}
		})
	}
}

func getRandomFloat32() float32 {
	return rand.Float32()*(MAX_F32-MIN_F32) + MIN_F32
}

func float32Equals(a, b, tolerance float32) bool {
	if a == b {
		return true
	}

	d := math.Abs(float64(a) - float64(b))

	if b == 0 {
		return d < float64(tolerance)
	}

	return (d / math.Abs(float64(b))) < float64(tolerance)
}
