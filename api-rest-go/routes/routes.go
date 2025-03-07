package routes

import (
	"api-rest-go/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade)
	log.Fatal(http.ListenAndServe(":8080", r))
}
