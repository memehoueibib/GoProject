package main

import (
    "fmt" // Permet d'afficher des messages à l'écran
    "os"  // Fournit des fonctions pour manipuler des fichiers et répertoires
)

/*
Exercice 1 :
Écrivez un programme qui ouvre un fichier et y écrit le message "Hello, world!".
Utilisez la fonction defer pour fermer le fichier, même si une erreur se produit.
*/

// Exo1 crée (ou remplace) un fichier "hello.txt", y écrit un message et assure sa fermeture.
func Exo1() {
    // Ouvre ou crée un fichier nommé "hello.txt".
    file, err := os.Create("hello.txt")
    // Vérifie si une erreur s’est produite lors de la création du fichier.
    if err != nil {
        // Affiche l'erreur et quitte la fonction si le fichier n'a pas pu être ouvert/créé.
        fmt.Println("Erreur lors de la création du fichier:", err)
        return
    }

    // Utilise defer pour s'assurer que le fichier soit fermé avant de quitter la fonction,
    // même si un problème survient plus bas.
    defer func() {
        // Tente de fermer le fichier.
        cerr := file.Close()
        // Affiche une éventuelle erreur de fermeture du fichier.
        if cerr != nil {
            fmt.Println("Erreur lors de la fermeture du fichier:", cerr)
        }
    }()

    // Écrit le texte "Hello, world!\n" dans le fichier.
    _, err = file.WriteString("Hello, world!\n")
    // Vérifie si une erreur d’écriture est survenue.
    if err != nil {
        // Affiche l’erreur et quitte la fonction si l'écriture a échoué.
        fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
        return
    }

    // Indique que l’opération s’est bien passée.
    fmt.Println("Exo1 -> \"Hello, world!\" a été écrit dans hello.txt.")
}
