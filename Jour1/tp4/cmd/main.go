package main

import (
	"fmt"
	"tp4/internal/vet"
)

func main() {
	fmt.Println("=== Rapport de simulation Clinique Vétérinaire ===")

	// 1) Créons un animal
	monChat := vet.Animal{
		Name:         "Minou",
		Age:          3,
		Race:         1234, // Ex. code de race
		Yiel:         "Miaou",
		HealthRecord: []string{"Vermifuge"},
	}

	// 2) Créons un client (celui qui possède l’animal)
	client := vet.Client{
		Person: vet.Person{
			Firstname: "Marie",
			Name:      "Dupont",
			Job:       "Enseignante",
		},
		Pet:     monChat.Name,
		PetInfo: "Très joueur et parfois capricieux",
	}

	// 3) Créons un docteur
	doctor := vet.Doctor{
		Person: vet.Person{
			Firstname: "Gregory",
			Name:      "House",
			Job:       "Docteur Vétérinaire",
		},
	}

	// 4) Créons une infirmière
	nurse := vet.Nurse{
		Person: vet.Person{
			Firstname: "Sophie",
			Name:      "Martin",
			Job:       "Infirmière",
		},
	}

	// --- Déroulement d'un mini-scénario ---

	// a) Le client parle de son animal
	client.ParleDeSonAnimal()

	// b) Le client répond au docteur (question de politesse)
	client.RepondAuDocteur()

	// c) Le docteur pose des questions
	doctor.PoserQuestions(monChat)

	// d) Le docteur crée une ordonnance
	ord := doctor.CreerOrdonnance(monChat)

	// e) Le docteur assigne l’ordonnance à l’infirmière
	doctor.AssignerOrdonnance(nurse, ord)

	// f) L’infirmière répond au docteur
	nurse.RepondreAuDocteur()

	// g) L’infirmière effectue les soins
	nurse.EffectuerSoins(monChat)

	// h) Le docteur demande au client de payer
	doctor.DemanderPaiement(client, 75.50)

	fmt.Println("=== Fin du scénario ===")
}
