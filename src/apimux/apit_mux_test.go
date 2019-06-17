package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func checkHttpStatus(gotStatus int, wantStatus int, t *testing.T) {
	if gotStatus != wantStatus {
		t.Errorf("handler returned wrong result. got %v want %v", gotStatus, wantStatus)
	}
}

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.ServeHTTP(rr, req)

	checkHttpStatus(rr.Code, 200, t)
}
