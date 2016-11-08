package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	// "path/filepath"

	_ "gopkg.in/mgo.v2"
)

// var Session *(mgo.Session);
func init() {
	// var err error
	// Session,err=mgo.Dial("localhost")
	// if err!=nil{
	// 	panic(err)
	// }
}

// 创建一个具有http.Handler接口的结构
// type messageHandler struct {
// 	message string
// }
//
// func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, m.message)
// }

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func NotExistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotExist 404")
}

func main() {
	// 解析输入的参数
	port := flag.String("p", "8000", "listening port")
	directory := flag.String("d", "public", "static file directory")
	flag.Parse()

	// 创建复用器
	mux := http.NewServeMux()
	// mh1 := &messageHandler{"Hello"}
	// mh2 := &messageHandler{"World"}

	mux.Handle("/hello", messageHandler("hello"))
	mux.Handle("/world", messageHandler("world"))

	// 创建不存在的404页面处理，这里调用HandlerFunc来处理，http.HandlerFunc返回一个符合Handler接口的对象
	mux.Handle("/404", http.HandlerFunc(NotExistHandler))
	mux.HandleFunc("/notexist", NotExistHandler)
	// 创建静态文件处理
	log.Printf("The static file dir is setting to %v", http.Dir(*directory))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*directory))))

	log.Println("Start the web server")

	server := &http.Server{
		Addr:           "localhost:" + *port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
