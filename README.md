# clicker
![Static Badge](https://img.shields.io/badge/license-MIT-blue)
Minimalist Go library for simulating mouse clicks, movements, and basic automation based on Robotgo
## About

A lightweight Go library for mouse automation and click simulation. Control cursor movement and perform clicks.
## Installation

```bash
go get github.com/Jofich/clicker@latest
```
## Quick Start

Here's a simple example of how to use clicker
```go
package main

import (
	"context"
	"sync"
	"time"

	"github.com/jofich/clicker"
)

func main() {
	posX, posY := clicker.GetGlobalMousePos()
  //create context with timeout to cancel clicking
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
    //create new clicker
		mouse := clicker.NewClicker()
		mouse.SetPosition(posX+i*50, posY)
		wg.Add(1)
    //start clicking
		go func() {
			mouse.StartClicking(ctx, time.Millisecond*500*time.Duration(i+1), clicker.Args{
				MouseButton: clicker.Left,
			})
			wg.Done()
		}()
	}
  // waits for the context timeout expires
	wg.Wait()
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support
If you encounter any problems or have questions, please [open an issue](https://github.com/jofich/clicker/issues/new)
