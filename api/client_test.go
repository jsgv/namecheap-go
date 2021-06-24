package api

import "net/url"

func testNewApiUrl(method string, v interface{}) (*url.URL, error) {
	apiKey := "theApiKey"
	apiUser := "theApiUser"
	clientIp := "127.0.0.1"

	client := NewClient(
		"https://api.namecheap.com/xml.response",
		apiKey,
		apiUser,
		clientIp,
	)

	u, err := url.Parse(client.prepareUrl(
		method,
		v,
	))

	return u, err
}
