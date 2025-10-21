package transport

import (
	"bytes"
	"cours/internal/constants"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WordstatRequest struct {
	Phrases []string `json:"phrases"`
	Period  Period   `json:"period"`
}

type Period struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

type WordstatResponse struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
}

type Data struct {
	Phrase     string         `json:"phrase"`
	Regions    []RegionStat   `json:"regions"`
	Searches   []SearchStat   `json:"searches"`
	Geo        []GeoStat      `json:"geo"`
	Demography []Demography   `json:"demography"`
	History    []HistoryPoint `json:"history"`
}

type RegionStat struct {
	RegionID int    `json:"regionId"`
	Name     string `json:"name"`
	Shows    int    `json:"shows"`
}

type SearchStat struct {
	Phrase string `json:"phrase"`
	Shows  int    `json:"shows"`
}

type GeoStat struct {
	RegionID int    `json:"regionId"`
	Name     string `json:"name"`
	Shows    int    `json:"shows"`
}

type Demography struct {
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Shows  int    `json:"shows"`
}

type HistoryPoint struct {
	Date  string `json:"date"`
	Shows int    `json:"shows"`
}

func FetchWordstatData(keyword, fromDate, toDate, oauthToken string) (*WordstatResponse, error) {
	requestBody := map[string]interface{}{
		"phrase":  keyword,
		"regions": []int{213, 2},
		"devices": []string{"phone", "tablet", "desktop"},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга JSON: %v", err)
	}

	req, err := http.NewRequest("POST", constants.ApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+oauthToken)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	req.Header.Set("Cookie", "_yasc=+xC4Te9mS5y8O0NjE2VA9NBNr8/FDEi4ZODKg32rM3vhmBJSq3yv50b2iWt27H0=")

	fmt.Printf("Токен: %s...\n", oauthToken[:20])
	fmt.Printf("Отправляем запрос на: %s\n", constants.ApiURL)

	client := &http.Client{
		Timeout: time.Duration(constants.RequestTimeout) * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	fmt.Printf("Получен ответ: %d\n", resp.StatusCode)
	fmt.Printf("Тело ответа: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернуло ошибку: %s\nТело ответа: %s", resp.Status, string(body))
	}

	var wordstatResp WordstatResponse
	err = json.Unmarshal(body, &wordstatResp)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	return &wordstatResp, nil
}
