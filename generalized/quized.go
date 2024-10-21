package main

import "fmt"

type Integer interface {
	~int | ~int8 | ~int16 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16
}

func Doubler[T Integer](v T) T {
	return v * 2
}

//------------------------------------------------------------------------------//

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (p PrintInt) String () string {
	return fmt.Sprintf("%d",p)
}

type PrintFloat float64

func (p PrintFloat) String() string {
	return fmt.Sprintf("%f", p)
}

func PrintIt[T Printable](t T){
	fmt.Println(t)
}