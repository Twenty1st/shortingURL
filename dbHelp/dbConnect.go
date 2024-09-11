package dbHelp

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func database() *sql.DB {
    connStr := "user=user123 password=123 dbname=short_urls_db sslmode=disable" // строка подключения
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