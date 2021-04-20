package interpreter

import (
	"proxychain-config-go/config"
	config_connection "proxychain-config-go/config/connection"
	helper_formatter "proxychain-config-go/helper/formatter"
	parser_parser_type "proxychain-config-go/interpreter/interpreter_type"
	loader_proxy "proxychain-config-go/loader/proxy"
	loader_proxy_proxy_source "proxychain-config-go/loader/proxy/proxy_source"
)

func Factory(
	opt config.Options,
	tpl []byte,
	f helper_formatter.FormatterInterface,
	c config_connection.ConnectionInterface,
) parser_parser_type.InterpreterTypeInterface {
	if opt.SearchType == "list" {
		pl := loader_proxy.Factory(opt, c)
		return parser_parser_type.NewListInterpreterType(opt, tpl, f, pl)
	}
	if opt.SearchType == "random" {
		ps := []loader_proxy_proxy_source.ProxySourceInterface{
			loader_proxy_proxy_source.NewGimmeProxyProxySource(c),
			loader_proxy_proxy_source.NewGetProxyListProxySource(c),
			loader_proxy_proxy_source.NewPubProxyProxySource(c),
		}
		pl := loader_proxy.NewRandomProxyLoader(ps)
		return parser_parser_type.NewRandomInterpreterType(opt, tpl, f, pl)
	}

	return nil
}
