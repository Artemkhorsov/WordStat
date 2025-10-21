package constants

const (
	ApiURL = "https://api.wordstat.yandex.net/v1/topRequests"

	OutputFile     = "wordstat_report.txt"
	RequestTimeout = 30
	MaxRegions     = 10
	MaxSearches    = 15
)
const (
	ErrNoOAuthToken      = "Не задан OAuth токен. Установите переменную окружения YANDEX_OAUTH_TOKEN"
	ErrInvalidDateFormat = "Неверный формат даты. Используйте YYYY-MM-DD"
	ErrAPIStatusNotOK    = "статус API не 'ok': %s"
)
const (
	ContentTypeJSON   = "application/json"
	AuthorizationType = "OAuth"
	UserAgent         = "Wordstat-Tool/1.0"
)
const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)
