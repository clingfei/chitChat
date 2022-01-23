package main

import (
	"chitChat/data"
	"net/http"
)

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)			
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

//index函数负责生成html文件并将其写入到ResponseWriter结构体中
func index(w http.ResponseWriter, r *http.Request) {					//请求参数可以通过访问request结构体来得到
	thread, err := data.Threads(); if err == nil {
		_, err := session(w, r)

		if err != nil {
			generateHTML(w, thread, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, thread, "layout", "private.navbar", "index")
		}
	}
}