package main

import (
	"chitChat/data"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address 			string
	ReadTimeout			int64
	WriteTimeout		int64
	Static				string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error){
	cookie, err := r.Cookie("cookie")		//如果不存在，cookie为空
	if err == nil {									//如果存在，检查Uuid是否存在，似乎并没有验证时间
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid Session")
		}
	}
	return
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
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

func info(args ...interface{}) {
	logger.SetPrefix("INFO")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func version() string {
	return "0.1"
}