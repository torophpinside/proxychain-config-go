package loader_proxy_proxy_source

import (
	"encoding/json"
	"errors"
	"fmt"
	config_connection "proxychain-config-go/config/connection"
	"strings"
)

type gimmeProxyData struct {
	IPPort string `json:"ipPort"`
	Type   string `json:"type"`
}

type gimmeProxyProxySource struct {
	conn config_connection.ConnectionInterface
}

func NewGimmeProxyProxySource(c config_connection.ConnectionInterface) ProxySourceInterface {
	return &gimmeProxyProxySource{
		conn: c,
	}
}

func (g *gimmeProxyProxySource) Get() (string, error) {
	pd := gimmeProxyData{}

	body, err := g.conn.Get("https://gimmeproxy.com/api/getProxy?post=true&anonymityLevel=1&protocol=http")
	if err != nil {
		return "", err
	}

	if strings.Contains(string(body), "No proxy") {
		return "", errors.New("no random proxy found (no proxy gimmeproxy)")
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s", pd.Type, pd.IPPort), nil
}
