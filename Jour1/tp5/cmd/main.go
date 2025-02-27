// main.go
package main

import (
	"tp5/internal/scenario"// Chemin d'import selon votre module ou dossier
)

func main() {
	// Nom du docteur
	doctorName := "M. Daribo"

	// Liste de patients
	patients := []string{"Perez", "Wendy", "Mathis", "Meme", "Moise", "Ihab"}

	// On lance le scénario
	scenario.RunScenario(doctorName, patients)
}
