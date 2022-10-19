package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rguichette/webclip/pkg/models/mysql"
)

//definne an application struct to hold the application-wide dependencies for the
//web application. For now we'll only include fiels for two custom loggers, but
//we'll add more to it as the build progresses.

type application struct {
	errLog        *log.Logger
	infoLog       *log.Logger
	clips         *mysql.ClipModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	//define a new command-line flag for the mysql dsn string
	dsn := flag.String("dsn", "web:pass/webclip?parseTime=True", "MySQL data source name")

	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	defer db.Close()
	//init ne template cache
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errLog.Fatal(err)
	}

	//Init new instance of Application containing the dependencies.
	//add it to the application
	app := &application{
		errLog:        errLog,
		infoLog:       infoLog,
		clips:         &mysql.ClipModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()

	errLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
