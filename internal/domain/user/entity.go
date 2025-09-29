package user

import (
	"context"
	"time"

	"github.com/Nezent/mig-test/internal/application/dto"
	"github.com/Nezent/mig-test/internal/domain/shared"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `json:"id" bun:",nullzero"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"-" bun:"password_hash"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

type UserRepository interface {
	CreateUser(user *User) (uuid.UUID, *shared.DomainError)
	GetUser() (*[]User, *shared.DomainError)
}

type UserService interface {
	CreateUser(req *dto.CreateUserRequest) (*dto.CreateUserResponse, *shared.DomainError)
	GetUser() (*dto.GetUserResponse, *shared.DomainError)
}

// BeforeAppendModel sets timestamps before insert/update.
func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	now := time.Now().UTC()
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = now
		u.UpdatedAt = now
	case *bun.UpdateQuery:
		u.UpdatedAt = now
	}
	return nil
}
