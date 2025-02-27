// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Fonction qui affiche un message
// func direBonjour(nom string, wg *sync.WaitGroup) {
// 	defer wg.Done() // Signale que la goroutine est terminée
// 	fmt.Printf("Bonjour, %s!\n", nom)
// }

// func main() {
// 	// Liste des noms à saluer
// 	noms := []string{"Alice", "Bob", "Charlie"}

// 	var wg sync.WaitGroup // Création du WaitGroup

// 	// Lancer une goroutine pour chaque nom
// 	for _, nom := range noms {
// 		wg.Add(1) // Ajoute une tâche au WaitGroup
// 		go direBonjour(nom, &wg)
// 	}

// 	wg.Wait() // Attend la fin de toutes les goroutines
// 	fmt.Println("Tous les messages ont été affichés.")
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Fonction qui calcule le carré d'un nombre et l'envoie dans le channel
// func calculerCarre(n int, ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Indique que cette goroutine est terminée
// 	fmt.Println(time.Now(),n*n)
// 	ch <- n * n     // Envoi du résultat au channel
// }

// func main() {
// 	nombres := []int{2, 3, 4} // Liste des nombres à calculer
// 	ch := make(chan int, len(nombres)) // Channel bufferisé
// 	var wg sync.WaitGroup // Création du WaitGroup

// 	// Lancer une goroutine pour chaque nombre
// 	for _, n := range nombres {
// 		wg.Add(1) // Ajoute une tâche au WaitGroup
// 		go calculerCarre(n, ch, &wg)
// 	}

// 	wg.Wait() // Attend que toutes les goroutines terminent
// 	close(ch) // Fermeture du channel après la fin des calculs

// 	// Récupérer et afficher les résultats
// 	for resultat := range ch {
// 		fmt.Printf("Carré de %d : %d\n", int(resultat/2), resultat)
// 	}
// }



// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// // Fonction worker (chaque worker récupère et traite une tâche)
// func worker(id int, tasks chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range tasks {
// 		fmt.Printf("Worker %d traite la tâche %d\n", id, task)
// 		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Simule le traitement
// 		fmt.Printf("Worker %d a terminé la tâche %d\n", id, task)
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	numWorkers := 3
// 	numTasks := 10

// 	tasks := make(chan int, numTasks) // Channel pour les tâches
// 	var wg sync.WaitGroup             // WaitGroup pour synchroniser

// 	// Lancer les workers
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, tasks, &wg)
// 	}

// 	// Ajouter des tâches dans le channel
// 	for i := 1; i <= numTasks; i++ {
// 		tasks <- i
// 	}

// 	close(tasks) // Fermeture du channel après l'ajout des tâches
// 	wg.Wait()    // Attente de la fin de tous les workers

// 	fmt.Println("Toutes les tâches ont été traitées.")
// }


package main

import (
	"fmt"
	"sync"
	"time"
)

// Structure représentant une commande
type Commande struct {
	ID   int
	Plat string
}

// Liste des plats disponibles
var platsDisponibles = []string{"Pizza", "Burger", "Sushi", "Steak", "Pâtes", "Steak Rare"}

func serveur(id int, commandes chan Commande, wg *sync.WaitGroup) {
	defer wg.Done() // Signale que le serveur a fini
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🛑 Serveur %d a rencontré une erreur : %v\n", id, r)
		}
	}()

	for commande := range commandes {
		if commande.Plat == "Steak Rare" {
			panic(fmt.Sprintf("Ingrédient manquant pour la commande %d", commande.ID))
		}

		fmt.Printf("✅ Serveur %d traite la commande %d : %s\n", id, commande.ID, commande.Plat)
		time.Sleep(time.Second) // Simulation du temps de traitement
		fmt.Printf("🍽 Serveur %d a terminé la commande %d\n", id, commande.ID)
	}
}

func main() {
	numServeurs := 3
	numCommandes := 10

	commandes := make(chan Commande, numCommandes) // Création du channel
	var wg sync.WaitGroup

	// Lancer les serveurs (workers)
	for i := 1; i <= numServeurs; i++ {
		wg.Add(1)
		go serveur(i, commandes, &wg)
	}

	// Génération et envoi des commandes
	for i := 1; i <= numCommandes; i++ {
		plat := platsDisponibles[i%len(platsDisponibles)] // Sélection d'un plat
		commandes <- Commande{ID: i, Plat: plat}          // Envoi dans le channel
	}

	close(commandes) // Fermeture du channel après l’ajout des commandes
	wg.Wait()        // Attendre que tous les serveurs terminent

	fmt.Println("🏁 Toutes les commandes ont été traitées.")
}
