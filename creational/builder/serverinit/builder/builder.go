package builder

import "time"

type Server struct {
	host     string
	port     int
	protocol string
	timeout  time.Duration
	maxConn  int
}

func MockServer(
	host string,
	port int,
	protocol string,
	timeout time.Duration,
	maxConn int,
) *Server {
	return &Server{
		host:     host,
		port:     port,
		protocol: protocol,
		timeout:  timeout,
		maxConn:  maxConn,
	}
}

type ServerBuilder struct {
	server *Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		server: &Server{},
	}
}

func (b *ServerBuilder) WithHost(host string) *ServerBuilder {
	b.server.host = host
	return b
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.server.port = port
	return b
}

func (b *ServerBuilder) WithProtocol(proto string) *ServerBuilder {
	b.server.protocol = proto
	return b
}

func (b *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	b.server.timeout = timeout
	return b
}

func (b *ServerBuilder) WithMaxConn(max int) *ServerBuilder {
	b.server.maxConn = max
	return b
}

func (b *ServerBuilder) Build() *Server {
	return b.server
}
