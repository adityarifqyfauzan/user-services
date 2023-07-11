package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/adityarifqyfauzan/user-services/app/validator"
	"github.com/gosimple/slug"
	"github.com/oklog/ulid/v2"
)

type Role struct {
	ID        ulid.ULID
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m Role) TableName() string {
	return "roles"
}

func NewRole(name string) (Role, error) {
	if err := validator.ValidateRole(name); err != nil {
		return Role{}, err
	}

	role := Role{
		ID:        ulid.Make(),
		Name:      name,
		Slug:      slug.Make(name),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return role, nil
}

func (m Role) BeforeCreate(tx *sql.Tx) error {
	var id uint8

	row := tx.QueryRow(fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE slug = $1 LIMIT 1", m.TableName()), m.Slug)
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	if id != 0 {
		return fmt.Errorf("role %s is exists", m.Name)
	}

	return nil
}

func (m Role) BeforeUpdate(tx *sql.Tx) error {
	// set updated at value
	m.UpdatedAt = time.Now()

	// check if role is exist
	var id uint8
	row := tx.QueryRow(fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE slug = $1 LIMIT 1", m.TableName()), m.Slug)
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	if id > 0 {
		return fmt.Errorf("role %s is exists", m.Name)
	}

	return nil
}

func (m Role) MarshalJson() ([]byte, error) {
	var role struct {
		ID   ulid.ULID `json:"id"`
		Name string    `json:"name"`
		Slug string    `json:"slug"`
	}

	role.ID = m.ID
	role.Name = m.Name
	role.Slug = m.Slug

	return json.Marshal(role)
}
