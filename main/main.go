package main

import (
	"net/http"
	"time"

	"lheinrich.de/extgo/shorts"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/handler"
)

// main function
func main() {
	setup()
}

// setup application
func setup() {
	// load config
	conf.Config = conf.ReadConfig("config.json")

	// initialize logging to file
	shorts.InitLoggingFile(time.Now().Format(conf.Config["app"]["logFile"]))

	// setups
	setupDB()
	setupWebserver()
}

// setup database
func setupDB() {
	// define config values
	host := conf.Config["postgresql"]["host"]
	port := conf.Config["postgresql"]["port"]
	database := conf.Config["postgresql"]["database"]
	username := conf.Config["postgresql"]["username"]
	password := conf.Config["postgresql"]["password"]
	ssl := conf.Config["postgresql"]["ssl"]

	// connect to postgresql database
	conf.DB = shorts.ConnectPostgreSQL(host, port, database, username, password, ssl)
}

// setup webserver
func setupWebserver() {
	// setup handlers
	setupHandlers()

	// define config values
	address := conf.Config["webserver"]["address"]
	cert := conf.Config["webserver"]["certificateFile"]
	key := conf.Config["webserver"]["keyFile"]

	// start http server
	http.ListenAndServeTLS(address, cert, key, nil)
}

// setup webserver handlers
func setupHandlers() {
	// register handlers
	http.HandleFunc("/index.html", handler.Index)
}
