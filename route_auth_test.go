package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Login(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/login", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	if strings.Contains(body, "Sign in") == false {
		t.Errorf("Body does not contain Sign in")
	}
}

func Test_Get_Signup(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc( "/signup", signup)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/signup", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	if strings.Contains(body, "Sign up") == false {
		t.Errorf("Body does not contain Sign up")
	}
}

func Test_Post_SignupAccount(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup_account", signupAccount)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"name": "bob", "email": "123456@qq.com", "password": "123456"}`)
	request, _ := http.NewRequest("POST", "/signup_account", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 302 {
		t.Errorf("Response code is %v", writer.Code)
	}
	url := writer.Header().Get("Location")
	if strings.EqualFold(url, "/login") == false {
		t.Errorf("Not redirect to /login")
		t.Errorf(url)
	}
}

func Test_Post_Authenticate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/authenticate", authenticate)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"email": "1234567@qq.com", "password": "123456"}`)
	request, _ := http.NewRequest("POST", "/authenticate", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 302 {
		t.Errorf("Response code is %v", writer.Code)
	}
	url := writer.Header().Get("Location")
	if strings.EqualFold(url, "/login") == true {
		t.Errorf("login failed")
	} else if strings.EqualFold(url, "/") == true{
		fmt.Println("Login successfully")
	} else {
		t.Errorf("redirect to error location")
	}
}

