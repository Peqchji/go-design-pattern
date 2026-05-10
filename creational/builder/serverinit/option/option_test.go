package option_test

import (
	"design_pattern/creational/builder/serverinit/option"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_FunctionalOptions(t *testing.T) {

	tests := []struct {
		name           string
		host           string
		port           int
		options        []option.ServerOption
		expectedServer *option.Server
	}{
		{
			name:           "Basic server",
			host:           "localhost",
			port:           8080,
			options:        []option.ServerOption{},
			expectedServer: option.MockServer("localhost", 8080, "", 0, 0),
		},
		{
			name: "Server with options",
			host: "localhost",
			port: 8080,
			options: []option.ServerOption{
				option.WithProtocol("http"),
				option.WithTimeout(10*time.Second),
				option.WithMaxConn(100),
			},
			expectedServer: option.MockServer("localhost", 8080, "http", 10*time.Second, 100),
		},
		{
			name: "Server with protocol options",
			host: "localhost",
			port: 8080,
			options: []option.ServerOption{
				option.WithProtocol("http"),
			},
			expectedServer: option.MockServer("localhost", 8080, "http", 0, 0),
		},
		{
			name: "Server with timeout options",
			host: "localhost",
			port: 8080,
			options: []option.ServerOption{
				option.WithTimeout(10*time.Second),
			},
			expectedServer: option.MockServer("localhost", 8080, "", 10*time.Second, 0),
		},
		{
			name: "Server with max options",
			host: "localhost",
			port: 8080,
			options: []option.ServerOption{
				option.WithMaxConn(100),
			},
			expectedServer: option.MockServer("localhost", 8080, "", 0, 100),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := option.NewServer(tt.host, tt.port, tt.options...)

			assert.NotNil(t, s)
			assert.Equal(t, tt.expectedServer, s)
		})
	}
}

func FuzzServer_FunctionalOptions(f *testing.F) {
	f.Add("localhost", 8080, "http", int64(1000), 100)

	f.Fuzz(
		func(t *testing.T, host string, port int, protocol string, timeoutMillis int64, maxConn int) {
			options := []option.ServerOption{
				option.WithProtocol(protocol),
				option.WithTimeout(time.Duration(timeoutMillis) * time.Millisecond),
				option.WithMaxConn(maxConn),
			}

			s := option.NewServer(host, port, options...)

			if s == nil {
				t.Errorf("Server should not be nil")
			}
		},
	)
}
