package main

import (
	"fmt"
	"sync/atomic"
)

type student struct {
	grade map[string]int
}

var val atomic.Value

//function has race conditions within it 
func seekingAtomicity () {

	val.Store(student{grade: map[string]int {}})

	wg.Add(4)

	go func () {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["maths"] = 98
	}()
	go func () {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["English"] = 78 
	}()
	go func () {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["Physics"] = 88 
	}()
	go func () {
		defer wg.Done()
		s := val.Load().(student)
		s.grade["Chemistry"] = 98 
	}()

	wg.Wait()
	fmt.Println(val.Load().(student))
}

//using the wait group store function to get rid of race conditions
func foundAtomicity () {

	val.Store(student{grade: map[string]int {}})

	wg.Add(4)

	go func () {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grade
		m["Geometry"] = 89
		val.Store(student{grade: m})
	}()

	go func () {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grade
		m["Computer"] = 97
		val.Store(student{grade: m})
	}()

	go func () {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grade
		m["VideoGraphy"] = 89
		val.Store(student{grade: m})
	}()	
	
	go func () {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grade
		m["Mechanics"] = 89
		val.Store(student{grade: m})
	}()

	wg.Wait()
	
	fmt.Println(val.Load().(student))
}