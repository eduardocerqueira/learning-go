package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestHomePage(t *testing.T) {
	expected := []byte("Welcome to the HomePage!")
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	checkError(err, t)

	res := httptest.NewRecorder()
	homePage(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v want 200", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf("Response body was '%v' want '%v'", expected, res.Body)
	}
}

func TestGetAllArticles(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/articles", nil)

	checkError(err, t)

	res := httptest.NewRecorder()
	getAllArticles(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v want 200", res.Code)
	}

	articles_json, err := json.Marshal(generateArticles())
	fmt.Println(string(articles_json))
	fmt.Println(res.Body)

	assert.JSONEq(t, string(articles_json), res.Body.String(), "Response body differs")
}
