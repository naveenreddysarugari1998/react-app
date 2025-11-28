package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var db *sql.DB

func main() {
    conn := "host=postgres user=postgres password=password dbname=users sslmode=disable"
    var err error

    db, err = sql.Open("postgres", conn)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/user", createUser)
    http.HandleFunc("/users", getUsers)

    fmt.Println("Backend running on :8080")
    http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    var u User
    json.NewDecoder(r.Body).Decode(&u)

    _, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", u.Name, u.Email)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Write([]byte(`{"status":"user added"}`))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    users := []User{}
    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }

    json.NewEncoder(w).Encode(users)
}
