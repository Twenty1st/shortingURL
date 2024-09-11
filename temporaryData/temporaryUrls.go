package temporaryData

var timesShortURL = make(map[string]string)

// Функция для добавления временных URL
func AddURL(shortURL, fullURL string) {
	timesShortURL[shortURL] = fullURL
}

// Функция для получения изначального URL по короткому
func GetFullURL(shortURL string) (string, bool) {
	fullURL, exists := timesShortURL[shortURL]
	return fullURL, exists
}

// Функция для получения всех коротких URL
func GetAllShortURLs() map[string]string {
	return timesShortURL
}