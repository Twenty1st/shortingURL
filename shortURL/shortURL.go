package shortURL

import "shortingURL/dbHelp"

func GetShortURL(full_url string, isSaveInDB bool) string {
	short_url := genNewShortUrl()
	
	if isSaveInDB{
		if !dbHelp.InsertNewURL(full_url, short_url){
			return ""
		}
	}
	
	return short_url
}