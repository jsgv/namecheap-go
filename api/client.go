package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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
	u := fmt.Sprintf(
		"%s?ApiKey=%s&ApiUser=%s&UserName=%s&ClientIp=%s&Command=%s",
		c.baseUrl,
		c.apiKey,
		c.apiUser,
		c.apiUser,
		c.clientIP,
		command,
	)

	params := make(map[string]interface{})

	v, isMap := opts.(map[string]interface{})
	if isMap {
		params = v
	} else {
		v := reflect.ValueOf(opts)

		if v.IsValid() {
			typeOfS := v.Type()

			for i := 0; i < v.NumField(); i++ {
				name := typeOfS.Field(i).Name

				switch typeOfS.Field(i).Type.Kind() {
				case reflect.Bool:
					value := v.Field(i).Bool()
					params[name] = strconv.FormatBool(value)
				case reflect.String:
					value := v.Field(i).String()
					if value != "" {
						params[name] = url.QueryEscape(value)
					}
				case reflect.Int:
					value := v.Field(i).Int()
					params[name] = fmt.Sprintf("%d", value)
				}
			}
		}
	}

	for k, v := range params {
		u = fmt.Sprintf(
			"%s&%s=%s",
			u,
			k,
			v,
		)
	}

	return u
}

func (c *Client) do(command string, opts interface{}, value interface{}) error {
	url := c.prepareUrl(command, opts)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return xml.Unmarshal(body, &value)
}
