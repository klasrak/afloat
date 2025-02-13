package main

import (
	"fmt"
	"sync"

	"github.com/klasrak/afloat"
)

func main() {
	var (
		f32 afloat.Float32
		f64 afloat.Float64
		wg  sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			f32.Add(1.0) // Concurrently add a float32 to f32
			f64.Add(1.0) // Concurrently add a float64 to f64
		}()
	}

	wg.Wait()

	fmt.Println(f32.Load()) // 100
	fmt.Println(f64.Load()) // 100
}
