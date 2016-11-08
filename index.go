package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}
func main() {
	port := flag.String("p", "8000", "listening port")
	directory := flag.String("d", "public", "static file directory")
	flag.Parse()

	mux := http.NewServeMux()
	// if path, err := filepath.Abs("public"); err == nil {
	// 	log.Printf("Abs path is %v", path)
	// }

	log.Printf("The static file dir is setting to %v", http.Dir(*directory))

	// mux.Handle("/", http.FileServer(http.Dir(*directory)))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*directory))))

	mh1 := &messageHandler{"Hello"}
	mh2 := &messageHandler{"World"}

	mux.Handle("/hello", mh1)
	mux.Handle("/world", mh2)

	log.Println("Start the web server")
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
