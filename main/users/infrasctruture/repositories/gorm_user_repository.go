package repositories

import "gorm.io/gorm"

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (itself *GormUserRepository) Create() error {
	return nil
}

func (itself *GormUserRepository) FindAll() error {
	return nil
}
func (itself *GormUserRepository) FindById() error {

	return nil
}
func (itself *GormUserRepository) Update() error {
	return nil
}
func (itself *GormUserRepository) Delete() error {
	return nil
}
