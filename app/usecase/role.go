package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adityarifqyfauzan/user-services/app/domain/model"
	"github.com/adityarifqyfauzan/user-services/app/domain/repository"
	"github.com/adityarifqyfauzan/user-services/app/errors"
)

type RoleUsecase struct {
	db   *sql.DB
	repo repository.RoleRepository
}

func NewRoleUsecase(db *sql.DB, repo repository.RoleRepository) RoleUsecase {
	return RoleUsecase{
		db:   db,
		repo: repo,
	}
}

var (
	errRoleSaveUsecase    = "unable to save new role"
	errRoleFindAllUsecase = "unable to find all roles"
)

func (uc RoleUsecase) Save(ctx context.Context, name string) (*model.Role, error) {

	newRole, err := model.NewRole(name)
	if err != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("%v: %v", errRoleSaveUsecase, err))
	}

	tx, err := uc.db.Begin()
	if err != nil {
		return nil, errors.NewUnprocessabelEntity(fmt.Sprintf("unable to begin transaction: %v", err))
	}
	defer tx.Rollback()

	err = uc.repo.Save(ctx, tx, newRole)
	if err != nil {
		return nil, errors.NewUnprocessabelEntity(fmt.Sprintf("%v: %v", errRoleSaveUsecase, err))
	}

	tx.Commit()
	return &newRole, nil
}

func (uc RoleUsecase) FindAll(ctx context.Context, page, size int) (any, error) {

	tx, err := uc.db.Begin()
	if err != nil {
		return nil, errors.NewUnprocessabelEntity(fmt.Sprintf("unable to begin transaction: %v", err))
	}
	defer tx.Rollback()
	roles, err := uc.repo.FindAll(ctx, tx, page, size)
	if err != nil {
		return nil, errors.NewUnprocessabelEntity(fmt.Sprintf("%v: %v", errRoleFindAllUsecase, err))
	}
	tx.Commit()
	return roles, nil
}
