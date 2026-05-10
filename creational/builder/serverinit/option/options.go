package option

import "time"

type Server struct {
	host     string
	port     int
	protocol string
	timeout  time.Duration
	maxConn  int
}

type ServerOption func(*Server)

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

func NewServer(host string, port int, opts ...ServerOption) *Server {
	server := &Server{host: host, port: port}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func WithProtocol(proto string) ServerOption {
	return func(s *Server) {
		s.protocol = proto
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(max int) ServerOption {
	return func(s *Server) {
		s.maxConn = max
	}
}

