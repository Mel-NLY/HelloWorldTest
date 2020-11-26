package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestIncHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/inc?input=2", nil)
	if err != nil {
		t.Fatalf("Cannot create request: %v", err)
	}

	rec := httptest.NewRecorder()

	incHandler(rec, req)

	result := rec.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status OK but got %v", result.Status)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("response body cannot read: %v", err)
	}

	data, err := strconv.Atoi(string(bytes.TrimSpace(body)))
	if err != nil {
		t.Fatalf("expected an integer but got %s", body)
	}

	if data != 3 {
		t.Fatalf("expected result of 3 but got %v", data)
	}
}
