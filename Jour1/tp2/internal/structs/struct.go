package structs

import (    
	"errors"
	"fmt"
	"strings"
)

// Rectangle stocke une largeur et une hauteur
type Rectangle struct {
	Width  int
	Height int
}

// Area affiche et retourne (au choix) l'aire ou le périmètre
func (r Rectangle) Area() int {
	fmt.Println("Area is ", r.Height*r.Width)
	return (r.Height * 2) + (r.Width * 2)
}

/*
   Exo 3
   Créer un programme qui permet de faire l’appel des élèves :
   1. Lire la liste des élèves et la liste des élèves présents
   2. Si l’élève est absent, modifier la case associée
   3. Faire la gestion d’erreur si l’élève n’est pas dans la liste
*/

// Teacher stocke le nom, le nombre total d'élèves et leur liste
type Teacher struct {
	Name          string
	TotalStudents int
	Students      []string
}


// MarkAbsent cherche l'élève en ignorant la casse.
// S'il est trouvé, on préfixe son nom par "ABSENT-".
func (t *Teacher) MarkAbsent(student string) error {
    studentLower := strings.ToLower(student)

    for i, s := range t.Students {
        if strings.ToLower(s) == studentLower {
            t.Students[i] = "ABSENT(S) -" + s
            return nil
        }
    }
    return errors.New("élève introuvable : " + student)
}

// IsStudentInClass vérifie la présence en ignorant la casse.
func (t *Teacher) IsStudentInClass(name string) bool {
    nameLower := strings.ToLower(name)

    for _, s := range t.Students {
        if strings.ToLower(s) == nameLower {
            return true
        }
    }
    return false
}