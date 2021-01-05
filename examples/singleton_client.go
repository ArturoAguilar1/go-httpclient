package examples

import (
	"github.com/ArturoAguilar1/go-httpclient/gohttp"
	"github.com/ArturoAguilar1/go-httpclient/gomime"
	"net/http"
	"time"
)

var(
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client{
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetHttpClient(nil).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Turonga-Computer").
		Build()
	return client
}
