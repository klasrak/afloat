## afloat

<img align="right" width="200px" height="150px" src="https://i.ibb.co/pR5T8vb/Whats-App-Image-2024-01-26-at-18-26-40.jpg">

Generate a coverage badge like this one for your Golang projects without uploading results to a third party.

![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
![Go](https://img.shields.io/badge/Go-v1.21-blue)


### How to use

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
