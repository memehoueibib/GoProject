package maths

import (
    "fmt"
    "sort"
)

// Add renvoie la somme de deux entiers.
func Add(n1, n2 int) int {
    return n1 + n2
}

// Fibo affiche la suite de Fibonacci (par exemple sur 10 termes).
func Fibo() {
    a, b := 0, 1
    for i := 0; i < 10; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }
}

// Ordered vérifie l'ordre croissant.
func Ordered(numbers []int) bool {
    for i := 0; i < len(numbers)-1; i++ {
        if numbers[i] > numbers[i+1] {
            return false
        }
    }
    return true
}

// AscendingSort trie un slice d'entiers en place.
func AscendingSort(numbers []int) {
    sort.Ints(numbers)
}
