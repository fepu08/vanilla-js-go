package main

import (
	"log"
	"net/http"

	"github.com/fepu08/vanilla-js-go/handlers"
	"github.com/fepu08/vanilla-js-go/logger"
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

	movieHandler := handlers.NewMovieHandler()

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failer %v", err)
	}
}
