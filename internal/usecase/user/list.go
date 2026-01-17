package usecase

import "github.com/zghost10/go-best-practices/internal/domain/user"

type ListUsersOutput struct {
	Users []ListUser `json:"users"`
}

type ListUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type ListUsersUseCase struct {
	repo user.IUserRepository
}

func NewListUsersUseCase(repo user.IUserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (u *ListUsersUseCase) Execute() (*ListUsersOutput, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	out := ListUsersOutput{
		Users: make([]ListUser, len(users)),
	}

	for i, user := range users {
		out.Users[i] = ListUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
	}

	return &out, nil
}
