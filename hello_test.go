package main

import (
        "io"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestHelloWorld(t *testing.T) {
        r := httptest.NewRequest(http.MethodGet, "/", nil)
        w := httptest.NewRecorder()

        helloWorld(w, r)

        resp := w.Result()

        if want, got := http.StatusOK, resp.StatusCode; want != got {
                t.Fatalf("Expected a %d response code, got %d", want, got)
        }

        body, _ := io.ReadAll(resp.Body)

        if string(body) != "Hello, world!" {
                t.Errorf(`Expected "Hello, world!", got %s`, string(body))
        }
}
