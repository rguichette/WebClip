package main

import (
	"log"
	"net/http"
)

//define a home handler function which write a byte slice containing
//"Hello from webclip" as the response body

func home(w http.ResponseWriter, r *http.Request) {
	//check if the current request url path exactly matches "/".
	//if it doesn't, use the http.NotFound() function to send 404 response to the client.
	//Importantly, we then return from the handler. If we don't return, the handler
	//would keep executing and also write the "Hello from webclip" message
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from webclip"))

}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	//Use r.Method to check whether the request is using POST or not. Note that
	//http.MethodPost is a constant equal to the string "POST"
	if r.Method != http.MethodPost {
		//if it's not, use the w.Writehead() methos to send a 405 status
		//code and the w.Write() method to write a "Methos Not allowed"
		// response body. We then return from the function so that the
		//subsequesnt code is not executed
		w.WriteHeader(405)
		w.Write([]byte("Method Not allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	//use the htttp.NerServeMux() funnction to init a new servemux, then
	//regiser the home function as a handler for the "/" url patter.

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//register the new handlers
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/createSnippet", createSnippet)

	//Use the http.ListenAndServe() functioin to start a new web server. We pass in
	//two parameters: the TCP network address to listen on( iin this: 4000)
	//and the servemux we just create. If http.ListernAndServe returns an error
	// we use the log.Fatalfunction to log the error message and exiitt. Note
	//that any error returned by the http.LostenAndServe() is always non-nil
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
