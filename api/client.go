package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Client struct {
	baseUrl  string
	apiKey   string
	apiUser  string
	clientIP string
}

func NewClient(baseUrl, apiKey, apiUser, clientIP string) *Client {
	return &Client{
		baseUrl:  baseUrl,
		apiKey:   apiKey,
		apiUser:  apiUser,
		clientIP: clientIP,
	}

}

func (c *Client) prepareUrl(command string, opts interface{}) string {
	url := fmt.Sprintf(
		"%s?ApiKey=%s&ApiUser=%s&UserName=%s&ClientIp=%s&Command=%s",
		c.baseUrl,
		c.apiKey,
		c.apiUser,
		c.apiUser,
		c.clientIP,
		command,
	)

	params := make(map[string]string)

	v := reflect.ValueOf(opts)
	if v.IsValid() {
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			value := v.Field(i).String()
			name := typeOfS.Field(i).Name

			if value != "" {
				params[name] = value
			}
		}
	}

	if params != nil {
		for k, v := range params {
			url = fmt.Sprintf(
				"%s&%s=%s",
				url,
				k,
				v,
			)
		}
	}

	return url
}

func (c *Client) do(command string, opts interface{}, value interface{}) {
	url := c.prepareUrl(command, opts)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	xml.Unmarshal(body, &value)
}
