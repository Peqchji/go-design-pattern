package builder_test

import (
	"testing"
	"time"

	"design_pattern/creational/builder/serverinit/builder"

	"github.com/stretchr/testify/assert"
)

func TestServer_ClassicBuilder(t *testing.T) {
	tests := []struct {
		name           string
		build          func() *builder.Server
		expectedServer *builder.Server
	}{
		{
			name: "Basic server",
			build: func() *builder.Server {
				return builder.NewServerBuilder().
					WithHost("localhost").
					WithPort(8080).
					Build()
			},
			expectedServer: builder.MockServer("localhost", 8080, "", 0, 0),
		},
		{
			name: "Full configuration",
			build: func() *builder.Server {
				return builder.NewServerBuilder().
					WithHost("api.example.com").
					WithPort(443).
					WithProtocol("https").
					WithTimeout(30 * time.Second).
					WithMaxConn(500).
					Build()
			},
			expectedServer: builder.MockServer("api.example.com", 443, "https", 30*time.Second, 500),
		},
		{
			name: "Only protocol",
			build: func() *builder.Server {
				return builder.NewServerBuilder().
					WithProtocol("ftp").
					Build()
			},
			expectedServer: builder.MockServer("", 0, "ftp", 0, 0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.build()
			assert.EqualValues(t, tc.expectedServer, s)
		})
	}
}

func FuzzServer_ClassicBuilder(f *testing.F) {
	f.Add("localhost", 8080, "http", int64(1000), 100)

	f.Fuzz(func(t *testing.T, host string, port int, protocol string, timeoutMillis int64, maxConn int) {
		s := builder.NewServerBuilder().
			WithHost(host).
			WithPort(port).
			WithProtocol(protocol).
			WithTimeout(time.Duration(timeoutMillis) * time.Millisecond).
			WithMaxConn(maxConn).
			Build()

		expected := builder.MockServer(
			host,
			port,
			protocol,
			time.Duration(timeoutMillis)*time.Millisecond,
			maxConn,
		)

		assert.EqualValues(t, expected, s)
	})
}