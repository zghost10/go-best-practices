package usecase

import (
	"github.com/google/uuid"
	"github.com/zghost10/go-best-practices/internal/domain/user"
)

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateUserOutput struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

type CreateUserUseCase struct {
	repo user.IUserRepository
}

func NewCreateUserUseCase(repo user.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (u *CreateUserUseCase) CreateUser(input CreateUserInput) (*CreateUserOutput, error) {
	user := user.User{
		ID:        uuid.New().String(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}
	err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
