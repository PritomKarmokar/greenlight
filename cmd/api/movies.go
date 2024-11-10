package main

import (
	"fmt"
	"greenlight.net/internal/data"
	"net/http"
	"time"
)

// using `Decode`
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	//err := json.NewDecoder(r.Body).Decode(&input)
	err := app.readJSON(w, r, &input)
	if err != nil {
		//app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

// using `Unmarshal`
//func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
//	var input struct {
//		Title   string   `json:"title"`
//		Year    int32    `json:"year"`
//		Runtime int32    `json:"runtime"`
//		Genres  []string `json:"genres"`
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		app.serverErrorResponse(w, r, err)
//		return
//	}
//
//	err = json.Unmarshal(body, &input)
//	if err != nil {
//		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	fmt.Fprintf(w, "%+v\n", input)
//}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		//http.NotFound(w, r)
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}
}
