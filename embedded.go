package main

import (
	"fmt"
	"time"
)

type Projects struct {
	established time.Time
	ongoing     bool
	stalled     bool
	promises    string
}

type ProjectBy struct {
	Projects
	president string
}

func (p Projects) PrintPresident(name string) string {
	return fmt.Sprintf("Mwizi by name %s", name)
}

func (p ProjectBy) PrintPresident(name string) string {
	return fmt.Sprintf("He that stole %s", name)
}

// adding to the method set of a containing struct
type Innered struct {
	A int
}

func (i Innered) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Innered) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Innered
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func callOuter() {
	o := Outer{
		Innered: Innered{
			A: 30,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
}
