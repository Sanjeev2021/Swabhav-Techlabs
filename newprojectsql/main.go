package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gormv1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//db.Create(&Product{Code: "L12342", Price: 10000})
	//db.Create(&Product{Code: "L12432", Price: 100034})

	// Read
	var product Product
	db.First(&product) // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println(product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 324242)

	// Delete - delete product
	db.Delete(&product)

}
