package main

import (
	"context"
	"database/sql"
	"fmt"
	"sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{ID: uuid.New().String(), Name: "Backend", Description: sql.NullString{String: "Backend Description", Valid: true}})
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, v := range categories {
	// 	fmt.Println(v.ID, v.Name, v.Description)
	// }
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "dd96d30a-d656-468c-b0e3-5f393b7091ae",
		Name:        "Backend updated",
		Description: sql.NullString{String: "Description updated", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, v := range categories {
		fmt.Println(v.ID, v.Name, v.Description)
	}

	err = queries.DeleteCategory(ctx, "dd96d30a-d656-468c-b0e3-5f393b7091ae")
	if err != nil {
		panic(err)
	}
}
