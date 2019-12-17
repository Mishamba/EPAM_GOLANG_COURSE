package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type errorToSend struct {
	title string
	body  string
}

func correctMethod(r *http.Request) bool {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		return false
	}

	return true
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:  "token",
		Value: r.FormValue("name") + ":" + r.FormValue("address"),
	}
	http.SetCookie(w, &c)
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	if !correctMethod(r) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		t, _ := template.ParseFiles("error.html")
		t.Execute(w, errorToSend{
			http.StatusText(http.StatusMethodNotAllowed),
			"can't answer on such a request...",
		})
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		t, _ := template.ParseFiles("error.html")
		t.Execute(w, errorToSend{
			http.StatusText(http.StatusNotFound),
			"can't parse file...",
		})
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, err)
			return
		}

		setCookie(w, r)
	}

	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handleConn)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
