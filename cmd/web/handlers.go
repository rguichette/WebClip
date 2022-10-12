package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // review
		return
	}
	//use the template.ParseFiles() function to read template file into a
	//templateset. If there's an error, we log the detailed error message and use
	//the http.Error() function to send a generic 500 internal server Error
	// response to the user.
	files := []string{"./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl"}

	//Use the template.ParseFiles() function to read the files and store the
	//templates in a template set. Notice that we can pass the slice of file paths
	//as a variadic parameter?
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Sever Error", http.StatusInternalServerError)
		return
	}
	//we the use the Execute() method on the template set to write the template
	//content as the response body. The last parameter to Execute() respresents any
	//dynamic data that we want to pass in which for not we'll leave as nil

	err = ts.Execute(w, nil) //review

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("hwllo from Webclip"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Display a specific snipper")
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//same as passing num 405 manually, but instead using golang constant
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
