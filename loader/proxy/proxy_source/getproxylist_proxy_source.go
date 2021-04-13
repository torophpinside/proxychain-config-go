package loader_proxy_proxy_source

import (
	"encoding/json"
	"fmt"
	"log"
	config_connection "proxychain-config-go/config/connection"
)

type getProxyData struct {
	AllowsHTTPS bool   `json:"allowsHttps"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
}

type getProxyListProxySource struct {
	conn config_connection.ConnectionInterface
}

func NewGetProxyListProxySource(c config_connection.ConnectionInterface) ProxySourceInterface {
	return &getProxyListProxySource{
		conn: c,
	}
}

func (g *getProxyListProxySource) Get() (string, error) {
	pd := getProxyData{}

	body, err := g.conn.Get("https://api.getproxylist.com/proxy?minDownloadSpeed=300&maxSecondsToFirstByte=10&allowsPost=1&maxConnectTime=15&protocol[]=http&allowsUserAgentHeader=1&lastTested=800")
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &pd)
	if err != nil {
		return "", err
	}

	if pd.IP == "" {
		log.Println("no random proxy found (getproxy)")
		return "://", nil
	}

	sec := "http"
	return fmt.Sprintf("%s://%s:%d", sec, pd.IP, pd.Port), nil
}
