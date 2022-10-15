package examples

import (
	"net/http"
	"time"

	"github.com/lucaseg/go-httpclient/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
