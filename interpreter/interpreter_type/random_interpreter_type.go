package parser_parser_type

import (
	"bytes"
	"fmt"
	"proxychain-config-go/config"
	helper_formatter "proxychain-config-go/helper/formatter"
	"proxychain-config-go/loader"
)

type randomInterpreterType struct {
	opt         config.Options
	tpl         []byte
	ipFormatter helper_formatter.FormatterInterface
}

func NewRandomInterpreterType(opt config.Options, tpl []byte, f helper_formatter.FormatterInterface) InterpreterTypeInterface {
	return &randomInterpreterType{
		opt:         opt,
		tpl:         tpl,
		ipFormatter: f,
	}
}

func (r *randomInterpreterType) Parse() ([]byte, error) {
	for i := 1; i < 7; i++ {
		p := loader.InitRandomProxy()

		lbl := fmt.Sprintf("#{PROXY%d}", i)

		r.ipFormatter.SetData(p)
		fd, err := r.ipFormatter.Exec()
		if err != nil {
			return nil, err
		}

		r.tpl = bytes.Replace(r.tpl, []byte(lbl), []byte(fd.(string)), -1)
	}

	return r.tpl, nil
}
