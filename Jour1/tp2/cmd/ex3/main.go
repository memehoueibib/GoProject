package main

import (
	"tp2/internal/structs"
    "strings"
	"fmt"
)

func main() {
    // Création du Teacher et de sa classe
    teacher := &structs.Teacher{
        Name:          "M. Daribo",
        TotalStudents: 20,
        Students:      []string{"Perez", "Wendy", "Mathis", "Meme", "Moise", "Ihab", " Reste des élèves"},
    }

    // Liste d'élèves qu'on déclare "présents" (certains en minuscules)
    presentStudents := []string{"Perez", "wendy", "Mathis", "Meme","Moise", "Ihab", "Maelle"}

    // 1) Vérifier si chaque élève "présent" se trouve dans la classe
    for _, p := range presentStudents {
        if !teacher.IsStudentInClass(p) {
            fmt.Printf("ERREUR : l'élève \"%s\" n'est pas dans la classe de %s\n", p, teacher.Name)
        }
    }

    // 2) Marquer absents ceux qui ne sont pas dans la liste "presentStudents"
    for _, student := range teacher.Students {
        // On compare en ignorant la casse
        if !inSliceIgnoreCase(student, presentStudents) {
            err := teacher.MarkAbsent(student)
            if err != nil {
                fmt.Printf("Impossible de marquer %s absent : %v\n", student, err)
            }
        }
    }

	// 3) Afficher la liste finale
	fmt.Printf("\n--- Liste finale pour %s ---\n", teacher.Name)
	for i, s := range teacher.Students {
		fmt.Printf(" - Élève %d : %s\n", i+1, s)
	}

	// Pour afficher la taille du slice "presentStudents"
	fmt.Printf("\nNombre total d'élèves déclarés présents : %d\n", len(presentStudents))

	}

// inSliceIgnoreCase renvoie true si "val" est dans "list", en ignorant la casse
func inSliceIgnoreCase(val string, list []string) bool {
    valLower := strings.ToLower(val)
    for _, item := range list {
        if strings.ToLower(item) == valLower {
            return true
        }
    }
    return false
}
