package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joshi4/uerrors"
)

func main() {
	http.HandleFunc("/feature", serveFeature)
	http.HandleFunc("/", serveIndex)
	http.ListenAndServe(":6060", nil)
}

func serveIndex(w http.ResponseWriter, req *http.Request) {
	uerr := uerrors.New("example: redirect failed", "Please visit http://localhost:6060/feature if you aren't redirected automatically")
	log.Println(uerr.Error())
	io.WriteString(w, uerr.UserError())
}

func serveFeature(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	uerr := uerrors.FromErrors(fmt.Errorf("example: error=true, endpoint=%s", path),
		fmt.Errorf("Please visit %s a little later", path))
	log.Println(uerr.Error())
	io.WriteString(w, uerr.UserError())
}
