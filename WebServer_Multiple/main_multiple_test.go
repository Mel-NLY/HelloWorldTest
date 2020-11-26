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
	testSuite := []struct {
		name   string
		value  string
		inc    int
		status int
		err    string
	}{
		{name: "increment of two", value: "2", inc: 3},
		{name: "missing a value", value: "", err: "missing value"},
		{name: "not a number", value: "WW", err: "not a number: WW"},
	}

	for _, testCase := range testSuite {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/inc?input="+testCase.value, nil)
			if err != nil {
				t.Fatalf("Cannot create request: %v", err)
			}
			rec := httptest.NewRecorder()
			incHandler(rec, req)

			result := rec.Result()
			defer result.Body.Close()

			body, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatalf("response body cannot read: %v", err)
			}

			if testCase.err != "" {
				if result.StatusCode != http.StatusBadRequest {
					t.Errorf("expected bad request; got %v", result.StatusCode)
				}
				if msg := string(bytes.TrimSpace(body)); msg != testCase.err {
					t.Errorf("expected message %q but got %q", testCase.err, msg)
				}
				return
			}

			if result.StatusCode != http.StatusOK {
				t.Errorf("expected status OK but got %v", result.Status)
			}

			data, err := strconv.Atoi(string(bytes.TrimSpace(body)))
			if err != nil {
				t.Fatalf("expected an integer but got %s", body)
			}

			if data != testCase.inc {
				t.Fatalf("expected result of %v but got %v", testCase.inc, data)
			}
		})
	}
}
