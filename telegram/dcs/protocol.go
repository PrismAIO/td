package dcs

import (
	"net"

	"github.com/PrismAIO/td/transport"
)

type protocol interface {
	Handshake(conn net.Conn) (transport.Conn, error)
}
