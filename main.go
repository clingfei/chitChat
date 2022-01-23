package main

import (
	"net/http"
	"time"
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
		Addr:  					config.Address,
		Handler:	 			mux,
		ReadTimeout: 			time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: 			time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 		1 << 20,
	}
	server.ListenAndServe()
}