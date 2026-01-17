package user

type IUserRepository interface {
	Create(user User) error
	Get(identifier string) (User, error)
	GetAll() ([]User, error)
	Update(user User) error
	Delete(identifier string) error
}
