package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/logeshwhatnot/rssagg/internal/database"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in the env")
	}

	dbUrl := os.Getenv("DB_URL")
	if port == "" {
		log.Fatal("DB_URL not found in the env")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("couldn't connect to the database: %v", err)
	}

	apiCfg := apiConfig{
		db: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Listening on port: %v", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
