package router

import (
	"log"
    "go-postgres/middleware"
	"github.com/gorilla/mux"
)

type App struct{
    Addr string
}


func (app *App) Router() (R *mux.Router){

	R = mux.NewRouter()

    R.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET")
    R.HandleFunc("/api/stock" , middleware.CreateStock).Methods("Post","OPTIONS")
    R.HandleFunc("/api/stocks/{id}", middleware.GetById).Methods("Get")
    R.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT" , "PATCH")
    R.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE")


	log.Printf("server is started on %v", app.Addr)

	return
}