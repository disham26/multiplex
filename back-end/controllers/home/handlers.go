package home

import (
	"TechVeritasMultiplex/helpers"
	"TechVeritasMultiplex/model"
	"TechVeritasMultiplex/structs"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//GetAllShowDetails API is an XHR developed to get the said show details
func GetAllShowDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside shows")
		str := model.Multiplex
		Data, _ := json.Marshal(str)
		helpers.JSON(string(Data))(w, r)
		return
	}
}

//GetMovieList API to get the list of the movies
func GetMovieList() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Data, _ := json.Marshal(model.Movies)
		helpers.JSON(string(Data))(w, r)
		return
	}
}

//GetMultiplex function will be useful in calling the seats
func GetMultiplex() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Data, _ := json.Marshal(model.Multiplex)
		helpers.JSON(string(Data))(w, r)
		return
	}
}

//BookShowSeats API to book the selected seats
func BookShowSeats() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Vars := mux.Vars(r)
		showid := Vars["showid"]
		showIDint, _ := strconv.Atoi(showid)
		log.Println(showIDint)
		showIDint--
		seats := Vars["seats"]
		individualSeats := strings.Split(seats, ",")
		shows := model.UpdateVariable(showIDint, individualSeats)
		var response structs.BookResponse
		response.Success = true
		response.Shows = shows
		Data, _ := json.Marshal(response)
		helpers.JSON(string(Data))(w, r)
		return
	}
}
