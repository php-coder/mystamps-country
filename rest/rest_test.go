package rest

import "errors"
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

func TestCountHandler_DbError(t *testing.T) {
	// given
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ignored/mapping", nil)

	db := &countryDBStub{errorMessage: "expected error from stub"}
	// when
	New(db).countHandler(rec, req)
	// then
	resp := rec.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("unexpected response status: expected %v, got %v", http.StatusInternalServerError, resp.StatusCode)
	}
}

type countryDBStub struct {
	result       int
	errorMessage string
}

func (c countryDBStub) CountAll() (int, error) {
	return c.result, errors.New(c.errorMessage)
}
