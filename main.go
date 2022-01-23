package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()					//多路复用器
	files := http.FileServer(http.Dir("/public"))				//创建了一个能够为指定目录中的静态文件服务的处理器，

	//静态文件应该在<root>/file下面，而请求的url应该是/static/file
	mux.Handle("/static/", http.StripPrefix("/static/", files))		//收到前缀为static的请求，移除掉/static/字符串，在Public目录中查找被请求的文件

	//index
	mux.HandleFunc("/", index)			//将对根目录的请求重定向到index处理器
	//error
	mux.HandleFunc("/err", err)

	//auth
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", logout)
	mux.HandleFunc("/signup_account", signupAccount)

	//thread
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)
	mux.HandleFunc("/thread/create", createThread)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
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