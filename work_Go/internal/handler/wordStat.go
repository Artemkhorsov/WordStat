package handler

import (
	"cours/internal/constants"
	"cours/internal/transport"
	"fmt"
	"os"
	"strings"
	"time"
)

func PrintUsage() {
	fmt.Println("Утилита для получения статистики Яндекс.Вордстат")
}

func IsValidDate(date string) bool {
	_, err := time.Parse(constants.DateFormat, date)
	return err == nil
}

func GenerateReport(stats *transport.WordstatResponse, keyword, fromDate, toDate string) string {
	var report strings.Builder

	report.WriteString("ОТЧЕТ ПО СТАТИСТИКЕ ЯНДЕКС.ВОРДСТАТ\n")
	report.WriteString(fmt.Sprintf("Ключевое слово: %s\n", keyword))
	report.WriteString(fmt.Sprintf("Период: с %s по %s\n", fromDate, toDate))
	report.WriteString(fmt.Sprintf("Дата генерации: %s\n\n", time.Now().Format(constants.DateTimeFormat)))

	if len(stats.Data) == 0 {
		report.WriteString("Нет данных для отображения\n")
		return report.String()
	}

	data := stats.Data[0]

	report.WriteString("ОБЩАЯ СТАТИСТИКА:\n")
	report.WriteString("-----------------\n")

	if len(data.History) > 0 {
		report.WriteString("\nИСТОРИЯ ПОКАЗОВ ПО ДНЯМ:\n")
		totalShows := 0
		for _, point := range data.History {
			report.WriteString(fmt.Sprintf("  %s: %d показов\n", point.Date, point.Shows))
			totalShows += point.Shows
		}
		report.WriteString(fmt.Sprintf("  Всего показов за период: %d\n", totalShows))
	}

	if len(data.Regions) > 0 {
		report.WriteString(fmt.Sprintf("\nТОП-%d РЕГИОНОВ:\n", constants.MaxRegions))
		for i, region := range data.Regions {
			if i >= constants.MaxRegions {
				break
			}
			report.WriteString(fmt.Sprintf("  %d. %s: %d показов\n", i+1, region.Name, region.Shows))
		}
	}

	if len(data.Searches) > 0 {
		report.WriteString(fmt.Sprintf("\nТОП-%d ПОХОЖИХ ЗАПРОСОВ:\n", constants.MaxSearches))
		for i, search := range data.Searches {
			if i >= constants.MaxSearches {
				break
			}
			report.WriteString(fmt.Sprintf("  %d. %s: %d показов\n", i+1, search.Phrase, search.Shows))
		}
	}

	if len(data.Demography) > 0 {
		report.WriteString("\nДЕМОГРАФИЧЕСКАЯ СТАТИСТИКА:\n")
		for _, demo := range data.Demography {
			report.WriteString(fmt.Sprintf("  %s, %s: %d показов\n", demo.Age, demo.Gender, demo.Shows))
		}
	}

	return report.String()
}

func SaveToFile(content, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func ValidateInput(keyword, fromDate, toDate string) error {
	if keyword == "" {
		return fmt.Errorf("ключевое слово не может быть пустым")
	}

	if !IsValidDate(fromDate) {
		return fmt.Errorf("неверный формат даты начала: %s", fromDate)
	}

	if !IsValidDate(toDate) {
		return fmt.Errorf("неверный формат даты окончания: %s", toDate)
	}

	return nil
}
