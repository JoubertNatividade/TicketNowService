package repositories

type DbUserRepository struct {
}

func NewDbUserRepository() DbUserRepository {
	return DbUserRepository{}
}

func (itself DbUserRepository) Create() error {
	return nil
}

func (itself DbUserRepository) FindAll() error {
	return nil
}
func (itself DbUserRepository) FindById() error {
	return nil
}
func (itself DbUserRepository) Update() error {
	return nil
}
func (itself DbUserRepository) Delete() error {
	return nil
}
