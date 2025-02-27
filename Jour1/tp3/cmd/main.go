package main

import (
    "tp3/internal/interfaces"
)

func main() {
    // Exemple d'utilisation
    perso := interfaces.Personnage{Image: "perso.png"}
    bg := interfaces.Background{Image: []string{"bg1.png", "bg2.png"}}

    // Polymorphisme : IDrawable peut être Personnage ou Background
    objets := []interfaces.IDrawable{perso, bg}
    for _, obj := range objets {
        interfaces.RenderObj(obj)
    }
}
