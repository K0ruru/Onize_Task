package auth

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.User, error)
	Create(ctx context.Context, user entity.User) error
	GetByEmailAndPassword(ctx context.Context, email string, passphrase string) (entity.User, error)
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func (r repository) Create(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Insert()
}

func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) GetByEmailAndPassword(ctx context.Context, email string, passphrase string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("user").Where(dbx.HashExp{"email": email, "passphrase": passphrase}).One(&user)
	return user, err
}



func NewRepo(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}
