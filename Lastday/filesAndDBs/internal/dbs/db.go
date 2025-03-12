// internal/dbs/db.go

package dbs

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

// DB est la connexion globale à la base de données.
var DB *sql.DB

// Todo représente la structure de données pour un todo.
type Todo struct {
    ID          int
    Title       string
    Description string
    Category    string
}

// init() s'exécute automatiquement à l'import du package.
// Il initialise la connexion à la base de données.
func init() {
    fmt.Println("Initializing database connection in package dbs...")

    // Configuration de la connexion MySQL.
    cfg := mysql.Config{
        User:                 os.Getenv("DBUSER"),  // ex: export DBUSER=ton_user
        Passwd:               os.Getenv("DBPASS"),  // ex: export DBPASS=ton_mdp
        Net:                  "tcp",
        Addr:                 "127.0.0.1:3306",
        DBName:               "todo_db",            // Nom de ta base (doit exister)
        AllowNativePasswords: true,
    }

    var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("Erreur lors de l'ouverture de la connexion: %v", err)
    }

    // Vérification de la connexion
    if err = DB.Ping(); err != nil {
        log.Fatalf("Impossible de joindre la base de données: %v", err)
    }

    fmt.Println("Connected to the database!")
}

// AddTodo insère un nouveau Todo et renvoie l'ID généré.
func AddTodo(title, description, category string) int {
    result, err := DB.Exec(
        "INSERT INTO todos (title, description, category) VALUES (?, ?, ?)",
        title, description, category,
    )
    if err != nil {
        log.Printf("Erreur lors de l'ajout du todo: %v\n", err)
        return 0
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        log.Printf("Erreur lors de la récupération de l'ID: %v\n", err)
        return 0
    }
    return int(lastID)
}

// EditTodo met à jour un Todo existant. 
// Si newTitle ou newDescription sont vides, on garde les valeurs actuelles.
// La fonction renvoie le Todo mis à jour.
func EditTodo(id int, newTitle, newDescription string) Todo {
    // On récupère d'abord le Todo actuel
    current := GetTodoById(id)
    if current.ID == 0 {
        // Pas trouvé, on renvoie un Todo vide
        return Todo{}
    }

    // Mise à jour uniquement si les nouvelles valeurs ne sont pas vides
    if newTitle != "" {
        current.Title = newTitle
    }
    if newDescription != "" {
        current.Description = newDescription
    }

    _, err := DB.Exec(
        "UPDATE todos SET title = ?, description = ? WHERE id = ?",
        current.Title, current.Description, id,
    )
    if err != nil {
        log.Printf("Erreur lors de la mise à jour du todo: %v\n", err)
        return Todo{}
    }

    return current
}

// RemoveTodo supprime un Todo par son ID et renvoie l'ID supprimé (ou 0 si erreur).
func RemoveTodo(id int) int {
    _, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        log.Printf("Erreur lors de la suppression du todo %d: %v\n", id, err)
        return 0
    }
    return id
}

// ListTodos renvoie la liste de tous les Todos.
func ListTodos() []Todo {
    rows, err := DB.Query("SELECT id, title, description, category FROM todos")
    if err != nil {
        log.Printf("Erreur lors de la récupération des todos: %v\n", err)
        return nil
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var t Todo
        if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Category); err != nil {
            log.Printf("Erreur lors du scan d'un todo: %v\n", err)
            continue
        }
        todos = append(todos, t)
    }
    return todos
}

// GetTodoById renvoie un Todo à partir de son ID (ou un Todo vide si non trouvé).
func GetTodoById(id int) Todo {
    var t Todo
    err := DB.QueryRow(
        "SELECT id, title, description, category FROM todos WHERE id = ?",
        id,
    ).Scan(&t.ID, &t.Title, &t.Description, &t.Category)

    if err == sql.ErrNoRows {
        // Aucun Todo trouvé
        return Todo{}
    } else if err != nil {
        log.Printf("Erreur lors de la récupération du todo %d: %v\n", id, err)
        return Todo{}
    }
    return t
}

// ListCategories renvoie toutes les catégories distinctes.
func ListCategories() []string {
    rows, err := DB.Query("SELECT DISTINCT category FROM todos")
    if err != nil {
        log.Printf("Erreur lors de la récupération des catégories: %v\n", err)
        return nil
    }
    defer rows.Close()

    var categories []string
    for rows.Next() {
        var cat string
        if err := rows.Scan(&cat); err != nil {
            log.Printf("Erreur lors du scan d'une catégorie: %v\n", err)
            continue
        }
        categories = append(categories, cat)
    }
    return categories
}

// GetTodoByCategory renvoie le premier Todo trouvé pour une catégorie donnée.
// (Si plusieurs Todos ont la même catégorie, on ne récupère que le premier.)
func GetTodoByCategory(categoryName string) Todo {
    var t Todo
    err := DB.QueryRow(
        "SELECT id, title, description, category FROM todos WHERE category = ? LIMIT 1",
        categoryName,
    ).Scan(&t.ID, &t.Title, &t.Description, &t.Category)

    if err == sql.ErrNoRows {
        // Aucun Todo trouvé pour cette catégorie
        return Todo{}
    } else if err != nil {
        log.Printf("Erreur lors de la récupération du todo pour la catégorie %s: %v\n", categoryName, err)
        return Todo{}
    }
    return t
}

// RemoveTodoByCategory supprime tous les Todos d'une catégorie et 
// renvoie le nombre de lignes supprimées.
func RemoveTodoByCategory(categoryName string) int {
    result, err := DB.Exec("DELETE FROM todos WHERE category = ?", categoryName)
    if err != nil {
        log.Printf("Erreur lors de la suppression des todos de la catégorie %s: %v\n", categoryName, err)
        return 0
    }

    rows, err := result.RowsAffected()
    if err != nil {
        log.Printf("Erreur lors de la récupération du nombre de lignes supprimées: %v\n", err)
        return 0
    }
    return int(rows)
}
