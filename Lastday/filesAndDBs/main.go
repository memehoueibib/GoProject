// main.go

package main

import (
    "fmt"

    "lastday/filesAndDBs/internal/dbs"
)

func main() {
    // 1. Ajouter un Todo
    newID := dbs.AddTodo("Apprendre Go", "Coder un petit projet", "Dev")
    fmt.Println("Nouveau Todo ID =", newID)

    // 2. Modifier un Todo (on met juste le titre si la description reste vide)
    updated := dbs.EditTodo(newID, "Apprendre Go en profondeur", "")
    fmt.Printf("Todo modifié: %+v\n", updated)

    // 3. Lister tous les Todos
    todos := dbs.ListTodos()
    fmt.Println("Liste des todos:")
    for _, t := range todos {
        fmt.Printf("- [%d] %s (%s)\n", t.ID, t.Title, t.Category)
    }

    // 4. Récupérer un Todo par son ID
    one := dbs.GetTodoById(newID)
    fmt.Printf("Récupéré par ID = %d: %+v\n", newID, one)

    // 5. Lister toutes les catégories
    categories := dbs.ListCategories()
    fmt.Println("Catégories disponibles:", categories)

    // 6. Récupérer un Todo par sa catégorie
    firstOfCategory := dbs.GetTodoByCategory("Dev")
    fmt.Printf("Premier Todo de la catégorie 'Dev': %+v\n", firstOfCategory)

    // 7. Supprimer un Todo
    removedID := dbs.RemoveTodo(newID)
    fmt.Printf("Todo %d supprimé\n", removedID)

    // 8. Supprimer tous les Todos d'une catégorie
    count := dbs.RemoveTodoByCategory("Dev")
    fmt.Printf("Nombre de Todos supprimés dans la catégorie 'Dev': %d\n", count)
}
