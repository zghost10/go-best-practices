package user

import "github.com/zghost10/go-best-practices/internal/domain/user"

type ListUsersOutput struct {
	Users []ListUser
}

type ListUser struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
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

	var listUsersOutput ListUsersOutput
	for _, user := range users {
		listUsersOutput.Users = append(listUsersOutput.Users, ListUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}

	return &listUsersOutput, nil
}
