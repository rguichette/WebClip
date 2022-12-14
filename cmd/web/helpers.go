package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//The serverError helper writes an error message and stack trace to the errorLog,
//then sends a generic 500 Internal Server Error respose to the user.

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//The clientError helper sends a specific status code and corresponding description
//to the user. We'll use this later in the book to send responses like 400 "bad" Request when there's a problem with the request that the user sent.

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//For consistency, we'll also implement a notound helper. This is simply a convenience wrapper around clienError which sends a 404 Not found responce to the user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Retrieve the appropraite temple set from the cache based on the page name. If no entry exists in the cache with the proviced name, call the serverError helper method.
	ts, ok := app.templateCache[name]

	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}
	//exec the template set, passing in any dynamic data.
	err := ts.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}

}
