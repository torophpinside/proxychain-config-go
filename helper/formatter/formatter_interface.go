package helper_formatter

type FormatterInterface interface {
	SetData(d interface{})
	Exec() (interface{}, error)
}
