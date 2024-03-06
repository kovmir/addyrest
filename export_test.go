package addyrest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewCustomClient(context.Background(), server.URL, "test_token")

	return func() {
		server.Close()
	}
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
