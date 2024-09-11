package shortURL

import (
	"math/rand"
	"shortingURL/dbHelp"
	"shortingURL/temporaryData"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genNewShortUrl() string{
	var short_url string
	for {
		short_url += string(charset[rand.Intn(len(charset))])
		if len(short_url) >= 6{
			if !IsExistUrl(short_url){
				break
			}
		}
	}
	return short_url
}

func IsExistUrl(short_url string) bool {	
	_, isExist := temporaryData.GetFullURL(short_url)
	if isExist{
		return true
	}
	if dbHelp.SelectShortUrl(short_url) != ""{
		return true;
	}
	return false
}