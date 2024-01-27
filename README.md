## afloat

<img align="right" width="200px" height="150px" src="https://i.ibb.co/pR5T8vb/Whats-App-Image-2024-01-26-at-18-26-40.jpg">

The `afloat` package provides a wrapper for atomic operations for the `float32` and `float64` types, which are not natively supported by the `sync/atomic` package.

![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
![Go](https://img.shields.io/badge/Go-v1.21-blue)


### How to use - WIP

```go
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
```
