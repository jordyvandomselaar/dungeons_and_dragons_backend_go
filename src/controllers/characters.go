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
			Name: "Gentoo",
		},
		{
			Name: "Jim",
		},
	}

	response, _ := json.Marshal(characters)

	w.Header().Add("content-type", "application/json")
	fmt.Fprint(w, string(response))
}
