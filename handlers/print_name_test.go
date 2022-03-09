package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrintName(t *testing.T) {
	const ROUTE = "/print/"
	const EXPECT = "Hi "

	testCases := []struct {
		name   string
		route  string
		expect string
	}{
		{
			name:   "Should return default value if no path provided",
			route:  ROUTE,
			expect: EXPECT + "lad!",
		},
		{
			name:   "1# Should use the provided path instead default one",
			route:  ROUTE + "user",
			expect: EXPECT + "user!",
		},
		{
			name:   "2# Should use the provided path instead default one",
			route:  ROUTE + "123",
			expect: EXPECT + "123!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.route, nil)
			w := httptest.NewRecorder()
			PrintName(w, req)
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
			if payload.Message != tc.expect {
				t.Errorf("expected `%s` got `%s` instead", tc.expect, payload.Message)
			}
		})
	}
}
