package model

import "loja/db"

type Product struct {
	Id, Quantity             int
	ProductName, Description string
	Price                    float64
}

func SelectAllProducts() []Product {
	db := db.ConnectToBd()
	defer db.Close()

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var description, productName string
		var price float64

		err = selectAllProducts.Scan(&id, &productName, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.ProductName = productName
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	return products
}

func InsertProduct(productName, description string, price float64, quantity int) {
	db := db.ConnectToBd()
	defer db.Close()

	if insertProduct, err := db.Prepare("INSERT INTO products(productName, description, price, quantity) VALUES($1, $2, $3, $4)"); err != nil {
		panic(err.Error())
	} else {
		insertProduct.Exec(productName, description, price, quantity)
	}
}

func DeleteProduct(idProduct string) {
	db := db.ConnectToBd()
	defer db.Close()

	if deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1"); err != nil {
		panic(err.Error())
	} else {
		deleteProduct.Exec(idProduct)
	}
}

func GetProductById(idProduct string) Product {
	db := db.ConnectToBd()
	defer db.Close()

	product := Product{}

	if p, err := db.Query("SELECT * FROM products where id = $1", idProduct); err != nil {
		panic(err.Error())
	} else {
		for p.Next() {
			var id, quantity int
			var description, productName string
			var price float64

			if err := p.Scan(&id, &productName, &description, &price, &quantity); err != nil {
				panic(err.Error())
			}

			product.Id = id
			product.ProductName = productName
			product.Description = description
			product.Price = price
			product.Quantity = quantity
		}

		return product
	}
}

func UpdateProduct(productName, description string, price float64, quantity, id int) {
	db := db.ConnectToBd()
	defer db.Close()

	updateProduct, err := db.Prepare("UPDATE products SET productName = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(productName, description, price, quantity, id)
}
