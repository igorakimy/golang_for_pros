package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// httptest.NewRecorder() возвращает объект httptest.ResponseRecorder
	// и используется для записи HTTP-ответа.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	// Убедиться, что код ответа соответствует ожидаемому.
	if status != http.StatusOK {
		t.Errorf("handler returned %v", status)
	}

	expect := `Fine!`
	// Убедиться, что тело ответа совпадает с ожидаемым.
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}

func TestStatusNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/StatusNotFound", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusNotFound)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned %v", rr.Code)
	}
}
