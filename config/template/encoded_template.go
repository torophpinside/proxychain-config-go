package config_template

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type encodedTemplate struct {
	encodedTemplateData string
}

func NewEncodedTemplate() TemplateInterface {
	return &encodedTemplate{
		encodedTemplateData: "ZHluYW1pY19jaGFpbgpwcm94eV9kbnMgCnRjcF9yZWFkX3RpbWVfb3V0IDE1MDAwCnRjcF9jb25uZWN0X3RpbWVfb3V0IDgwMDAKW1Byb3h5TGlzdF0KI3tQUk9YWTF9CiN7UFJPWFkyfQoje1BST1hZM30KI3tQUk9YWTR9CiN7UFJPWFk1fQoje1BST1hZNn0KI3tQUk9YWTd9CiN7UFJPWFk4fQoje1BST1hZOX0KI3tQUk9YWTEwfQoje1BST1hZMTF9CiN7UFJPWFkxMn0KI3tQUk9YWTEzfQoje1BST1hZMTR9CiN7UFJPWFkxNX0KI3tQUk9YWTE2fQoje1BST1hZMTd9CiN7UFJPWFkxOH0KI3tQUk9YWTE5fQ==",
	}
}

func (et *encodedTemplate) GetTemplate() ([]byte, error) {
	sDec, err := base64.StdEncoding.DecodeString(et.encodedTemplateData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error decoding tpl: %s", err))
	}
	return []byte(sDec), nil
}
