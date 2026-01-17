package user

import (
	"fmt"

	"github.com/zghost10/go-best-practices/internal/domain/user"
)

func NewInMemoryUserRepo() user.IUserRepository {
	return &UserInMemoryRepo{
		users: make(map[string]user.User),
	}
}

type UserInMemoryRepo struct {
	users map[string]user.User
}

func (u *UserInMemoryRepo) Create(user user.User) error {
	u.users[user.ID] = user
	return nil
}

func (u *UserInMemoryRepo) Get(identifier string) (user.User, error) {
	usr, ok := u.users[identifier]
	if !ok {
		return user.User{}, fmt.Errorf("user not found")
	}
	return usr, nil
}

func (u *UserInMemoryRepo) GetAll() ([]user.User, error) {
	users := make([]user.User, len(u.users))
	i := 0
	for _, user := range u.users {
		users[i] = user
		i++
	}
	return users, nil
}

func (u *UserInMemoryRepo) Update(user user.User) error {
	u.users[user.ID] = user
	return nil
}

func (u *UserInMemoryRepo) Delete(identifier string) error {
	delete(u.users, identifier)
	return nil
}
