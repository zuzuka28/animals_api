package controller

import (
	"animals_api/internal/entity"
	"github.com/pkg/errors"
	"net/mail"
	"strconv"
)

func validateAccountId(value string) (int, error) {
	accountId, err := strconv.Atoi(value)
	if err != nil {
		return -1, errors.New("accountId is null")
	}

	if accountId <= 0 {
		return -1, errors.New("accountId must be > 0")
	}

	return accountId, nil
}

func validatePointId(value string) (int64, error) {
	accountId, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		return -1, errors.New("pointId is null")
	}

	if accountId < 0 {
		return -1, errors.New("pointId must be >= 0")
	}

	return accountId, nil
}

func validateSize(value string) (int, error) {
	size, err := strconv.Atoi(value)
	if err != nil {
		return -1, errors.New("size is null")
	}

	if size < 0 {
		return -1, errors.New("size must be >= 0")
	}

	return size, nil
}

func validateFrom(value string) (int, error) {
	from, err := strconv.Atoi(value)
	if err != nil {
		return -1, errors.New("from is null")
	}

	if from <= 0 {
		return -1, errors.New("from must be > 0")
	}

	return from, nil
}

func validateAccount(account entity.User) (entity.User, error) {
	// todo: check spaces

	valid := entity.User{Id: account.Id}

	if account.FirstName == "" {
		return valid, errors.New("bad first name")
	}
	if account.LastName == "" {
		return valid, errors.New("bad last name")
	}
	if _, err := mail.ParseAddress(account.Email); err != nil {
		return valid, errors.New("bad email")
	}
	if account.Password == "" {
		return valid, errors.New("bad password")
	}

	return valid, nil
}

func validatePoint(value entity.Location) error {
	// todo: check null
	if value.Latitude < -90 || value.Latitude > 90 {
		return errors.New("lat is null")
	}

	if value.Latitude < -180 || value.Latitude > 180 {
		return errors.New("lat is null")
	}

	return nil
}
