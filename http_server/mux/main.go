package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", myMux)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend - " + http.StatusText(404)))
	}

}
