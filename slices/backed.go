// an attempt to change the bacling array of a slice
package main

import (
	"fmt"
	"os"
	"regexp"
)

var digitRegexep = regexp.MustCompile("[0-9]+")

func findDigits (filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexep.Find(b)
}

func main () {
	b := findDigits("rando.txt")
	c := []byte{}
	numbers := []int{2,45,6,7,8,3,5}
	fmt.Println(c)		
	c = append(c,b...)
	fmt.Println(&b[0])
	fmt.Println(&c[0])
	fmt.Println(c)
	numbers = Eppend(numbers, 56,43,3,5,5,67)
	fmt.Println(numbers)
	
}

func Eppend(slice []int, elements ...int) []int {
    n := len(slice)
    total := len(slice) + len(elements)
    if total > cap(slice) {
        // Reallocate. Grow to 1.5 times the new size, so we can still grow.
        newSize := total*3/2 + 1
        newSlice := make([]int, total, newSize)
        copy(newSlice, slice)
        slice = newSlice
    }
     fmt.Println(len(slice))
    // slice = slice[:total]
    copy(slice[n:], elements)
    return slice
}


