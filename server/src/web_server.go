package src

import (
	"fmt"
	"net/http"

	"github.com/glaukiol1/golssh/server/server/helpers"
)

var password = "PASSWORD"

var pass_hash, _ = helpers.HashPassword(password)

func auth(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	if helpers.CheckPasswordHash(req.URL.Query().Get("password"), pass_hash) {
		fmt.Fprintf(w, "success")
	} else {
		fmt.Fprintf(w, "error")
	}
}

func StartWebServer() {

	http.HandleFunc("/auth", auth)
	http.ListenAndServe(":22", nil)
}
