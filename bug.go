```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        
        var x int
        
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        x++
                }()
        }
        wg.Wait()
        fmt.Println(x)
}
```