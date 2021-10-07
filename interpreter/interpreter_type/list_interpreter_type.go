package parser_parser_type

import (
	"bytes"
	"errors"
	"fmt"
	"proxychain-config-go/config"
	_helperFormatter "proxychain-config-go/helper/formatter"
	_loaderProxy "proxychain-config-go/loader/proxy"
)

type listInterpreterType struct {
	opt         config.Options
	tpl         []byte
	ipFormatter _helperFormatter.FormatterInterface
	proxyLoader _loaderProxy.ProxyLoaderInterface
}

func NewListInterpreterType(
	opt config.Options,
	tpl []byte,
	f _helperFormatter.FormatterInterface,
	pl _loaderProxy.ProxyLoaderInterface,
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
	m := l.opt.Max
	if m > len(lp) {
		m = len(lp)
	}
	for i := 1; i <= m; i++ {
		lbl := fmt.Sprintf("#{PROXY%d}", i)
		fd, err := l.ipFormatter.With(lp[i-1]).Format()
		if err != nil {
			return nil, err
		}
		l.tpl = bytes.Replace(l.tpl, []byte(lbl), []byte(fd.(string)), -1)
	}

	return l.tpl, nil
}
