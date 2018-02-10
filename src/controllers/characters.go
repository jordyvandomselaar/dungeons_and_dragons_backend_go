package controllers

import (
	"encoding/json"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"log"
)

type Characters struct {
	Controller
}

func (c *Characters) Index(w http.ResponseWriter, r *http.Request) {
	characters := []entities.Character{
		{
			Id:   1,
			Name: "Gentoo",
		},
		{
			Id:   2,
			Name: "Jim",
		},
		{
			Id:   3,
			Name: "Bob",
		},
	}

	response, _ := json.Marshal(characters)

	c.JsonResponse(w, response)
}

func (c *Characters) Show(w http.ResponseWriter, r *http.Request) {
	characters := []entities.Character{
		{
			Id:   1,
			Name: "Gentoo",
		},
		{
			Id:   2,
			Name: "Jim",
		},
		{
			Id:   3,
			Name: "Bob",
		},
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		log.Fatal(err)
	}

	var result entities.Character

	for _, character := range characters {
		if character.Id == id {
			result = character
		}
	}

	json, err := json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	c.JsonResponse(w, json)
}
