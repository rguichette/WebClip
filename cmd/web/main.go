package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Create a file server whiich serves files out of the "./ui/static" directory.
	//Note that the path given to the http.Dir function is relative to the project
	//direcory root.

	fileServer := http.FileServer(http.Dir("./ui/static/")) //review
	//Use the mux.Handle() function to register the file server as the handler fo
	//all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //review

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
