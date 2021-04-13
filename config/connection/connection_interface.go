package config_connection

type ConnectionInterface interface {
	Connect() error
	Get(u string) ([]byte, error)
}
