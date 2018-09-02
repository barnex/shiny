package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	fs := http.FileServer(http.Dir(*dir))
	log.Fatal(http.ListenAndServe(*listen,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r)
			fs.ServeHTTP(w, r)
		})))
}
