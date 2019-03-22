package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/tvducmt/Auto-grading/api"
)

var (
	postgresURI string
	db          *sqlx.DB
)

func createRouter(env api.Env) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", api.HandleAPI(env.Index).Reg)
	router.Post("/register", api.HandleAPI(env.Register).Reg)
	return router
}

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}
// 	postgresURI = fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		"localhost", 5432, "postgres", "123456", "pratice_golang")
// 	db, err = sqlx.Connect("postgres", postgresURI)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Database is connected.")
// }

func main() {
	env := api.Env{
		DB: db,
	}
	//createRouter
	fmt.Println("Server on port 3000...")
	//Create Server
	router := createRouter(env)
	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
