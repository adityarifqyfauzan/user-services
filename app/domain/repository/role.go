package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/adityarifqyfauzan/user-services/app/domain/model"
	"github.com/adityarifqyfauzan/user-services/utils"
	"github.com/oklog/ulid/v2"
)

var (
	errRoleNotFound = errors.New("role: not found")
)

type RoleRepository struct {
}

func NewRoleRepository() RoleRepository {
	return RoleRepository{}
}

func (rps RoleRepository) FindByID(ctx context.Context, tx *sql.Tx, id ulid.ULID) (model.Role, error) {

	m := new(model.Role)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", m.TableName())

	row := tx.QueryRowContext(ctx, query, id)
	if err := row.Scan(&m.ID, &m.Name, &m.Slug, &m.CreatedAt, &m.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return *m, errRoleNotFound
		}
		return *m, err
	}

	return *m, nil
}

func (rps RoleRepository) FindAll(ctx context.Context, tx *sql.Tx, page, size int) ([]model.Role, error) {

	// get pagination
	offset := utils.GetOffset(page, size)

	var m model.Role
	roles := make([]model.Role, 0)

	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", m.TableName())
	rows, err := tx.QueryContext(ctx, query, size, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, m)
	}

	return roles, nil
}

func (rps RoleRepository) Save(ctx context.Context, tx *sql.Tx, m model.Role) error {

	err := m.BeforeCreate(tx)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (id, name, slug, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", m.TableName())

	_, err = tx.ExecContext(ctx, query, m.ID, m.Name, m.Slug, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (rps RoleRepository) Update(ctx context.Context, tx *sql.Tx, m model.Role) error {

	err := m.BeforeUpdate(tx)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET name = $2, slug = $3, update_at = $4 WHERE id = $5", m.TableName())
	_, err = tx.ExecContext(ctx, query, m.Name, m.Slug, m.UpdatedAt, m.ID)
	if err != nil {
		return err
	}

	return nil
}
