package main

import (
	"cours/internal/constants"
	"cours/internal/handler"
	"cours/internal/transport"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Не удалось загрузить .env файл: %v", err)
		log.Println("Проверьте наличие файла .env в корне проекта")
	}

	oauthToken := os.Getenv("YANDEX_OAUTH_TOKEN")
	if oauthToken == "" {
		log.Fatal("Не задан OAuth токен. Установите переменную окружения YANDEX_OAUTH_TOKEN или создайте файл .env")
	}

	if len(os.Args) < 4 {
		handler.PrintUsage()
		return
	}

	keyword := os.Args[1]
	fromDate := os.Args[2]
	toDate := os.Args[3]

	if !handler.IsValidDate(fromDate) || !handler.IsValidDate(toDate) {
		log.Fatal("Неверный формат даты. Используйте YYYY-MM-DD")
	}

	log.Printf("Запрашиваем статистику для слова: '%s'\n", keyword)
	log.Printf("Период: с %s по %s\n\n", fromDate, toDate)

	stats, err := transport.FetchWordstatData(keyword, fromDate, toDate, oauthToken)
	if err != nil {
		log.Fatalf("Ошибка при получении данных: %v", err)
	}

	report := handler.GenerateReport(stats, keyword, fromDate, toDate)

	err = handler.SaveToFile(report, constants.OutputFile)
	if err != nil {
		log.Fatalf("Ошибка при сохранении в файл: %v", err)
	}

	log.Printf("Отчет сохранен в файл: %s\n", constants.OutputFile)
}
