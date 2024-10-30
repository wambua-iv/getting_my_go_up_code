package main

import (
	"fmt"
	"runtime"
	"time"
)

//this function waits for no one
func startConcurrent () {
	go func()  {
		//need concurrent code here
		concuT1()
		concuT2()
		concuT3()
	}()
	fmt.Println("Main function exiting")
}

func concuT1() {
	time.Sleep(500*time.Millisecond)
	fmt.Println("from function 1")
}
func concuT2 () {
	time.Sleep(100* time.Millisecond)
	fmt.Println("from function two")
}
func concuT3 () {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(runtime.NumCPU())
}

//attempted to force a wait from main

func attemptedToWait() {
	now := time.Now()
	time.Sleep(2 * time.Second)
	startConcurrent()
	fmt.Printf("Time elapsed is %v", time.Since(now))
}