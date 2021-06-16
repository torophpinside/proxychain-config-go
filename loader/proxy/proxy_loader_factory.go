package loader_proxy

import (
	"proxychain-config-go/config"
	_configConnection "proxychain-config-go/config/connection"
	"strings"
)

func Factory(opt config.Options, conn _configConnection.ConnectionInterface) ProxyLoaderInterface {
	bc := _configConnection.NewBasicConnection()
	_ = bc.Connect()

	if strings.ToLower(opt.Src) == "alt" {
		return NewAltProxyLoader(conn)
	}
	if strings.ToLower(opt.Src) == "nova" {
		return NewNovaProxyLoader(conn)
	}
	if strings.ToLower(opt.Src) == "plus" {
		return NewProxyListPlusProxyLoader(bc)
	}
	if strings.ToLower(opt.Src) == "free" {
		return NewFreeProxyListProxyLoader(bc)
	}
	if strings.ToLower(opt.Src) == "spysone" {
		return NewSpysoneProxyLoader(bc)
	}
	if strings.ToLower(opt.Src) == "all" {
		return NewAllProxyLoader(bc, conn)
	}

	return nil
}
