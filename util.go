package main

import (
	"chitChat/data"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error){
	cookie, err := r.Cookie("_cookie")		//如果不存在，cookie为空
	if err == nil {									//如果存在，检查Uuid是否存在，似乎并没有验证时间
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid Session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface {}, fn ...string) {				// ...表示可变参数，接受0个或多个值作为参数
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}