package parser_parser_type

import (
	"bytes"
	"errors"
	"fmt"
	"proxychain-config-go/config"
	helper_formatter "proxychain-config-go/helper/formatter"
	loader_proxy "proxychain-config-go/loader/proxy"
)

type randomInterpreterType struct {
	opt         config.Options
	tpl         []byte
	ipFormatter helper_formatter.FormatterInterface
	proxyLoader loader_proxy.ProxyLoaderInterface
}

func NewRandomInterpreterType(
	opt config.Options,
	tpl []byte,
	f helper_formatter.FormatterInterface,
	pl loader_proxy.ProxyLoaderInterface,
) InterpreterTypeInterface {
	return &randomInterpreterType{
		opt:         opt,
		tpl:         tpl,
		ipFormatter: f,
		proxyLoader: pl,
	}
}

func (r *randomInterpreterType) Parse() ([]byte, error) {
	l, err := r.proxyLoader.Load()
	if err != nil {
		return nil, err
	}
	if len(l) == 0 {
		return nil, errors.New("no proxy found")
	}

	for i := 1; i < len(l); i++ {
		lbl := fmt.Sprintf("#{PROXY%d}", i)
		fd, err := r.ipFormatter.With(l[i-1]).Format()
		if err != nil {
			return nil, err
		}

		r.tpl = bytes.Replace(r.tpl, []byte(lbl), []byte(fd.(string)), -1)
	}

	return r.tpl, nil
}
