package src

import (
	"io/ioutil"
	"net/http"
)

func Authenticate(password string, ip string) {
	resp, err := http.Get("http://" + ip + ":22/auth?password=" + password)

	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
}
