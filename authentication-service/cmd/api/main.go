package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}


func main() {

	app := Config{}

	log.Panicln("Starting authentication service")
	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}	

