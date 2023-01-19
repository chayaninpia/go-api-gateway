package main

import (
	"apigw/src/configs"
	"apigw/src/routers"
	"log"
)

func main() {

	log.Println(`------- Start API Gateway -------`)
	configs.Init()
	routers.Init()
	log.Println(`------- END API Gateway -------`)
}
