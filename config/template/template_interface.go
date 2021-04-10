package config_template

type TemplateInterface interface {
	GetTemplate() ([]byte, error)
}
