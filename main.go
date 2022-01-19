package main

import (
	"log"
	"net/http"

	"github.com/Shubhamag12/HMS/routes"
)

func main()  {
	router := routes.Router()
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", router))
}