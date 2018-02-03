package main

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "net/http"
    "github.com/jordyvandomselaar/dungeons_and_dragons_backend_go/src/routes"
)

func main() {
    r := routes.Initialize()

    http.Handle("/", r)

    http.ListenAndServe("localhost:8000", nil)
}
