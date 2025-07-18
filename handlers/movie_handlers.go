package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fepu08/vanilla-js-go/data"
	"github.com/fepu08/vanilla-js-go/logger"
	"github.com/fepu08/vanilla-js-go/models"
)

type MovieHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

func NewMovieHandler(movieRepo data.MovieStorage, logInstance *logger.Logger) *MovieHandler {
	return &MovieHandler{storage: movieRepo, logger: logInstance}
}

func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (movieHandler *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		// TODO: log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (movieHandler *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := movieHandler.storage.GetTopMovies()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		movieHandler.logger.Error("Get Top Movies Error", err)
		return
	}
	movieHandler.writeJSONResponse(w, movies)
}

func (movieHandler *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          3,
			TMDB_ID:     181,
			Title:       "Die Hard",
			ReleaseYear: 1995,
			Genres:      []models.Genre{{ID: 1, Name: "Action"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max", LastName: "Wick"}},
		},
		{
			ID:          4,
			TMDB_ID:     181,
			Title:       "Mad Max",
			ReleaseYear: 2004,
			Genres:      []models.Genre{{ID: 1, Name: "Action"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Jon", LastName: "Doe"}},
		},
	}

	movieHandler.writeJSONResponse(w, movies)
}
