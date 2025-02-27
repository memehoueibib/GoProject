package main

import "fmt"

func main() {
    var nom string
    var age int
    fmt.Print("Entre ton nom : ")
    fmt.Scanln(&nom)
    fmt.Print("Vous quel age ? : ")
    fmt.Scanln(&age)
    fmt.Println("Salut,", nom, "!", "")
}