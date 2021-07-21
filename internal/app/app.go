package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/speps/go-hashids/v2"
)

// App with a router and db as dependencies
type App struct {
	router *mux.Router
	db     *sql.DB
	h      *hashids.HashID
}

func NewApp() *App {
	a := new(App)
	a.initRouters()
	a.initDB()

	return a
}

// initRouters init routes
func (a *App) initRouters() {
	a.router.HandleFunc("/", a.Shorten).Methods("POST")
	a.router.HandleFunc("/", a.Longer).Methods("POST")
}

// InitDB init routes
func (a *App) initDB() {
	a.db = InitDB("mysql")
}

func (a *App) initHash() {
	hd := hashids.NewData()
	hd.Salt = "salt"
	hd.MinLength = 6
	a.h, _ = hashids.NewWithData(hd)
}

// Run the app
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.router))
}
