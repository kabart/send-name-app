package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCountChar(t *testing.T) {

	var testNames = []struct {
		name  string
		chars int
	}{
		{"Kasia", 5},
		{"Michał", 6},
		{"Róża", 4},
		{"Jan Maria", 8},
	}

	for _, testCase := range testNames {

		if i := countChar(testCase.name); i != testCase.chars {
			t.Errorf("\nNumber of characters in name %s - expected: %d\n Got: %d\n", testCase.name, testCase.chars, i)
		}
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestHealthCheckHandler(t *testing.T) {

	request, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheck)
	handler.ServeHTTP(rec, request)

	checkResponseCode(t, http.StatusOK, rec.Code)
}

func TestSendNameHandler(t *testing.T) {

	var testCases = []struct {
		input  string
		output string
	}{
		{
			"name=Katarzyna",
			"Ilość znaków w imieniu Katarzyna: 9\n",
		},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		request, err := http.NewRequest("POST", "/send-name", reader)
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(sendName)
		handler.ServeHTTP(rec, request)

		checkResponseCode(t, http.StatusOK, rec.Code)

		if got := rec.Body.String(); got != testCase.output {
			t.Errorf("\nExpected: %s\n Got: %s\n", testCase.output, rec.Body.String())
		}
	}
}

func checkResponseCode(t *testing.T, expected, actual int, testName ...string) {
	if expected != actual {
		t.Errorf("Error in test%s: expected response code %d. Got %d\n", testName, expected, actual)
	}
}
