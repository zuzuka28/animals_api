package service

import "animals_api/internal/entity"

type UserService interface {
	GetUser(id int) (entity.User, error)
	SearchUser(ctx *entity.AccountRequestCtx) ([]entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) error
}

type AccountService struct {
}

func (s *AccountService) GetUser(id int) (entity.User, error) {
	return entity.User{
		Id:        1,
		FirstName: "test",
		LastName:  "tester",
		Email:     "test@test.ru",
	}, nil
}

func (s *AccountService) SearchUser(ctx *entity.AccountRequestCtx) ([]entity.User, error) {
	users := make([]entity.User, 0, 10)
	for i := 0; i < ctx.Size; i++ {
		users = append(users, entity.User{
			Id:        i,
			FirstName: "test",
			LastName:  "tester",
			Email:     "test@test.ru",
		})
	}
	return users, nil
}

func (s *AccountService) UpdateUser(user entity.User) (entity.User, error) {
	// todo: read api declaration
	return entity.User{
		Id:        1,
		FirstName: "test_update",
		LastName:  "tester",
		Email:     "test@test.ru",
	}, nil
}

func (s *AccountService) DeleteUser(id int) error {
	// todo: read api declaration
	return nil
}
