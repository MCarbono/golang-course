//create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));
//docker-compose exec mysql bash
//mysql -u root -p goexpert
package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// product := NewProduct("notebook", 1899.99)
	// if err = insertProduct(db, product); err != nil {
	// 	panic(err)
	// }
	// product.Price = 100.0
	// if err = updateProduct(db, product); err != nil {
	// 	panic(err)
	// }
	// product, err = selectOneProduct(context.Background(), db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// products, err := selectAllProducts(db)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(products)
	if err = deleteProduct(db, "ed3c7c80-7018-4004-a680-433c8c888eab"); err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id= ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	return err
}

func selectOneProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	return &p, err
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, ID string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ID)
	return err
}
