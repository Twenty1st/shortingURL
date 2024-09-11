package main

import (
	"flag"
	"shortingURL/dbHelp"
	"shortingURL/shortURL"
	"shortingURL/temporaryData"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Определение флага -d
	useDatabase := flag.Bool("d", false, "Use PostgreSQL for storage")
	flag.Parse()

	// Обработка post  запроса
	app.Post("/", func(c *fiber.Ctx) error {
		url := string(c.Body())
		if !IsCorrectUrlAddress(url){
			return c.SendString(
				"Error: incorrect url address",
			)
		}

		// Получаем короткий url
		short_url := shortURL.GetShortURL(url, *useDatabase)
		// Если произошла ошибка при создании короткого url, выводим сообщение об ошибке
		if short_url == "" {
				return c.SendString(
					"Error: error for create short url",
				)
		}
		// Сохраняем во временном словаре
		temporaryData.AddURL(short_url, url)
		// Выводим короткий url
		return c.SendString(
				"url: "	+ url + "\n" +
				"shortURL: http://127.0.0.1:3000/"+short_url,
		)
	})

	// Обработка get запроса, вывод полного url
	app.Get("/:short_url", func(c *fiber.Ctx) error {
		shortURL := c.Params("short_url")
		full_url, check := temporaryData.GetFullURL(shortURL)

		if check{
			return c.SendString(full_url)
		}else{
			full_url = dbHelp.SelectFullUrl(shortURL)
			if(full_url == ""){
				return c.SendString("Don`t know this url!")
			}
			return c.SendString(full_url)
		}
	})

	// Обработка get запроса с перенаправлением по указанному url
	// app.Get("/:short_url", func(c *fiber.Ctx) error {
	// 	shortURL := c.Params("short_url")
	// 	full_url, check := temporaryData.GetFullURL(shortURL)

	// 	if check {
	// 			return c.Redirect(full_url, fiber.StatusFound)
	// 	} else {
	// 			full_url = dbHelp.SelectFullUrl(shortURL)
	// 			if full_url == "" {
	// 					return c.SendString("Don't know this url!")
	// 			}
	// 			return c.Redirect(full_url, fiber.StatusFound)
	// 	}
	// })


	app.Listen(":3000")
}