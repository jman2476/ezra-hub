package apicaller

import (
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Client struct {
	httpClient  http.Client
	baseURL     string
	token       string
	refresh     string
	lastRefresh time.Time
}

func NewClient(timeout time.Duration) Client {
	godotenv.Load("./app/client/.env")
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: os.Getenv("SERVER_URL"),
	}
}
