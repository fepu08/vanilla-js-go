package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fepu08/vanilla-js-go/models"
)

type MovieHandler struct {
	// TODO
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
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
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max", LastName: "Wick"}},
		},
		{
			ID:          2,
			TMDB_ID:     181,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 1, Name: "Scifi"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Jon", LastName: "Doe"}},
		},
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
