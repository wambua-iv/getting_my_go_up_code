package main

import (
	"sync"
)

type value struct {
	mu    sync.Mutex
	value int
}

var data int

func main() {
	con()
}

func con() {
	// go func() {
	// 	data++
	// }()

	// if data == 0 {
	// 	fmt.Printf("the value is %v  ", data)
	// } else {
	// 	fmt.Printf("the value is %v.\n", data)
	// }

	// var wg sync.WaitGroup
	// printSum := func(v1, v2 *value) {
	// 	defer wg.Done()
	// 	v1.mu.Lock()
	// 	defer v1.mu.Unlock()
	// 	time.Sleep(2 * time.Second)
	// 	v2.mu.Lock()
	// 	defer v2.mu.Unlock()
	// 	fmt.Printf("sum=%v\n", v1.value+v2.value)
	// }
	// var a, b value
	// wg.Add(2)
	// go printSum(&a, &b)
	// go printSum(&b, &a)
	// wg.Wait()

	startConcurrent()
	attemptedToWait()
}
