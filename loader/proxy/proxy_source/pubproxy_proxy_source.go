package loader_proxy_proxy_source

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	config_connection "proxychain-config-go/config/connection"
	"strings"
)

type proxyItem struct {
	IPPort string `json:"ipPort"`
	Type   string `json:"type"`
}

type proxyData struct {
	Data  []proxyItem `json:"data"`
	Count int         `json:"count"`
}

type pubProxyProxySource struct {
	conn config_connection.ConnectionInterface
}

func NewPubProxyProxySource(c config_connection.ConnectionInterface) ProxySourceInterface {
	return &pubProxyProxySource{
		conn: c,
	}
}

func (p *pubProxyProxySource) Get() (string, error) {
	pd := proxyData{}

	body, err := p.conn.Get("http://pubproxy.com/api/proxy?user_agent=true&speed=3&type=http&last_check=600&format=json")
	if err != nil {
		return "", err
	}

	if strings.Contains(string(body), "No proxy") {
		return "", errors.New("no random proxy found (no proxy pubproxy)")
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		return "", err
	}

	if pd.Count == 0 {
		log.Println("no random proxy found (pubproxy)")
		return "://", nil
	}

	return fmt.Sprintf("%s://%s", pd.Data[0].Type, pd.Data[0].IPPort), nil
}
