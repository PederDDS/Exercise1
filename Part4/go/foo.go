// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    //TODO: increment i 1000000 times
	for x := 0; x<=1000000; x++ {
	i++
	}
}

func decrementing() {
    //TODO: decrement i 1000000 times
	for k:= 0; k<=1000000; k++ {
	i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())    
	runtime.GOMAXPROCS(runtime.NumCPU())                                    
    // TODO: Spawn both functions as goroutines
	
   	go incrementing()
    	go decrementing()
    	time.Sleep(100*time.Millisecond)
  	Println("The magic number is:", i)
}
