package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	Index(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	payload := struct {
		Message string `json:"message"`
	}{}
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected `200` got %d instead", res.StatusCode)
	}
	if payload.Message != "Hello From Earth!!" {
		t.Errorf("expected `Hello From Earth!!` got %s instead", payload.Message)
	}
}
