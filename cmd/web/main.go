package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//definne an application struct to hold the application-wide dependencies for the
//web application. For now we'll only include fiels for two custom loggers, but
//we'll add more to it as the build progresses.

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {

	//Define a new command-line flag with the name "addr", a default value of ":4000"
	//and some short help text explaining what the flag controls. The value of the
	//flag will be stored in the add variable at runtime.

	addr := flag.String("addr", ":4000", "HTTP network address")
	//Importantly, we use the flag.Parse() function to parse the command-line flag.
	//This reads in the command-line flag value and assigns it to the addr
	//variable. Need to call this BEFORE using the addr variable
	// otherwise it will always contain the default of ":4000"
	//if any errors are encountered during parsing the application will be terminated.

	flag.Parse()

	//use log.New() to create a logger for writing information messages. This takes
	//three parameters: the destination to write the logs, a string
	//prefix for messages (Info followed by tab) and flags to indicate what
	//additional information to include (local data and time) Note that the flags
	//are joind using the bitwise OR operator |
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//create a logger for writing error messages in the same way, but use stderr as
	//the destination and use the log.Lshortfile flag to include the relevant
	//file name and line number.

	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Init new instance of Application containing the dependencies.
	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	//Create a file server whiich serves files out of the "./ui/static" directory.
	//Note that the path given to the http.Dir function is relative to the project
	//direcory root.

	fileServer := http.FileServer(http.Dir("./ui/static/")) //review
	//Use the mux.Handle() function to register the file server as the handler fo
	//all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //review
	// mux.Handle("/static/", fileServer) //review

	// The value returned from the flag.String() function is a pointer to the flag
	//value not the value itself. (i.e: prefix it with the *symbol) before using it.
	//Note that we're using the log.PringF=f() function to interpolate the address with the log message.

	//Initialize a new Http.Server struct. We set the Addr and Handler fields so
	//that the server users the same network address and routes as before, and set
	//the ErrorLof field so that the server now uses the custom errorLog logger in
	//the event of any probles.

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	//Call the ListenAndSeve() method on out new http.Server struct.
	err := srv.ListenAndServe()
	// log.Printf("starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	errLog.Fatal(err)
}
