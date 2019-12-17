package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type answer struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestUri string      `json:"request_uri"`
	Header     http.Header `json:"headers"`
}

func HandleAnsw(w http.ResponseWriter, r *http.Request) {
	ans := answer{
		r.Host,
		r.UserAgent(),
		r.RequestURI,
		r.Header,
	}

	buf, err := json.Marshal(ans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Set("Content-Type", "applicatioin/json; charset=UTF-8")
	w.Write(buf)
}

func main() {
	http.HandleFunc("/", HandleAnsw)
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}
