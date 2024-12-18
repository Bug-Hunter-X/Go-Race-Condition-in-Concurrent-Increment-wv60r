# Go Race Condition Example

This repository demonstrates a common race condition in Go when multiple goroutines concurrently modify a shared variable without proper synchronization. The `bug.go` file contains the buggy code, which leads to an unpredictable result.  The `bugSolution.go` file provides a corrected version using a mutex to prevent the race condition.