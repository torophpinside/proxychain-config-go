package parser_parser_type

import (
	"bytes"
	"errors"
	"fmt"
	"proxychain-config-go/config"
	helper_formatter "proxychain-config-go/helper/formatter"
	loader_proxy "proxychain-config-go/loader/proxy"
)

type listInterpreterType struct {
	opt         config.Options
	tpl         []byte
	ipFormatter helper_formatter.FormatterInterface
	proxyLoader loader_proxy.ProxyLoaderInterface
}

func NewListInterpreterType(
	opt config.Options,
	tpl []byte,
	f helper_formatter.FormatterInterface,
	pl loader_proxy.ProxyLoaderInterface,
) InterpreterTypeInterface {
	return &listInterpreterType{
		opt:         opt,
		tpl:         tpl,
		ipFormatter: f,
		proxyLoader: pl,
	}
}

func (l *listInterpreterType) Parse() ([]byte, error) {
	lp, err := l.proxyLoader.Load()
	if err != nil {
		return nil, err
	}
	if len(lp) == 0 {
		return nil, errors.New("no proxy found")
	}
	for i := 1; i <= l.opt.Max; i++ {
		lbl := fmt.Sprintf("#{PROXY%d}", i)
		fd, err := l.ipFormatter.With(lp[i-1]).Format()
		if err != nil {
			return nil, err
		}
		l.tpl = bytes.Replace(l.tpl, []byte(lbl), []byte(fd.(string)), -1)
	}

	return l.tpl, nil
}
