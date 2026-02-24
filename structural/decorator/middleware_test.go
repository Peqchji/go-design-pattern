package middleware_test

import (
	"design_pattern/pkg"
	"design_pattern/structural/decorator/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StatusRecorderTestcase struct {
	name     string
	capture  int
	expected pkg.Result[int]
}

func TestStatusRecorder_CapturesCode(t *testing.T) {
	table := []StatusRecorderTestcase{
		{
			name:    "Capture status 200",
			capture: http.StatusOK,
			expected: pkg.Result[int]{
				Result: http.StatusOK,
				Error:  nil,
			},
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			dummyWriter := httptest.NewRecorder()
			s := middleware.NewStatusRecorder(
				dummyWriter,
				tc.capture,
			)

			s.WriteHeader(tc.expected.Result)

			assert.Equal(t, tc.expected.Result, s.GetRecord())
		})
	}

}

func TestLoggingMiddleware_Integration(t *testing.T) {
	realHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type", "application/json")
		w.WriteHeader(http.StatusTeapot)

		w.Write([]byte(`{"message": "I am a teapot"}`))
	})

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tea", nil)

	handler := middleware.LoggingMiddleware(realHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTeapot {
		t.Errorf("Client received wrong code: got %v, want %v", rr.Code, http.StatusTeapot)
	}

	expectedBody := `{"message": "I am a teapot"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Client received wrong body: got %v, want %v", rr.Body.String(), expectedBody)
	}

	if rr.Header().Get("X-Content-Type") != "application/json" {
		t.Errorf("Client missing headers")
	}
}
