package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listen = flag.String("http", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	fs := http.FileServer(http.Dir(*dir))
	log.Fatal(http.ListenAndServe(*listen,

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.Path)
			switch r.Method {
			default:
				http.Error(w, "bad request", http.StatusBadRequest)
			case "GET":
				fs.ServeHTTP(w, r)
			case "PUT":
				handleUpload(w, r)
			}

		})))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Println("upload:", err)
		return
	}
	fmt.Println("upload:", string(data))
}
