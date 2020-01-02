package rest

import "testing"
import "net/http"
import "net/http/httptest"

func TestCountHandler_InvalidMethod(t *testing.T) {
	// given
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/ignored/mapping", nil)

	// when
	New(nil).countHandler(rec, req)

	// then
	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("unexpected response status: expected %v, got %v", http.StatusMethodNotAllowed, resp.StatusCode)
	}
}
