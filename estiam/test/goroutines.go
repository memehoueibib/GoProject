package main

import (
    "fmt"
    "time"
)

func direBonjour() {
    fmt.Println("Bonjour depuis une goroutine !")
}

func main4() {
    go direBonjour() // Lance en parallèle
    fmt.Println("Salut depuis main !")
    time.Sleep(time.Second) // Attends un peu pour voir les deux messages
}