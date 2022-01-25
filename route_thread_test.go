package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var cookie = http.Cookie{
	Name: "cookie",
	Value: "c2222f79-0011-495f-46ac-b69f315582ea",
	HttpOnly: true,
}

func Test_Get_newThread(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/threads/new", newThread)

	writer := httptest.NewRecorder()

	request := httptest.NewRequest("GET", "/threads/new", nil)
	request.AddCookie(&cookie)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
