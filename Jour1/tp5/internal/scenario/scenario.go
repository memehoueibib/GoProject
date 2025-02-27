// scenario.go
package scenario

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RunScenario gère la logique de saisie et lance des "scénarios fous" selon le patient choisi
func RunScenario(doctorName string, patients []string) {
	fmt.Printf("Bienvenue dans le cabinet médical de %s !\n", doctorName)
	fmt.Println("Patients possibles :", strings.Join(patients, ", "))
	fmt.Println("Entrez le nom d'un patient pour lancer un scénario (ou 'exit' pour quitter).")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture, réessayez.")
			continue
		}
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Fin de la consultation, à bientôt !")
			break
		}

		// Vérifier si le patient fait partie de la liste
		if !estDansListe(input, patients) {
			fmt.Println("Ce patient n'est pas dans la liste. Essayez un autre nom, ou 'exit'.")
			continue
		}

		// Lancer le scénario fou
		scenarioFou(doctorName, input)
	}
}

// estDansListe vérifie si le nom saisi est dans la liste des patients
func estDansListe(name string, list []string) bool {
	name = strings.ToLower(name)
	for _, p := range list {
		if strings.ToLower(p) == name {
			return true
		}
	}
	return false
}

// scenarioFou affiche un scénario amusant pour un patient donné
func scenarioFou(doctor, patient string) {
	fmt.Printf("\n*** Scénario délirant entre %s et %s ***\n", doctor, patient)

	switch strings.ToLower(patient) {
	case "perez":
		fmt.Printf("[%s] arrive en sautant à cloche‐pied, persuadé d'être un kangourou.\n", patient)
		fmt.Printf("[%s] %s l'examine et confirme qu'il a avalé trop de chewing-gums\n", doctor, doctor)
	case "wendy":
		fmt.Printf("[%s] se présente avec des lunettes 3D, dit qu'elle voit la salle en mode cartoon.\n", patient)
		fmt.Printf("[%s] %s lui prescrit un jus de fruit magique pour revenir à la réalité.\n", doctor, doctor)
	case "mathis":
		fmt.Printf("[%s] prétend entendre des voix lui chantant du reggae.\n", patient)
		fmt.Printf("[%s] %s lui conseille de devenir DJ pour exprimer ce talent.\n", doctor, doctor)
	case "meme":
		fmt.Printf("[%s] crée des memes internet sur place, photographiant le docteur et en faisant un montage.\n", patient)
		fmt.Printf("[%s] %s réclame 50%% des droits d'auteur sur chaque meme viral.\n", doctor, doctor)
	case "moise":
		fmt.Printf("[%s] tente de diviser la salle d'attente en deux, comme la mer.\n", patient)
		fmt.Printf("[%s] %s le prévient que c'est interdit par la sécurité incendie.\n", doctor, doctor)
	case "ihab":
		fmt.Printf("[%s] est arrivé en volant sur un tapis de yoga, persuadé qu'il a découvert la lévitation.\n", patient)
		fmt.Printf("[%s] %s observe, perplexe, et prend des notes pour un futur best-seller.\n", doctor, doctor)
	}

	fmt.Printf("\n[%s] Après ce spectacle incroyable, %s se gratte la tête :\n", doctor, doctor)
	fmt.Println("    « Hmmm... Je n'avais encore jamais vu ça dans toute ma carrière ! »")
	fmt.Printf("[%s] Diagnostique : 'Vous souffrez d'un excès d'imagination combiné à un déficit de logique.'\n", doctor)
	fmt.Printf("[%s] Conseille : 'Dormez 8h, faites 15 minutes de danse, et évitez de manger des escargots au chocolat.'\n", doctor)
	fmt.Printf("*** Fin de consultation pour %s ***\n\n", patient)
}
