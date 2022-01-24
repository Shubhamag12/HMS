package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shubhamag12/HMS/conf"
	"github.com/Shubhamag12/HMS/routes"
)

func main()  {
	router := routes.Router()
	apiBase := fmt.Sprintf("%s:%s", conf.EnvMap["HOST"], conf.EnvMap["PORT"])
	log.Fatal(http.ListenAndServe(apiBase, router))
}