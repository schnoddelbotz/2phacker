package webserver

import (
	"log"
	"net/http"
)

func Serve(docRoot string, port string) {
	fs := http.FileServer(http.Dir(docRoot))
	http.Handle("/", fs)

	log.Printf("Listening on %s...", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
