package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)


type Service interface{ 
  Create(ctx context.Context, req CreateUserRequest) (User, error)
  Get(ctx context.Context, id string) (User, error)
  Count(ctx context.Context) (int, error)
  Query(ctx context.Context, offset, limit int) ([]User, error)
  Update(ctx context.Context, id string, req UpdateUserRequest) (User, error)  
  Delete(ctx context.Context, id string) (User, error)
}

type User struct {
	entity.Users
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	No_telp  string `json:"no_telp"`
 }

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	No_telp  string `json:"no_telp"`
 }


func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.Email, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.No_telp, validation.Required, validation.Length(0, 128)),
	)
}

func (m UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.Email, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
    validation.Field(&m.No_telp, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new album service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Create creates a new album.
func (s service) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

  id := entity.GenerateID()

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// Convert the hashed password byte slice to a string
	hashedPasswordString := string(hashedPassword)

	err = s.repo.Create(ctx, entity.Users{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPasswordString, // Store the hashed password
		No_telp:   req.No_telp,
	})
	if err != nil {
		return User{}, err
	}
	return s.Get(ctx, id)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range items {
		result = append(result, User{item})
	}
	return result, nil
}

func (s service) Update(ctx context.Context, id string, req UpdateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		return user, err
	}
	user.Name = req.Name
	user.Email = req.Email
	user.No_telp = req.No_telp

	// Hash the password if it's provided in the update request
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, err
		}
		user.Password = string(hashedPassword)
	}

	if err := s.repo.Update(ctx, user.Users); err != nil {
		return user, err
	}
	return user, nil
}

// Delete deletes the album with the specified ID.
func (s service) Delete(ctx context.Context, id string) (User, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return User{}, err
	}
	return user, nil
}


