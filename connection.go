package gedis

import (
	"fmt"
	"net"
)

func makeAddr(config *Config) string {
	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}

func newConn(config *Config) (net.Conn, error) {
	return net.Dial("tcp", makeAddr(config))
}
