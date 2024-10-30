package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//this function waits for no one
func startConcurrent () {
	wg.Add(3)
	go func()  {
		//need concurrent code here
		concuT1(&wg)
		concuT2(&wg)
		concuT3(&wg)
	}()
	wg.Wait()
	fmt.Println("Main function exiting")
}

func concuT1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500*time.Millisecond)
	fmt.Println("from function 1")
}
func concuT2 (wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100* time.Millisecond)
	fmt.Println("from function two")
}
func concuT3 (wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(runtime.NumCPU())
}

//attempted to force a wait from main

func attemptedToWait() {
	now := time.Now()
 	startConcurrent()
	fmt.Printf("Time elapsed is %v", time.Since(now))
}