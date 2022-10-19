package examples

import (
	"net/http"
	"time"

	"github.com/lucaseg/go-httpclient/gohttp"
	"github.com/lucaseg/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetUserAgent("Lucas-Agent").
		SetResponseTimeout(3 * time.Second).
		Build()

	return client
}
