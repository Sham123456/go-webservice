package main

import (
	"net/http"

	"github.com/Sham123456/go-webservice/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
