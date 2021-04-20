package loader_proxy

import (
	"proxychain-config-go/config"
	config_connection "proxychain-config-go/config/connection"
	"strings"
)

func Factory(opt config.Options, conn config_connection.ConnectionInterface) ProxyLoaderInterface {
	if strings.ToLower(opt.Src) == "alt" {
		return NewAltProxyLoader(conn)
	}
	if strings.ToLower(opt.Src) == "nova" {
		return NewNovaProxyLoader(conn)
	}
	if strings.ToLower(opt.Src) == "plus" {
		return NewProxyListPlusProxyLoader(conn)
	}

	return nil
}
