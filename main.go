package main

import (
	"net/http"

	"github.com/ceciquin/WebServiceGo/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
