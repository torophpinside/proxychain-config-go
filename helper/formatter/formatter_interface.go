package helper_formatter

type FormatterInterface interface {
	With(d interface{}) FormatterInterface
	Format() (interface{}, error)
}
