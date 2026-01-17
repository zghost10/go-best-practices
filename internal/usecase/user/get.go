package usecase

import "github.com/zghost10/go-best-practices/internal/domain/user"

type GetUserInput struct {
	Identifier string
}

type GetUserOutput struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

type GetUserUseCase struct {
	repo user.IUserRepository
}

func NewGetUserUseCase(repo user.IUserRepository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (u *GetUserUseCase) GetUser(input GetUserInput) (*GetUserOutput, error) {
	user, err := u.repo.Get(input.Identifier)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
