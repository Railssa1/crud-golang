package controller

import (
	"html/template"
	"log"
	"loja/model"
	"net/http"
	"strconv"
)

var t = template.Must(template.ParseGlob("templates/html/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.SelectAllProducts()
	t.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		productName := r.FormValue("productName")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro ao converter o preço")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro ao converter o preço")
		}

		model.InsertProduct(productName, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	model.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := model.GetProductById(idProduct)
	t.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idProduct := r.FormValue("id")
		productName := r.FormValue("productName")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro ao converter preço")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro ao converter quantidade")
		}

		convertedId, err := strconv.Atoi(idProduct)
		if err != nil {
			log.Println("Erro ao converter id")
		}

		model.UpdateProduct(productName, description, convertedPrice, convertedQuantity, convertedId)
	}

	http.Redirect(w, r, "/", 301)
}