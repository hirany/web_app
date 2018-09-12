package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
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
	var addr = flag.String("addr", ":8000", "アプリケーション")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "ht.html"})
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
