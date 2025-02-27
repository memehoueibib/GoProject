package main

import (
    "bufio"   // Permet de lire du texte ligne par ligne (depuis la console, par exemple)
    "fmt"     // Permet d'afficher des messages à l'écran et d’utiliser Scanf
    "os"      // Fournit des fonctions pour manipuler le système (entrée, sortie, etc.)
    "strconv" // Fournit des fonctions pour convertir des chaînes de caractères en nombres
)

/*
Exercice 2 :
Écrivez un programme qui lit un nombre entier à partir de l'entrée standard.
Si le nombre est négatif, provoquez une exception (panic) avec le message
"Le nombre doit être positif". Utilisez la fonction recover pour gérer
l’exception et imprimer le message d’erreur.
*/

// Exo2 lit un entier depuis l'entrée standard, déclenche un panic s'il est négatif et gère l'erreur avec recover().
func Exo2() {
    // Le defer ci-dessous permet de capturer un éventuel panic grâce à recover().
    defer func() {
        // Vérifie si un panic s'est produit.
        if r := recover(); r != nil {
            // Affiche le message du panic capturé.
            fmt.Println("Exo2 -> Erreur capturée par recover():", r)
        }
    }()

    // Invite l'utilisateur à saisir un nombre entier positif.
    fmt.Print("Exo2 -> Entrez un nombre entier positif: ")

    /*
       ─────────────────────────────────────────────────────────────────────────────────
       APPROCHE 1 : Utiliser un bufio.Scanner
       ─────────────────────────────────────────────────────────────────────────────────
    */
    scanner := bufio.NewScanner(os.Stdin)
    // Vérifie si l’on parvient à lire une ligne.
    if scanner.Scan() {
        // Récupère la chaîne saisie par l’utilisateur.
        input := scanner.Text()

        // Convertit la chaîne en entier.
        number, err := strconv.Atoi(input)
        // Déclenche un panic si la conversion échoue (par exemple, si le texte n’est pas un nombre valide).
        if err != nil {
            panic("Valeur non entière ou invalide.")
        }

        // Déclenche un panic si le nombre est négatif.
        if number < 0 {
            panic("Le nombre doit être positif")
        }

        // Si tout va bien, on affiche le nombre lu.
        fmt.Println("Exo2 -> Le nombre saisi est:", number)

    } else {
        // Si la lecture échoue (EOF ou autre), on déclenche un panic.
        panic("Impossible de lire l'entrée.")
    }

    /*

       // var number int
       // _, err := fmt.Scanf("%d", &number)
       // if err != nil {
       //     panic("Valeur non entière ou invalide.")
       // }
       
       // if number < 0 {
       //     panic("Le nombre doit être positif")
       // }
       
       // fmt.Println("Exo2 -> Le nombre saisi est:", number)
    */
}
