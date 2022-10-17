package main

import (
	"database/sql"
	"flag"
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
	errLog  *log.Logger
	infoLog *log.Logger
	clips   *mysql.ClipModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	//define a new command-line flag for the mysql dsn string
	dsn := flag.String("dsn", "web:password/webclip?parseTime=True", "MySQL data source name")

	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	defer db.Close()
	//init a mysql clipModel instance and add it to the application dependencies

	//Init new instance of Application containing the dependencies.
	app := &application{
		errLog:  errLog,
		infoLog: infoLog,
		clips:   &mysql.ClipModel{DB: db},
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
