package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fepu08/vanilla-js-go/data"
	"github.com/fepu08/vanilla-js-go/handlers"
	"github.com/fepu08/vanilla-js-go/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logger, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("failed to initialize logger %v", err)
		panic(err)
	}

	return logger
}

func main() {
	logInstance := initializeLogger()
	defer logInstance.Close()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL not set in environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Data Repository for Movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movie repository")
	}

	// Initialize Account Repository for Users
	accountRepo, err := data.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize account repository")
	}

	// Handlers
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	accountHandler := handlers.NewAccountHandler(accountRepo, logInstance)

	// Endpoints
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register/", accountHandler.Register)
	http.HandleFunc("/api/account/authenticate/", accountHandler.Authenticate)

	catchAllClientRoutesHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}

	http.HandleFunc("/movies", catchAllClientRoutesHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failer %v", err)
	}
}
