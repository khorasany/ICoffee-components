package helpers

import (
	"log"
	"net/http"
)

func GetVerbose(r *http.Request) {
	log.Println(r.Method, r.Host, r.RequestURI)
}
