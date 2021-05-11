package config_connection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type basicConnection struct {
	client *http.Client
}

func NewBasicConnection() ConnectionInterface {
	return &basicConnection{}
}

func (t *basicConnection) Connect() error {
	c := &http.Client{}
	t.client = c
	return nil
}

func (t *basicConnection) Get(u string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0")
	resp, err := t.client.Do(req)
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
