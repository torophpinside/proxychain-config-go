package interpreter

import (
	"proxychain-config-go/config"
	helper_formatter "proxychain-config-go/helper/formatter"
	parser_parser_type "proxychain-config-go/interpreter/interpreter_type"
)

func Factory(opt config.Options, tpl []byte) parser_parser_type.InterpreterTypeInterface {
	if opt.SearchType == "list" {
		return parser_parser_type.NewListInterpreterType(opt, tpl, helper_formatter.NewIPFormatterHelper())
	}

	return nil
}
