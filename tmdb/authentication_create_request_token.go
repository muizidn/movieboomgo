package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type RespCreateRequestToken struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}

type RespCreateRequestTokenErr struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

func CreateRequestToken(apiKey string) (*RespCreateRequestToken, *RespCreateRequestTokenErr, error) {

	url := "https://api.themoviedb.org/3/authentication"

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("tmdb_create_request_token",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	client := &http.Client{}

	resp, err := client.Get(fmt.Sprintf("https://api.themoviedb.org/3/authentication/token/new?api_key=%s", apiKey))

	if err != nil {
		sugar.Errorf("Failure", err)
	}

	sugar.Infof("Response statusCode %d", resp.StatusCode)

	respBody, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200:
		var result RespCreateRequestToken
		json.Unmarshal(respBody, &result)
		return &result, nil, nil
	case 401:
		var result RespCreateRequestTokenErr
		json.Unmarshal(respBody, &result)
		return nil, &result, nil
	default:
		return nil, nil, fmt.Errorf("TMDB StatusCode: %d", resp.StatusCode)
	}
}
