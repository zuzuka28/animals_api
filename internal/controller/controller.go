package controller

import (
	"animals_api/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Controller struct {
	router      *gin.Engine
	service     service.Service
	userService service.UserService
}

const accountIdKey = "accountId"
const animalIdKey = "animalId"
const typeIdKey = "typeId"
const pointIdKey = "pointId"

func NewController(engine *gin.Engine, service service.Service, userService service.UserService) *Controller {
	r := engine
	if r == nil {
		r = gin.Default()
	}

	c := Controller{
		router:      r,
		service:     service,
		userService: userService,
	}

	log.Printf("hello")
	c.router.Use(gin.ErrorLogger())

	c.router.GET(fmt.Sprintf("/accounts/:%s", accountIdKey), c.Account)
	c.router.PUT(fmt.Sprintf("/accounts/:%s", accountIdKey), c.UpdateAccount)
	c.router.DELETE(fmt.Sprintf("/accounts/:%s", accountIdKey), c.DeleteAccount)
	c.router.GET("/accounts/search", c.SearchAccount)

	c.router.GET(fmt.Sprintf("/animals/:%s", animalIdKey), c.Animal)
	c.router.GET("/animals/search", c.SearchAnimal)
	c.router.GET(fmt.Sprintf("/animals/types/:%s", typeIdKey), c.AnimalType)
	c.router.GET(fmt.Sprintf("/animals/:%s/locations", animalIdKey), c.AnimalLocations)

	c.router.GET(fmt.Sprintf("/locations/:%s", pointIdKey), c.Location)
	c.router.POST("/locations", c.NewLocation)
	c.router.PUT(fmt.Sprintf("/locations/:%s", pointIdKey), c.UpdateLocation)
	c.router.DELETE(fmt.Sprintf("/locations/:%s", pointIdKey), c.DeleteLocation)

	return &c
}

func (c *Controller) Run(addr ...string) error {
	return c.router.Run(addr...)
}
