package controller

import (
	"animals_api/internal/entity"
	"animals_api/internal/utils/typed_errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	firstNameKey = "firstName"
	lastNameKey  = "lastName"
	emailKey     = "email"
	fromKey      = "from"
	sizeKey      = "size"
)

func (c *Controller) Account(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.User{}

	if params.Id, err = validateAccountId(ctx.Param(accountIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.userService.GetUser(params.Id)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) SearchAccount(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.AccountRequestCtx{}

	params.FirstName = ctx.Query(firstNameKey)
	params.LastName = ctx.Query(lastNameKey)
	params.Email = ctx.Query(emailKey)

	if params.From, err = validateFrom(ctx.Query(fromKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if params.Size, err = validateSize(ctx.Query(sizeKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.userService.SearchUser(params)
	if err != nil {

		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) UpdateAccount(ctx *gin.Context) {
	var (
		err     error
		account entity.User
	)

	if err = ctx.BindJSON(account); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if account.Id, err = validateAccountId(ctx.Param(accountIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if account, err = validateAccount(account); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.userService.UpdateUser(account)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.Forbidden:
			ctx.Status(http.StatusForbidden)
			return
		case typed_errors.Conflict:
			ctx.Status(http.StatusConflict)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) DeleteAccount(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.AccountRequestCtx{}

	if params.Id, err = validateAccountId(ctx.Param(accountIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err = c.userService.DeleteUser(params.Id); err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.Forbidden:
			ctx.Status(http.StatusForbidden)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) Animal(ctx *gin.Context) {

}

func (c *Controller) SearchAnimal(ctx *gin.Context) {

}

func (c *Controller) AnimalType(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.AnimalType{}

	if params.Id, err = validateAnimalTypeId(ctx.Param(typeIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.service.GetType(params.Id)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) NewAnimalType(ctx *gin.Context) {
	var (
		err        error
		animalType entity.AnimalType
	)

	if err = ctx.BindJSON(animalType); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	resp, err := c.service.NewType(animalType)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.Conflict:
			ctx.Status(http.StatusConflict)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) UpdateAnimalType(ctx *gin.Context) {
	var (
		err        error
		animalType entity.AnimalType
	)

	if err = ctx.BindJSON(animalType); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if animalType.Id, err = validateAnimalTypeId(ctx.Param(typeIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if animalType, err = validateAnimalType(animalType); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.service.UpdateType(animalType)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) DeleteAnimalType(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.AnimalType{}

	if params.Id, err = validateAnimalTypeId(ctx.Param(typeIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err = c.service.DeleteAnimalType(params.Id); err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) Location(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.Location{}

	if params.Id, err = validatePointId(ctx.Param(pointIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.service.GetLocation(params.Id)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) NewLocation(ctx *gin.Context) {
	var (
		err      error
		location entity.Location
	)

	if err = ctx.BindJSON(location); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if err = validatePoint(location); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.service.NewLocation(location)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.Conflict:
			ctx.Status(http.StatusConflict)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) UpdateLocation(ctx *gin.Context) {
	var (
		err      error
		location entity.Location
	)

	if err = ctx.BindJSON(location); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if location.Id, err = validatePointId(ctx.Param(pointIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err = validatePoint(location); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	resp, err := c.service.UpdateLocation(location)
	if err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		case typed_errors.Conflict:
			ctx.Status(http.StatusConflict)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) DeleteLocation(ctx *gin.Context) {
	var (
		err error
	)

	params := &entity.Location{}

	if params.Id, err = validatePointId(ctx.Param(pointIdKey)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err = c.service.DeleteLocation(params.Id); err != nil {
		switch typed_errors.GetType(err) {
		case typed_errors.NotFound:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) AnimalLocations(ctx *gin.Context) {

}
