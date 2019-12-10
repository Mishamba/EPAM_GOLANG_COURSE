package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type answer struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestUri string      `json:"request_uri"`
	H          http.Header `json:"headers"`
}

func HandleAnsw(w http.ResponseWriter, r *http.Request) {
	ans := answer{
		r.Host,
		r.UserAgent(),
		r.RequestURI,
		r.Header,
	}

	w.Header().Set("Content-Type", "applicatioin/json; charset=UTF-8")
	buf, _ := json.Marshal(ans)
	w.Write(buf)
}

func main() {
	http.HandleFunc("/", HandleAnsw)
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}
