package dbHelp

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func database() *sql.DB {
    username := ""
    password := ""
    dbname := "short_urls_db"
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println("Error opening database:", err)
        return nil
    }

    err = db.Ping()
    if err != nil {
        log.Println("Error connecting to database:", err)
        return nil
    }

    return db
}