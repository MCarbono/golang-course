package main

import (
	"api/configs"
	"api/internal/entity"
	"api/internal/infra/database"
	"api/internal/infra/webserver/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/jwtauth"

	_ "api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcelo Carbono

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := database.NewConnection("test.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productRepository := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productRepository)

	userRepository := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userRepository, configs.JWTExpiresIn)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(JSONHeader)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	log.Fatal(http.ListenAndServe(":8000", r))
}

func JSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Products CURL
//curl -X POST 'localhost:8000/products' -H 'Content-Type: application/json' -d '{"name": "Product 4", "price": 40.0}'
//curl localhost:8000/products/84ab5774-d534-4a46-a7c9-763c9a52c725
//curl -X PUT 'localhost:8000/products/342862c2-bc98-4d64-8a0f-a4bf6997e642' -H 'Content-Type: application/json' -d '{"name": "Product Updated", "price": 5.0}'
//curl -X DELETE 'localhost:8000/products/342862c2-bc98-4d64-8a0f-a4bf6997e642'
//curl localhost:8000/products -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA3MDk2MjYsInN1YiI6ImYyOTVmMDZjLThjOTQtNGUyZC04ODc3LWY0ZjY5ZDMzODMyYiJ9.w3kODqSwtzrgJbGA8ikCZe9MhVhuRNAS8Sfeqg2zj1U'

//User CURL
//curl -X POST 'localhost:8000/users' -H 'Content-Type: application/json' -d '{"name": "User 1", "email": "user1@gmail.com", "password": "123456"}'
//curl -X POST 'localhost:8000/users/generate_token' -H 'Content-Type: application/json' -d '{"email": "user1@gmail.com", "password": "123456"}'

//swag
//iniciar folder de docs/atualizar docs
//swag init -g cmd/server/main.go
