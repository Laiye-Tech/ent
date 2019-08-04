// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/migrate/entv1/user"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age     *int32
	name    *string
	address *string
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int32) *UserCreate {
	uc.age = &i
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetAddress sets the address field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.address = &s
	return uc
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("entv1: missing required field \"age\"")
	}
	if uc.name == nil {
		return nil, errors.New("entv1: missing required field \"name\"")
	}
	if err := user.NameValidator(*uc.name); err != nil {
		return nil, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
	}
	switch uc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return uc.sqlSave(ctx)
	default:
		return nil, errors.New("entv1: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		res sql.Result
		u   = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(user.Table).Default(uc.driver.Dialect())
	if uc.age != nil {
		builder.Set(user.FieldAge, *uc.age)
		u.Age = *uc.age
	}
	if uc.name != nil {
		builder.Set(user.FieldName, *uc.name)
		u.Name = *uc.name
	}
	if uc.address != nil {
		builder.Set(user.FieldAddress, *uc.address)
		u.Address = *uc.address
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}