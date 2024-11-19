package server_http_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	server_http "github.com/VandiKond/parse-ru-time-duration-go/internal/http"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("1 м"))
	w := httptest.NewRecorder()
	Handler := server_http.ParseHandler{Url: ""}
	Handler.ServeHTTP(w, req)
	res := w.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("unreadable body")
	}
	if string(body) != "1m0s" {
		t.Fatalf("expected 1m0s, got %s", string(body))
	}
}
