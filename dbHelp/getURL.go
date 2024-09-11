package dbHelp

import (
	"database/sql"
	"log"
)

func SelectFullUrl(short_url string) string {
	db := database()
    if db == nil {
        return ""
    }
    defer db.Close()

    query := "SELECT url_full FROM public.url_addresses WHERE url_short = $1"
    var url_full string

    err := db.QueryRow(query, short_url).Scan(&url_full)
    if err != nil {
        if err == sql.ErrNoRows {
            return ""
        }
        log.Fatal("Error select:", err)
        return ""
    }

    return url_full
}

func SelectShortUrl(full_url string) string {
	db := database()
    if db == nil {
        return ""
    }
    defer db.Close()

    query := "SELECT url_short FROM public.url_addresses WHERE url_full = $1"
    var url_short string

    err := db.QueryRow(query, full_url).Scan(&url_short)
    if err != nil {
        if err == sql.ErrNoRows {
            return ""
        }
        log.Fatal("Error select:", err)
        return ""
    }

    return url_short
}