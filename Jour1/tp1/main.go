package main

import (
    "fmt"
    "tp1/maths" 
)

func main() {
    sum := maths.Add(3, 4)
    fmt.Println("Sum(3,4) =", sum)

    fmt.Println("Suite Fibonacci :")
    maths.Fibo()

    numbers := []int{5, 3, 6, 2, 1}
    fmt.Println("Is ordered before sorting ?", maths.Ordered(numbers))

    maths.AscendingSort(numbers)
    fmt.Println("Sorted numbers:", numbers)
    fmt.Println("Is ordered after sorting ?", maths.Ordered(numbers))
}
