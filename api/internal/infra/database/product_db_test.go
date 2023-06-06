package database

import (
	"api/internal/entity"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	db, err := NewConnection("file::memory:")
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("John", 10)
	assert.NoError(t, err)
	produtDB := NewProduct(db)
	err = produtDB.Create(product)
	assert.Nil(t, err)
	insertedProduct, err := produtDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, insertedProduct.Name, product.Name)
	assert.Equal(t, insertedProduct.Price, product.Price)
	assert.NotNil(t, insertedProduct.CreatedAt, product.CreatedAt)
}

func TestFindAllProducts(t *testing.T) {
	db, err := NewConnection("file::memory:")
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := NewConnection("file::memory:")
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	product.Name = "Product 2"
	product.Price = 15
	productDB := NewProduct(db)
	err = productDB.Update(product)
	assert.NoError(t, err)
	updatedProduct, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, updatedProduct.Name)
	assert.Equal(t, product.Price, updatedProduct.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := NewConnection("file::memory:")
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
}
