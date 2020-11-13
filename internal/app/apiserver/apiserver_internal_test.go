package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer(t *testing.T) {
	s := New(NewConfig())
	rec := (httptest.NewRecorder())
	req := (httptest.NewRequest(http.MethodGet, "/hello", nil))
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hi, Dude")
}
