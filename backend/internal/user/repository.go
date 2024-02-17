package user

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface { 
  Get(ctx context.Context, id string) (entity.Users, error)
  GetUserByUsername(ctx context.Context, username string) (entity.Users, error)
  Count(ctx context.Context) (int, error)
  Query(ctx context.Context, offset, limit int) ([]entity.Users, error)
  Create(ctx context.Context, user entity.Users) error
  Update(ctx context.Context, user entity.Users) error
  Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

 // Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Users, error) {
	var user entity.Users
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) GetUserByUsername(ctx context.Context, username string) (entity.Users, error) {
    var user entity.Users
    
    // Construct the condition expression for the WHERE clause
    condition := dbx.HashExp{"name": username}

    err := r.db.With(ctx).Select().From("users").Where(condition).Row(&user)
    return user, err
}


// Count returns the number of the album records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("users").Row(&count)
	return count, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&users)
	return users, err
}

func (r repository) Create(ctx context.Context, user entity.Users) error {
	return r.db.With(ctx).Model(&user).Insert()
}

// Update saves the changes to an album in the database.
func (r repository) Update(ctx context.Context, user entity.Users) error {
	return r.db.With(ctx).Model(&user).Update()
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	user, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}




//
