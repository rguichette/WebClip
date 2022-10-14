package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//Change the signature of the home handler so it is deined as a method against
// *application

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
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
		//Because the home handler function is now a method against application
		//it can access its fields, including the error logger. We'll write the log
		//message to this instead of the standard logger.
		app.errLog.Println(err.Error())
		// log.Println(err.Error())
		app.serverError(w, err)
		// http.Error(w, "Internal Sever Error", http.StatusInternalServerError)
		return
	}
	//we the use the Execute() method on the template set to write the template
	//content as the response body. The last parameter to Execute() respresents any
	//dynamic data that we want to pass in which for not we'll leave as nil

	err = ts.Execute(w, nil) //review

	if err != nil {
		//also update the code here to use the error logger from the application struct.
		app.errLog.Println(err.Error())
		// log.Println(err.Error())
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	// w.Write([]byte("hwllo from Webclip"))
}

//change the signature of the showSnippet handler so it is defined as a method
//against *application.

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprint(w, "Display a specific snipper")
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//same as passing num 405 manually, but instead using golang constant
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
