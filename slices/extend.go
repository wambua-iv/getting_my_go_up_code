package main


// func Eppend(slice []int, elements ...int) []int {
//     n := len(slice)
//     total := len(slice) + len(elements)
//     if total > cap(slice) {
//         // Reallocate. Grow to 1.5 times the new size, so we can still grow.
//         newSize := total*3/2 + 1
//         newSlice := make([]int, total, newSize)
//         copy(newSlice, slice)
//         slice = newSlice
//     }
//     fmt.Println(len(slice))
//     slice = slice[:total]
//     copy(slice[n:], elements)
//     return slice
// }