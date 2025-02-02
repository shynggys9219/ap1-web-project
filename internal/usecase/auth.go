package usecase

import (
	"github.com/shynggys9219/ap1-web-project/internal/model"
)

type Auth struct {
	authRepo AuthRepo
}

func NewAuth(authRepo AuthRepo) *Auth {
	return &Auth{
		authRepo: authRepo,
	}
}

func (uc *Auth) Create(user model.User) (int, error) {
	return uc.authRepo.Create(user)
}

func (uc *Auth) Get(id int) (*model.User, error) {
	return uc.authRepo.Get(id)
}

func (uc *Auth) Update(id int) (*model.User, error) {
	return nil, nil
}

func (uc *Auth) Delete(id int) error {
	return nil
}
