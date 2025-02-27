package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bienvenue sur la page d'accueil !")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "message": "Salut depuis l'API JSON !",
        "status":  "succès",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", homeHandler)         // Route "/"
    http.HandleFunc("/api", apiHandler)       // Route "/api"
    fmt.Println("Serveur démarré sur http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}