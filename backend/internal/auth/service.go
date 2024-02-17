package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, input LoginUser) (string, error)

	SignUp(ctx context.Context, input CreateUser) (string, error)
}

// Identity represents an authenticated user identity.
type User struct {
	entity.User
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
	repo            Repository
}

type CreateUser struct {
	Name       string `json:"name"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
}

type LoginUser struct {
	Email string `json:"email"`
	Passphrase string `json:"passphrase"`
}

func (cu CreateUser) Validate() error {
	return validation.ValidateStruct(&cu,
		validation.Field(&cu.Name, validation.Required),
		validation.Field(&cu.Passphrase, validation.Required, validation.Length(8, 128)),
		validation.Field(&cu.Email, validation.Required),
	)
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, logger log.Logger, repo Repository) Service {
	return &service{signingKey, tokenExpiration, logger, repo}
}

// SignUp implements Service.
func (s *service) SignUp(ctx context.Context, req CreateUser) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}
	id := entity.GenerateID()
	err := s.repo.Create(ctx, entity.User{
		ID:         id,
		Name:       req.Name,
		Passphrase: req.Passphrase,
		Email:      req.Email,
	})

	if err != nil {
		return "", err
	}
	return s.generateJWT(User{User: entity.User{
		ID:         id,
		Name:       req.Name,
		Passphrase: req.Passphrase,
		Email:      req.Email,
	}})
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s *service) Login(ctx context.Context, loginReq LoginUser) (string, error) {
	user, err := s.authenticate(ctx, loginReq)
	if err != nil {
		// Handle the error, for example, return an empty token or log the error.
		s.logger.Error("Authentication failed", "error", err)
		return "", errors.Unauthorized("authentication failed")
	}

	return s.generateJWT(user)
}

func (s *service) authenticate(ctx context.Context, loginReq LoginUser) (User, error) {
	user, err := s.repo.GetByEmailAndPassword(ctx, loginReq.Email, loginReq.Passphrase)
	if err != nil {
		// Handle the error, for example, return an empty user or log the error.
		s.logger.Error("Authentication failed", "error", err)
		return User{}, errors.Unauthorized("authentication failed")
	}

	return User{User: user}, nil
}





// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(user User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
