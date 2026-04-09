// Go by Example: For
// https://gobyexample.com/for
package main

import "fmt"

func main() {
	// Go has only one looping construct, the for loop. Here are the three basic types.
	// The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
	// The classic initial/condition/after for loop, with the init and post statements optional.
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }
	// The range form of the for loop iterates over a slice or map. This is the most common
    for i := range 3 {
        fmt.Println("range", i)
    }
	// You can also break and continue to control the loop flow.
    for {
        fmt.Println("loop")
        break
    }
	// This loop will print only odd numbers.
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}