package config_connection

import (
	"errors"
	"fmt"
	"io/ioutil"
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
