package dbHelp

import (
	"log"
)

func InsertNewURL(full_url, short_url string) bool {
	db := database()
	if db == nil{
		return false
	}

	query := "INSERT INTO public.url_addresses (url_full, url_short) VALUES ($1, $2)"
	_, err := db.Exec(query, full_url, short_url)
	if err != nil {
		// Обработка ошибки
		log.Fatal("Error inserting:", err)
		return false
	}
	defer db.Close()
	db.Close()
	return true
}