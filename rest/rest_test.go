package rest

import "errors"
import "io/ioutil"
import "testing"
import "net/http"
import "net/http/httptest"
import "strconv"

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

func TestCountHandler_ExpectedResult(t *testing.T) {
	// given
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ignored/mapping", nil)

	expected := 100
	db := &countryDBStub{result: expected}
	// when
	New(db).countHandler(rec, req)
	// then
	resp := rec.Result()
	defer resp.Body.Close()

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("unexpected response status, want: %v, got: %v", want, got)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unexpected error during reading of response body: %v", err)
	}

	if got, want := string(bytes), strconv.Itoa(expected); got != want {
		t.Errorf("unexpected result, want: %v got: %v", want, got)
	}
}

type countryDBStub struct {
	result       int
	errorMessage string
}

func (c countryDBStub) CountAll() (int, error) {
	var err error = nil
	if c.errorMessage != "" {
		err = errors.New(c.errorMessage)
	}
	return c.result, err
}
