//create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));
//docker-compose exec mysql bash
//mysql -u root -p goexpert
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(product)
		}
	}
}

// db.Create(&Product{
// 	Name:  "Notebook",
// 	Price: 1000.00,
// })

//create batch
// products := []Product{
// 	{Name: "Notebook", Price: 1000.00},
// 	{Name: "Mouse", Price: 50.00},
// 	{Name: "Keyboard", Price: 100.00},
// }
// db.Create(&products)

//select one
//var products Product
// db.First(&product, 1)
// fmt.Println(product)
// db.First(&product, "name = ?", "Mouse")
// fmt.Println(product)

//select all
// var products []Product
// db.Find(&products)
// for _, v := range products {
// 	fmt.Println(v)
// }

// var products []Product
// db.Limit(2).Offset(2).Find(&products)
// for _, v := range products {
// 	fmt.Println(v)
// }

//where
// var products []Product
// db.Where("price > ?", 90).Find(&products)
// for _, v := range products {
// 	fmt.Println(v)
// }

// var products []Product
// db.Where("name LIKE ?", "%k%").Find(&products)
// for _, v := range products {
// 	fmt.Println(v)
// }

// var p Product
// db.First(&p, 1)
// p.Name = "New mouse"
// db.Save(&p)

// var p2 Product
// db.First(&p2, 1)
// fmt.Println(p2)
// db.Delete(&p2)

// db.Create(&Product{
// 	Name:  "Notebook",
// 	Price: 1000.00,
// })

// var p Product
// db.First(&p, 1)
// p.Name = "New Mouse 2"
// db.Save(&p)
