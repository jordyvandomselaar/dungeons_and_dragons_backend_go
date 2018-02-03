package routes

import (
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/controllers/authentication"
    "github.com/gorilla/mux"
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/controllers"
)

func Initialize() *mux.Router {
    r := mux.NewRouter()

    rc := authentication.Register{}
    cc := controllers.Characters{}

    r.HandleFunc("/register", rc.Index).Name("authentication.register").Methods("POST")
    r.HandleFunc("/characters", cc.Index).Name("characters.index").Methods("GET")

    return r
}
