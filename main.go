package main

import (
	"log"

	"github.com/go-Blog/db"
	"github.com/go-Blog/logger"
	"github.com/go-Blog/service"
)

func init() {
	db.InitDbClient()
	logger.InitLog()

}
func main() {
	log.Println("service start......")
	service.StartService("3000")
}
