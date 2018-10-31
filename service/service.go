package service

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

// StartService strat service
func StartService(port string) {
	router := NewRouter()
	http.Handle("/", router)
	log.Println("Starting HTTP service at " + port)
	// add cors in develop enviroment
	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	// err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting http listenr at port " + port)
		log.Println("Error: " + err.Error())
	}

}
