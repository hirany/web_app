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
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
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
	gomniauth.SetSecurityKey("セキュリティーキー")
	gomniauth.WithProviders(
		google.New("", "", "http://localhost:8000/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/room", r)
	go r.run()

	http.Handle("/chat", MustAuth(&templateHandler{filename: "ht.html"}))
	http.Handle("/login", &templateHandler{filename: "loginpage1.html"})
	http.HandleFunc("/auth/", loginHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Webサーバを開始 ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
