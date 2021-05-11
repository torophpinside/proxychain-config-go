package interpreter

import (
	"proxychain-config-go/config"
	_configConnection "proxychain-config-go/config/connection"
	helper_formatter "proxychain-config-go/helper/formatter"
	parser_parser_type "proxychain-config-go/interpreter/interpreter_type"
	loader_proxy "proxychain-config-go/loader/proxy"
)

func Factory(
	opt config.Options,
	tpl []byte,
	f helper_formatter.FormatterInterface,
	c _configConnection.ConnectionInterface,
) parser_parser_type.InterpreterTypeInterface {
	if opt.SearchType == "list" {
		pl := loader_proxy.Factory(opt, c)
		return parser_parser_type.NewListInterpreterType(opt, tpl, f, pl)
	}

	return nil
}
