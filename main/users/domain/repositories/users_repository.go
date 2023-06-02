package repositories

type IUsersRepository interface {
	Create() error
	FindAll() error
	FindById() error
	Update() error
	Delete() error
}
