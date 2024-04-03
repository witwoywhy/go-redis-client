package gedis

const (
	DefaultConnectionPool = 1
)

type Config struct {
	Host           string
	Port           string
	Password       string
	DB             int
	ConnectionPool int
}
