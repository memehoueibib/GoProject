package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Journal contient : 
//   - currentDay : le jour sur lequel on ajoute les commentaires
//   - content    : un map associant un jour (int) à une liste de commentaires (slice de string)
type Journal struct {
    CurrentDay int
    Content    map[int][]string
}

// addDay incrémente le jour courant et prépare une liste vide pour ce nouveau jour.
func (j *Journal) addDay() {
    j.CurrentDay++
    if j.Content[j.CurrentDay] == nil {
        j.Content[j.CurrentDay] = []string{}
    }
    fmt.Printf("Nouveau jour créé : %d\n", j.CurrentDay)
}

// addComment ajoute un commentaire dans la liste du jour courant.
func (j *Journal) addComment(comment string) {
    j.Content[j.CurrentDay] = append(j.Content[j.CurrentDay], comment)
    fmt.Printf("Commentaire ajouté au jour %d\n", j.CurrentDay)
}

// getDay affiche tous les commentaires du jour spécifié.
func (j *Journal) getDay(day int) {
    comments, existe := j.Content[day]
    if !existe {
        fmt.Printf("Le jour %d n'existe pas.\n", day)
        return
    }
    if len(comments) == 0 {
        fmt.Printf("Aucun commentaire pour le jour %d.\n", day)
        return
    }
    fmt.Printf("Commentaires du jour %d :\n", day)
    for i, c := range comments {
        fmt.Printf("  %d) %s\n", i+1, c)
    }
}

// saveJournal demande un nom de fichier et enregistre le journal en JSON.
func (j *Journal) saveJournal(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("impossible de créer le fichier : %v", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(j); err != nil {
        return fmt.Errorf("erreur d'encodage JSON : %v", err)
    }

    fmt.Printf("Journal sauvegardé dans %s\n", filename)
    return nil
}

// loadJournal charge un fichier JSON dans la struct Journal.
func (j *Journal) loadJournal(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("impossible d'ouvrir le fichier : %v", err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(j); err != nil {
        return fmt.Errorf("erreur de décodage JSON : %v", err)
    }

    fmt.Printf("Journal chargé depuis %s\n", filename)
    return nil
}

func Exo4() {
    // Initialisation du journal : jour courant = 0, map vide.
    journal := &Journal{
        CurrentDay: 0,
        Content:    make(map[int][]string),
    }

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("=== Journal simplifié + save/load ===")
    fmt.Println("Commandes disponibles :")
    fmt.Println("  /add-day               -> créer un nouveau jour")
    fmt.Println("  /add-comment           -> ajouter un commentaire au jour courant")
    fmt.Println("  /get-day               -> afficher un jour (on vous demandera le numéro)")
    fmt.Println("  /save                  -> sauvegarder le journal (on vous demandera le nom du fichier)")
    fmt.Println("  /load-[fichier.json]   -> charger le journal depuis un fichier")
    fmt.Println("  /quit                  -> quitter le programme")

    for {
        fmt.Print("\nEntrez une commande : ")
        if !scanner.Scan() {
            fmt.Println("Fin de la lecture (EOF ?). Au revoir.")
            break
        }
        cmd := strings.TrimSpace(scanner.Text())

        switch {
        case cmd == "/add-day":
            journal.addDay()

        case cmd == "/add-comment":
            if journal.CurrentDay == 0 {
                fmt.Println("Aucun jour n'existe encore. Utilisez /add-day d'abord.")
                continue
            }
            fmt.Print("Entrez votre commentaire : ")
            if scanner.Scan() {
                comment := scanner.Text()
                journal.addComment(comment)
            } else {
                fmt.Println("Lecture du commentaire impossible.")
            }

        case cmd == "/get-day":
            fmt.Print("Quel jour voulez-vous afficher ? ")
            if scanner.Scan() {
                dayStr := scanner.Text()
                day, err := strconv.Atoi(dayStr)
                if err != nil {
                    fmt.Println("Veuillez entrer un nombre valide.")
                    continue
                }
                journal.getDay(day)
            }

        case cmd == "/save":
            fmt.Print("Nom du fichier pour la sauvegarde (ex: journal.json) : ")
            if scanner.Scan() {
                filename := scanner.Text()
                if err := journal.saveJournal(filename); err != nil {
                    fmt.Println("Erreur lors de l'enregistrement :", err)
                }
            } else {
                fmt.Println("Impossible de lire le nom du fichier.")
            }

        case strings.HasPrefix(cmd, "/load-"):
            // Extraire la partie après /load-
            filename := strings.TrimPrefix(cmd, "/load-")
            if filename == "" {
                fmt.Println("Syntaxe : /load-nomFichier.json")
                continue
            }
            if err := journal.loadJournal(filename); err != nil {
                fmt.Println("Erreur lors du chargement :", err)
            }

        case cmd == "/quit":
            fmt.Println("Au revoir.")
            return

        default:
            fmt.Println("Commande inconnue. Essayez /add-day, /add-comment, /get-day, /save, /load-..., /quit.")
        }
    }
}
