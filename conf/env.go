package conf

import (
	"log"

	"github.com/joho/godotenv"
)

var EnvMap map[string]string

func init()  {
	var err error
	EnvMap, err = godotenv.Read(".dev.env")
	if err != nil {
		log.Fatal(err)
	}
}