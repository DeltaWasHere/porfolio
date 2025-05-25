package client

import (
	"net/http"
	"time"
)

const(
	BaseURLV2 = "https://api.github.com/"
)

type GitApiClient struct{
	BaseUrl string
	HTTPClient *http.Client
	apiKey string
}

func NewClient() *GitApiClient{
	return &GitApiClient{
		BaseUrl: BaseURLV2,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}



