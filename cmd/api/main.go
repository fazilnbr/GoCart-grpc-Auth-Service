package main

import (
	"fmt"
	"log"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/config"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	fmt.Println("config : ", config)
}
