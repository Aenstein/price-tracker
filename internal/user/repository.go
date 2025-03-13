package user

import "price-tracker/pkg/db"

type UserRepository struct {
	*db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Db: database,
	}
}

func (r *UserRepository) CreateUser(user *User) (*User, error) {
	result := r.Db.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	result := r.Db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}