package helper_formatter

import (
	"fmt"
	"strings"
)

type ipFormatterHelper struct {
	data string
}

func NewIPFormatterHelper() FormatterInterface {
	return &ipFormatterHelper{}
}

func (i *ipFormatterHelper) With(d interface{}) FormatterInterface {
	i.data = d.(string)
	return i
}

func (i *ipFormatterHelper) Format() (interface{}, error) {
	spl := strings.Split(i.data, ":")
	ip := spl[1]
	ip = ip[2:len(spl[1])]
	return fmt.Sprintf("%s    %s    %s", spl[0], ip, spl[2]), nil
}
