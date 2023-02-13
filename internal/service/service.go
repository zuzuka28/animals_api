package service

import (
	"animals_api/internal/entity"
)

type Service interface {
	GetAnimal(id int)
	GetLocation(id int64) (entity.Location, error)
	UpdateLocation(location entity.Location) (entity.Location, error)
	NewLocation(location entity.Location) (entity.Location, error)
	DeleteLocation(id int64) error
	GetType(id int64) (entity.AnimalType, error)
	NewType(animalType entity.AnimalType) (entity.AnimalType, error)
	UpdateType(animalType entity.AnimalType) (entity.AnimalType, error)
	DeleteAnimalType(id int64) error
}

type AnimalService struct {
}

func (s *AnimalService) GetAnimal(id int) {
}

func (s *AnimalService) GetLocation(id int64) (entity.Location, error) {
	return entity.Location{
		Id:        1,
		Latitude:  90,
		Longitude: -90,
	}, nil
}

func (s *AnimalService) NewLocation(location entity.Location) (entity.Location, error) {
	return entity.Location{
		Id:        1,
		Latitude:  90,
		Longitude: -90,
	}, nil
}

func (s *AnimalService) UpdateLocation(location entity.Location) (entity.Location, error) {
	return entity.Location{
		Id:        2,
		Latitude:  90,
		Longitude: -90,
	}, nil
}

func (s *AnimalService) DeleteLocation(id int64) error {
	// todo: read api declaration
	return nil
}

func (s *AnimalService) GetType(id int64) (entity.AnimalType, error) {
	// todo: read api declaration
	return entity.AnimalType{
		Id:   1,
		Type: "test type",
	}, nil
}

func (s *AnimalService) NewType(animalType entity.AnimalType) (entity.AnimalType, error) {
	// todo: read api declaration
	return entity.AnimalType{
		Id:   1,
		Type: "test type",
	}, nil
}

func (s *AnimalService) UpdateType(animalType entity.AnimalType) (entity.AnimalType, error) {
	// todo: read api declaration
	return entity.AnimalType{
		Id:   1,
		Type: "test type",
	}, nil
}

func (s *AnimalService) DeleteAnimalType(id int64) error {
	// todo: read api declaration
	return nil
}
