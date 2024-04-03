package gedis

import (
	"net"
)

type pool struct {
	id   int
	conn net.Conn
}
