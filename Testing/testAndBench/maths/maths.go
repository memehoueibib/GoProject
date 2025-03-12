package maths

import "fmt"

// addition de deux nombres
func Add(n1, n2 int) int {
	return n1 + n2
}

// Fonction qui affiche les 50 premiers
// termes de la suite de Fibonacci
func Fibo() {
	a := 0
	b := 1
	for i := 0; i < 50; i++ {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
	}
}

// Ordered vérifie si le slice de nombres est trié par ordre croissant.
func Ordered(numbers []int) bool {
	if len(numbers) <= 1 {
		return true
	}
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}

// AscendingSort trie des nombres en ordre croissant (méthode basique).
func AscendingSort(numbers []int) {
	for !Ordered(numbers) {
		for i := 1; i < len(numbers); i++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
			}
		}
	}
}
