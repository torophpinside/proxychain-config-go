package config_connection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type torConnection struct {
	client *http.Client
}

func NewTorConnection() ConnectionInterface {
	return &torConnection{}
}

func (t *torConnection) Connect() error {
	up, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to parse proxy URL: %v\n", err))
	}

	tr := &http.Transport{Proxy: http.ProxyURL(up)}
	c := &http.Client{Transport: tr}

	t.client = c

	return nil
}

func (t *torConnection) Get(u string) ([]byte, error) {
	resp, err := t.client.Get(u)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to issue GET request: %v\n", err))
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read the body: %v\n", err))
	}

	return body, nil
}
