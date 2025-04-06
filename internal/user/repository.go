package user

import "api/pkg/db"

// создание пользователя, поиско по емэйл
type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email=?", email) //DB.First возвращает первое совпадение
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
