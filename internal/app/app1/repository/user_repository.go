package repository

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/cryptoz"
	"github.com/pepeunlimited/users/internal/app/app1/ent"
	"github.com/pepeunlimited/users/internal/app/app1/ent/user"
	"time"
)

type Role string
const (
	User Role = "User"
	Admin 	  = "Admin"
)

var (
	ErrUserBanned 		= errors.New("users: user is banned")
	ErrUserLocked 		= errors.New("users: user is locked")
	ErrUserDeleted    	= errors.New("users: user is deleted")
	ErrUsernameExist    = errors.New("users: username already exist")
	ErrEmailExist 		= errors.New("users: email already exist")
)

// Access the `users` table
// 	- one-to-many to `tickets`
type UserRepository interface {
	CreateUser(ctx context.Context, username string, email string, password string, role Role) (*ent.User, error)
	GetUserById(ctx context.Context, id int) (*ent.User, error)
	GetUserByUsername(ctx context.Context, username string) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	GetUsers(ctx context.Context, offset int, limit int) ([]*ent.User, error)
	UpdateUser(ctx context.Context, user *ent.UserUpdateOne) (*ent.User, error)
	UpdatePassword(ctx context.Context, userId int, current string, new string) (*ent.User, error)
	DeleteUsers(ctx context.Context)
	DeleteUser(ctx context.Context, userId int) error
	BanUser(ctx context.Context, userId int) error
	LockUser(ctx context.Context, userId int) error
	GetUserTicketsByUserId(ctx context.Context, userId int) (*ent.User, []*ent.Ticket, error)
	// wipes `users` and `tickets` database tables
	DeleteAll(ctx context.Context)

	UnbanUser(ctx context.Context, userId int) (*ent.User, error)
	UnlockUser(ctx context.Context, userId int) (*ent.User, error)
	UndoDeleteByUserId(ctx context.Context, userId int) (*ent.User, error)
}

type userMySQL struct {
	client *ent.Client
	crypto cryptoz.Crypto
}

func (repo userMySQL) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return repo.client.User.Query().Where(user.Username(username)).Only(ctx)
}

func (repo userMySQL) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return repo.client.User.Query().Where(user.Email(email)).Only(ctx)
}

func (repo userMySQL) UnbanUser(ctx context.Context, userId int) (*ent.User, error) {
	_, err := repo.client.User.Update().SetIsBanned(false).SetLastModified(time.Now().UTC()).Where(user.ID(userId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return repo.GetUserById(ctx, userId)
}

func (repo userMySQL) UnlockUser(ctx context.Context, userId int) (*ent.User, error) {
	_, err := repo.client.User.Update().SetIsLocked(false).SetLastModified(time.Now().UTC()).Where(user.ID(userId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return repo.GetUserById(ctx, userId)
}

func (repo userMySQL) UndoDeleteByUserId(ctx context.Context, userId int) (*ent.User, error) {
	_, err := repo.client.User.Update().SetIsDeleted(false).SetLastModified(time.Now().UTC()).Where(user.ID(userId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return repo.GetUserById(ctx, userId)
}

func (repo userMySQL) UpdatePassword(ctx context.Context, userId int, current string, new string) (*ent.User, error) {
	user, err := repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	err = repo.crypto.Check(user.Password, current)
	if err != nil {
		return nil, err
	}
	hash, err := repo.crypto.Hash(new)
	if err != nil {
		return nil, err
	}
	return repo.UpdateUser(ctx, user.Update().SetPassword(hash))
}

func (repo userMySQL) BanUser(ctx context.Context, userId int) error {
	user, err := repo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	return user.Update().SetIsBanned(true).SetLastModified(time.Now().UTC()).Exec(ctx)
}

func (repo userMySQL) LockUser(ctx context.Context, userId int) error {
	user, err := repo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	return user.Update().SetIsLocked(true).SetLastModified(time.Now().UTC()).Exec(ctx)
}

func (repo userMySQL) DeleteUser(ctx context.Context, userId int) error {
	user, err := repo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	return user.Update().SetIsDeleted(true).SetLastModified(time.Now().UTC()).Exec(ctx)
}

func (repo userMySQL) DeleteAll(ctx context.Context) {
	repo.client.Ticket.Delete().ExecX(ctx)
	repo.client.User.Delete().ExecX(ctx)
}

func (repo userMySQL) GetUserTicketsByUserId(ctx context.Context, userId int) (*ent.User, []*ent.Ticket, error) {
	user, err := repo.client.User.Get(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	tickets, err := user.QueryTickets().All(ctx)
	if err != nil {
		return nil, nil, err
	}
	return user, tickets, nil
}

func (repo userMySQL) CreateUser(ctx context.Context, username string, email string, password string, role Role) (*ent.User, error) {
	if _, err := repo.GetUserByUsername(ctx, username); err == nil {
		return nil, ErrUsernameExist
	}
	if _, err := repo.GetUserByEmail(ctx, email); err == nil {
		return nil, ErrEmailExist
	}
	password, err := repo.crypto.Hash(password)
	save, err := repo.client.
		User.
		Create().
		SetRole(string(role)).
		SetUsername(username).
		SetEmail(email).
		SetPassword(password).
		SetLastModified(time.Now().UTC()).
		SetIsDeleted(false).
		SetIsBanned(false).
		SetIsLocked(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (repo userMySQL) GetUserById(ctx context.Context, id int) (*ent.User, error) {
	user, err := repo.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if user.IsBanned {
		return nil, ErrUserBanned
	}
	if user.IsLocked {
		return nil, ErrUserLocked
	}
	if user.IsDeleted {
		return nil, ErrUserDeleted
	}
	return user, nil
}

func (repo userMySQL) GetUsers(ctx context.Context, offset int, limit int) ([]*ent.User, error) {
	return repo.client.User.Query().Limit(limit).Offset(offset).Where(user.IsDeleted(false), user.IsLocked(false), user.IsBanned(false)).All(ctx)
}

func (repo userMySQL) UpdateUser(ctx context.Context, user *ent.UserUpdateOne) (*ent.User, error) {
	updated, err := user.SetLastModified(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (repo userMySQL) DeleteUsers(ctx context.Context) {
	repo.client.User.Delete().ExecX(ctx)
}

func NewUserRepository(client *ent.Client) UserRepository {
	return userMySQL{client:client, crypto: cryptoz.NewCrypto()}
}