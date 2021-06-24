package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
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

	yesnoRx := regexp.MustCompile(`^yesno`)
	intnozeroRx := regexp.MustCompile(`intnozero`)
	params := make(map[string]interface{})

	v, isMap := opts.(map[string]interface{})
	if isMap {
		for k, v2 := range v {
			switch t := v2.(type) {
			case string:
				params[k] = url.QueryEscape(t)
			default:
				params[k] = v2
			}
		}
	} else {
		v := reflect.ValueOf(opts)

		if v.IsValid() {
			typeOfS := v.Type()

			for i := 0; i < v.NumField(); i++ {
				name := typeOfS.Field(i).Name
				namecheapStructTag := typeOfS.Field(i).Tag.Get("namecheap")

				switch typeOfS.Field(i).Type.Kind() {
				case reflect.Bool:
					// Namecheap wants certain "boolean" fields in yes|no format.
					// Struct fields are boolean values that get converted to yes|no strings.
					// Using struct tags for now.
					value := v.Field(i).Bool()

					if yesnoRx.MatchString(namecheapStructTag) {
						if value {
							params[name] = "yes"
						} else {
							params[name] = "no"
						}
					} else {
						params[name] = strconv.FormatBool(value)
					}
				case reflect.String:
					value := v.Field(i).String()
					if value != "" {
						params[name] = url.QueryEscape(value)
					}
				case reflect.Int:
					value := v.Field(i).Int()

					// Does not allow a value of zero or less to be set.
					if intnozeroRx.MatchString(namecheapStructTag) {
						if value > 0 {
							params[name] = fmt.Sprintf("%d", value)
						}
					} else {
						params[name] = fmt.Sprintf("%d", value)
					}
				case reflect.Ptr:
					// allows for a 'third' state for boolean. [nil, true, false]
					// for namecheap boolean parameters that are not required and have no default value
					value := v.Field(i).Elem()

					if value.IsValid() {
						switch value.Kind() {
						case reflect.Bool:
							params[name] = strconv.FormatBool(value.Bool())
						}
					}
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
