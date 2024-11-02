package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type student struct {
	grade map[string]int
}
type value map[string]int

func main() {
	preMutexed()
	enabledMutex()
}

func preMutexed() {

	var wg sync.WaitGroup
	var val atomic.Value

	val.Store(student{grade: map[string]int{}})

	wg.Add(4)

	go func() {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["Geometry"] = 89
		val.Store(s)
	}()

	go func() {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["Computer"] = 97
		val.Store(s)
	}()

	go func() {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["VideoGraphy"] = 89
		val.Store(s)
	}()

	go func() {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["Mechanics"] = 89
		val.Store(s)
	}()

	wg.Wait()

	fmt.Println(val)

}

func addToStruct(wg *sync.WaitGroup, m value, st student) {

	var mutex sync.Mutex
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	for course, score := range m {

		st.grade[course] = score
	}

}

func enabledMutex() {
	//keyvalueslice := make([]map[string]interface{}, 1, 1)

	val := map[string]int{"concurrency": 76,
		"errors": 76, "stdlib": 76, "context": 76,
		"Computer": 97, "Geometry": 89, "Mechanics": 89, "VideoGraphy": 89,
	}

	var wg sync.WaitGroup
	st := student{grade: map[string]int{}}
	st2 := student{grade: map[string]int{}}

	wg.Add(1)
	now := time.Now()
	go addToStruct(&wg, val, st)
	fmt.Println("go", time.Since(now))
	wg.Wait()

	now2 := time.Now()
	for course, score := range val {
		st2.grade[course] = score
	}
	fmt.Println("main", time.Since(now2))

}
