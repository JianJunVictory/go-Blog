package main

import (
	"log"
	"./service"
)
func main(){
	log.Println("service start......")
	service.StartService("3000")

}