package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/hirany/web_app/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}
func main() {
	var addr = flag.String("addr", ":8000", "アプリケーションのアドレス")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	//http.Handle("/chat!", MustAuth(&templateHandler{filename: "ht.html"}))
	http.Handle("/login", &templateHandler{filename: "loginpage1.html"})
	//http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバを開始 ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
