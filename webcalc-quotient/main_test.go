package main

import (
	// "encoding/json"

	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = SetupRouter()

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestQuotient1(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=4&y=2")

	assert.Equal(t, http.StatusOK, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":false`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"string":"4 / 2 = 2"`, res)
	assert.True(t, stringMatch)

	answerMatch, _ := regexp.MatchString(`"answer":2`, res)
	assert.True(t, answerMatch)
}

func TestQuotient2(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=-4&y=3")

	assert.Equal(t, http.StatusOK, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":false`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"string":"-4 / 3 = -1"`, res)
	assert.True(t, stringMatch)

	answerMatch, _ := regexp.MatchString(`"answer":-1`, res)
	assert.True(t, answerMatch)
}

func TestQuotient3(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=20&y=-6")

	assert.Equal(t, http.StatusOK, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":false`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"string":"20 / -6 = -3"`, res)
	assert.True(t, stringMatch)

	answerMatch, _ := regexp.MatchString(`"answer":-3`, res)
	assert.True(t, answerMatch)
}

func TestQuotient4(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=-6&y=-9")

	assert.Equal(t, http.StatusOK, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":false`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"string":"-6 / -9 = 0"`, res)
	assert.True(t, stringMatch)

	answerMatch, _ := regexp.MatchString(`"answer":0`, res)
	assert.True(t, answerMatch)
}

func TestError1(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?y=6")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"x parameter not provided"`, res)
	assert.True(t, stringMatch)
}

func TestError2(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=m&y=6")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"x parameter is not a valid integer"`, res)
	assert.True(t, stringMatch)
}

func TestError3(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=2")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"y parameter not provided"`, res)
	assert.True(t, stringMatch)
}

func TestError4(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=2&y=t")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"y parameter is not a valid integer"`, res)
	assert.True(t, stringMatch)
}

func TestError5(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=2&y=0")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"Divisor \(y\) must not be 0"`, res)
	assert.True(t, stringMatch)
}

func TestError6(t *testing.T) {
	w := performRequest(router, "GET", "/quotient?x=2&y=4&z=3")

	assert.Equal(t, http.StatusBadRequest, w.Code)

	res := w.Body.String()

	errorMatch, _ := regexp.MatchString(`"error":true`, res)
	assert.True(t, errorMatch)

	stringMatch, _ := regexp.MatchString(`"message":"Unrecognised parameter\(s\) provided"`, res)
	assert.True(t, stringMatch)
}
