package main

import (
	"animals_api/internal/controller"
	"animals_api/internal/service"
	"log"
)

func main() {
	as := &service.AnimalService{}
	us := &service.AccountService{}
	c := controller.NewController(nil, as, us)

	err := c.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
