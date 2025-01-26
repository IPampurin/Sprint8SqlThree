package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo1.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	product := "Облачное хранилище"
	price := 300
	// название продукта и цена передаются через параметры
	res, err := db.Exec("INSERT INTO products (product, price) VALUES (:product, :price)",
		sql.Named("product", product),
		sql.Named("price", price))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}
