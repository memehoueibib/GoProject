package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

/*
Exercice 3 :
1) Un struct Request (Id int)
2) Un struct Service (Id int + Lst []*Request)
3) Méthodes :
   - RegisterRequest(*Request) error
   - ProcessRequest(id int) error (vérifie si l'ID existe dans Lst)
4) Exo3 : propose un menu pour enregistrer ou traiter une requête
*/

// Request représente une requête, identifiée par un ID.
type Request struct {
    Id int
}

// Service gère un ensemble de requêtes (Lst).
type Service struct {
    Id  int
    Lst []*Request
}

// RegisterRequest ajoute la requête au slice Lst.
func (s *Service) RegisterRequest(request *Request) error {
    s.Lst = append(s.Lst, request)
    return nil
}

// ProcessRequest cherche un ID parmi les requêtes stockées.
// Si on le trouve, on considère qu’on “traite” la requête, sinon on renvoie une erreur.
func (s *Service) ProcessRequest(id int) error {
    for _, req := range s.Lst {
        if req.Id == id {
            fmt.Printf("Traitement de la requête ID %d\n", id)
            return nil
        }
    }
    return fmt.Errorf("Aucune requête avec l'ID %d n'a été trouvée", id)
}

// Exo3 propose un menu interactif pour enregistrer et traiter des requêtes.
func Exo3() {
    // Initialisation du service (avec une slice vide pour Lst).
    service := &Service{
        Id:  1,
        Lst: make([]*Request, 0),
    }

    // Scanner pour lire les entrées utilisateur.
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n=== MENU ===")
        fmt.Println("1) Enregistrer une requête")
        fmt.Println("2) Traiter une requête existante")
        fmt.Println("3) Quitter")
        fmt.Print("Votre choix : ")

        if !scanner.Scan() {
            // Si on ne peut pas lire, on arrête.
            fmt.Println("Lecture impossible (EOF ou erreur). Arrêt.")
            return
        }
        choiceStr := scanner.Text()

        choice, err := strconv.Atoi(choiceStr)
        if err != nil {
            fmt.Println("Veuillez entrer un nombre valide.")
            continue
        }

        switch choice {
        case 1:
            // Enregistrer une requête
            fmt.Print("Entrez l'ID de la requête à enregistrer : ")
            if !scanner.Scan() {
                fmt.Println("Lecture impossible, on revient au menu.")
                continue
            }
            idStr := scanner.Text()
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("ID invalide :", err)
                continue
            }

            req := &Request{Id: id}
            if err := service.RegisterRequest(req); err != nil {
                fmt.Println("Erreur lors de l'enregistrement :", err)
            } else {
                fmt.Printf("Requête ID %d enregistrée.\n", id)
            }

        case 2:
            // Traiter une requête
            fmt.Print("Entrez l'ID de la requête à traiter : ")
            if !scanner.Scan() {
                fmt.Println("Lecture impossible, on revient au menu.")
                continue
            }
            idStr := scanner.Text()
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("ID invalide :", err)
                continue
            }

            if err := service.ProcessRequest(id); err != nil {
                fmt.Println("Erreur lors du traitement :", err)
            } else {
                fmt.Printf("Requête ID %d traitée avec succès.\n", id)
            }

        case 3:
            // Quitter
            fmt.Println("Au revoir.")
            return

        default:
            fmt.Println("Choix invalide, réessayez.")
        }
    }
}


