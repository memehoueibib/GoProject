package vet

import "fmt"

// Person : info de base sur un individu
type Person struct {
	Firstname string
	Name      string
	Job       string
}

// Animal : informations sur l’animal
type Animal struct {
	Name         string
	Age          int
	Race         int
	Yiel         string              // Ex. "Miaou", "Wouf"
	HealthRecord []string            // Dossier médical (ex: médocs déjà pris)
}

// Doctor : un docteur vétérinaire
type Doctor struct {
	Person
	AssignedPatient string
}

// Nurse : une infirmière
type Nurse struct {
	Person
	AssignedPet string
	PillsList   string
}

// Client : un client (maître de l’animal)
type Client struct {
	Person
	Pet     string // Nom de l’animal
	PetInfo string // Info complémentaire
}

// Ordonnance : simple struct pour prescrire des médicaments
type Ordonnance struct {
	Medicaments []string
}

// ---------------------------------------------------------------------
// Interfaces

// IDoctor : rôles/méthodes d’un docteur
type IDoctor interface {
	PoserQuestions(a Animal)
	CreerOrdonnance(a Animal) Ordonnance
	DemanderPaiement(c IClient, montant float64)
	AssignerOrdonnance(inf INurse, ord Ordonnance)
}

// INurse : rôles/méthodes d’une infirmière
type INurse interface {
	RepondreAuDocteur()
	LireOrdonnance(ord Ordonnance)
	EffectuerSoins(a Animal)
}

// IClient : rôles/méthodes d’un client
type IClient interface {
	ParleDeSonAnimal()
	RepondAuDocteur()
}

// ---------------------------------------------------------------------
// Implémentations du Doctor (satisfait IDoctor)

func (d Doctor) PoserQuestions(a Animal) {
	fmt.Printf("[Docteur %s] Pose des questions sur l’animal %s.\n", d.Name, a.Name)
}

func (d Doctor) CreerOrdonnance(a Animal) Ordonnance {
	fmt.Printf("[Docteur %s] Création d’ordonnance pour l’animal %s.\n", d.Name, a.Name)
	// Ex. simple : on prescrit toujours les mêmes médicaments.
	return Ordonnance{Medicaments: []string{"Antibiotique", "Vitamines"}}
}

func (d Doctor) DemanderPaiement(c IClient, montant float64) {
	fmt.Printf("[Docteur %s] Demande au client de payer %.2f €.\n", d.Name, montant)
}

func (d Doctor) AssignerOrdonnance(inf INurse, ord Ordonnance) {
	fmt.Printf("[Docteur %s] Assigne l’ordonnance à l’infirmière.\n", d.Name)
	inf.LireOrdonnance(ord)
}

// ---------------------------------------------------------------------
// Implémentations du Nurse (satisfait INurse)

func (n Nurse) RepondreAuDocteur() {
	fmt.Printf("[Infirmière %s] Je suis disponible, docteur.\n", n.Name)
}

func (n Nurse) LireOrdonnance(ord Ordonnance) {
	n.PillsList = fmt.Sprintf("%v", ord.Medicaments)
	fmt.Printf("[Infirmière %s] Lit l’ordonnance : %v.\n", n.Name, ord.Medicaments)
}

func (n Nurse) EffectuerSoins(a Animal) {
	fmt.Printf("[Infirmière %s] Soigne l’animal %s.\n", n.Name, a.Name)
}

// ---------------------------------------------------------------------
// Implémentations du Client (satisfait IClient)

func (cl Client) ParleDeSonAnimal() {
	fmt.Printf("[Client %s] Bonjour, mon animal s’appelle %s.\n", cl.Name, cl.Pet)
	if cl.PetInfo != "" {
		fmt.Printf("[Client %s] Infos sur mon animal : %s\n", cl.Name, cl.PetInfo)
	}
}

func (cl Client) RepondAuDocteur() {
	fmt.Printf("[Client %s] Je réponds au docteur sur l’état de mon animal %s.\n", cl.Name, cl.Pet)
}
