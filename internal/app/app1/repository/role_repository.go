package repository

import (
	"context"
	"errors"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/ent/role"
	"github.com/pepeunlimited/users/internal/app/app1/ent/user"
)

type Role string
const (
	User Role = 	"user"
	Admin 	  = 	"admin"
	Reviewer  =		"reviewer"
)


var (
	ErrDuplicateRole = errors.New("roles: inserted role already exist for the user")
)

// Access the `roles` table
// 	- many-to-one to `users`
type RoleRepository interface {
	AddRole(ctx context.Context, userId int, role Role) error
	DeleteRole(ctx context.Context, userId int, role Role) error
	FindRolesByUserId(ctx context.Context, userId int) ([]*ent.Role, error)
}

type rolesMySQL struct {
	client *ent.Client
}

func (r rolesMySQL) FindRolesByUserId(ctx context.Context, userId int) ([]*ent.Role, error) {
	return r.client.Role.Query().Where(role.HasUsersWith(user.ID(userId))).All(ctx)
}

func (r rolesMySQL) AddRole(ctx context.Context, userId int, role Role) error {
	roles, err := r.FindRolesByUserId(ctx, userId)
	if err != nil {
		return err
	}
	for _, r := range roles {
		if r.Role == string(role) {
			return ErrDuplicateRole
		}
	}
	_, err = r.client.Role.Create().SetUsersID(userId).SetRole(string(role)).Save(ctx)
	return err
}

func (r rolesMySQL) DeleteRole(ctx context.Context, userId int, srole Role) error {
	_, err := r.client.Role.Delete().Where(role.Role(string(srole)), role.HasUsersWith(user.ID(userId))).Exec(ctx)
	return err
}

func NewRolesRepository(client *ent.Client) RoleRepository {
	return &rolesMySQL{client:client}
}