package interfaces

import "fmt"

// IDrawable est l'interface pour tout ce qui "se dessine"
type IDrawable interface {
    Draw()
}

// Personnage est un exemple de struct qui implémente IDrawable
type Personnage struct {
    Image string
}

func (p Personnage) Draw() {
    fmt.Println("Personnage avec l'image:", p.Image)
}

// Background est un autre struct implémentant IDrawable
type Background struct {
    Image []string
}

func (b Background) Draw() {
    fmt.Println("Background avec les images:", b.Image)
}

// RenderObj appelle la méthode Draw() de n'importe quel IDrawable
func RenderObj(obj IDrawable) {
    obj.Draw()
}
