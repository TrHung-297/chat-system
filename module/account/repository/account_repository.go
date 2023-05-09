package repository

import "github.com/jmoiron/sqlx"

type IAccountRepository interface {

}

type AccountRepository struct {
	*sqlx.DB
	//ProfileRepository repository.IProfileRepository
}

func NewAccountRepository(sqlx *sqlx.DB) IAccountRepository {
	accountRepo := AccountRepository{
		sqlx,
	}
	//accountRepo.ProfileRepository = repository.NewProfileRepository(sqlx)
	return &accountRepo
}