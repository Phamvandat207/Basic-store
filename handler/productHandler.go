package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Phamvandat207/Basic-store/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var db *sql.DB

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []model.Product
	result, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var product model.Product
		err := result.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO product (name, price) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	price := keyVal["price"]
	_, err = stmt.Exec(name, price)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New product was created")
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name, price FROM product WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var product model.Product
	for result.Next() {
		err := result.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(product)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE product SET name = ? , price = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["name"]
	newPrice := keyVal["price"]
	_, err = stmt.Exec(newName, newPrice, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Product with ID = %s was updated", params["id"])
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Product with ID = %s was deleted", params["id"])
}