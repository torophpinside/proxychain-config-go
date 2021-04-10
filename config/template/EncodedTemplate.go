package config_template

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

type encodedTemplate struct {
	encodedTemplateData string
}

func NewEncodedTemplate() TemplateInterface {
	return &encodedTemplate{
		encodedTemplateData: "ZHluYW1pY19jaGFpbgpwcm94eV9kbnMgCnRjcF9yZWFkX3RpbWVfb3V0IDE1MDAwCnRjcF9jb25uZWN0X3RpbWVfb3V0IDgwMDAKW1Byb3h5TGlzdF0KI3tQUk9YWTF9CiN7UFJPWFkyfQoje1BST1hZM30KI3tQUk9YWTR9CiN7UFJPWFk1fQoje1BST1hZNn0KI3tQUk9YWTd9CiN7UFJPWFk4fQoje1BST1hZOX0=",
	}
}

func (et *encodedTemplate) GetTemplate() ([]byte, error) {
	sDec, err := base64.StdEncoding.DecodeString(et.encodedTemplateData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error decoding tpl: %s", err))
	}
	tpl := []byte(sDec)
}
