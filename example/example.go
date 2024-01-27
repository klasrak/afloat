package main

import (
	"fmt"
	"sync"

	"github.com/klasrak/afloat"
)

func main() {
	var (
		v  afloat.Float32
		wg sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			v.Add(1.0) // Concurrently add a float32 to v
		}()
	}

	wg.Wait()

	fmt.Println(v.Load()) // 100
}
