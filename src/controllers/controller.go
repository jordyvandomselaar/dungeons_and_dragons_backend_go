package controllers

import (
	"net/http"
	"fmt"
)

type Controller struct{}

func(c *Controller) JsonResponse(w http.ResponseWriter, data []byte) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	fmt.Fprint(w, string(data))
}
