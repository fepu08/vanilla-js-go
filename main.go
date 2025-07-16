package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/fepu08/vanilla-js-go/handlers"
	"github.com/fepu08/vanilla-js-go/logger"
	"github.com/joho/godotenv"
)

func initializeLogger() *logger.Logger {
	logger, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("failed to initialize logger %v", err)
		panic(err)
	}

	defer logger.Close()
	return logger
}

func main() {
	logInstance := initializeLogger()

	// Reading Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	// Connect to DB
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal("Failed to connect to the DB: %v", err)
	}
	defer db.Close()

	movieHandler := handlers.NewMovieHandler()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failer %v", err)
	}
}
