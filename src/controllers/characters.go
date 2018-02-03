package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/domain/entities"
	"net/http"
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

	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(response))
}
