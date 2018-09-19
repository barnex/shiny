package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

var (
	flagPort   = flag.String("http", ":8080", "listen address")
	flagDir    = flag.String("dir", ".", "directory to serve")
	flagLevels = flag.String("levels", "levels", "directory to save levels")
)

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*flagDir))
	http.Handle("/", withLog(fs))
	http.Handle("/put/", withLog(http.StripPrefix("/put/", http.HandlerFunc(handlePut))))

	log.Printf("listening on %v", *flagPort)
	log.Fatal(http.ListenAndServe(*flagPort, nil))
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	saveLevel(w, r)
	buildLevels()
}

func saveLevel(w http.ResponseWriter, r *http.Request) {
	levelNum := path.Dir(r.URL.Path)
	f, err := os.Create(path.Join(*flagLevels, levelNum+".level"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(path.Base(r.URL.Path))
}

func buildLevels() {
	out, err := exec.Command("./build_levels.sh").CombinedOutput()
	os.Stderr.Write(out)
	if err != nil {
		fmt.Println(err)
	}
}

func withLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
