package interpreter

import (
	"proxychain-config-go/config"
	_configConnection "proxychain-config-go/config/connection"
	_helperFormatter "proxychain-config-go/helper/formatter"
	_parserType "proxychain-config-go/interpreter/interpreter_type"
	_loaderProxy "proxychain-config-go/loader/proxy"
)

func Factory(
	opt config.Options,
	tpl []byte,
	f _helperFormatter.FormatterInterface,
	c _configConnection.ConnectionInterface,
) _parserType.InterpreterTypeInterface {
	if opt.SearchType == "list" {
		pl := _loaderProxy.Factory(opt, c)
		return _parserType.NewListInterpreterType(opt, tpl, f, pl)
	}

	return nil
}
