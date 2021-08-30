package main

import (
	"database/sql"
	"github.com/Phamvandat207/Basic-store/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/basic-store")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/posts", handler.GetProducts).Methods("GET")
	router.HandleFunc("/posts", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/posts/{id}", handler.GetProduct).Methods("GET")
	router.HandleFunc("/posts/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/posts/{id}", handler.DeleteProduct).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
