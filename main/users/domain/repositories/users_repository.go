package repositories

type UsersRepository interface {
	Create() error
	FindAll() error
	FindById() error
	Update() error
	Delete() error
}
