package service

import (
	"log"
	"net/http"
	
)

func StartService(port string){
	r := NewRouter()
	http.Handle("/",r)
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting http listenr at port " + port)
		log.Println("Error: " + err.Error())
	}


}