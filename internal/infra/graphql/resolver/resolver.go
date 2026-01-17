package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

import (
	user "github.com/zghost10/go-best-practices/internal/usecase/user"
)

type Resolver struct {
	CreateUserUseCase *user.CreateUserUseCase
	GetUserUseCase    *user.GetUserUseCase
	ListUsersUseCase  *user.ListUsersUseCase
}
