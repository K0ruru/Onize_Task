package auth

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)


type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

type Repository interface {
    // GetUserByUsername retrieves a user by username
    GetUserByUsername(ctx context.Context, username string) (*entity.Users, error)}


// GetUserByUsername retrieves a user by their username from the database.
func (r repository) GetUserByUsername(ctx context.Context, username string) (*entity.Users, error) {
    var user entity.Users
    
    // Construct the condition expression for the WHERE clause
    condition := dbx.HashExp{"name": username}

    // Use Select() to get all columns for the user
    query := r.db.With(ctx).Select().From("users").Where(condition)

    // Use One() instead of Row() to retrieve a single row
    err := query.One(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

