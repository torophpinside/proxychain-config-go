package parser_parser_type

type InterpreterTypeInterface interface {
	Parse() ([]byte, error)
}
