package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	
	"github.com/qiangxue/go-rest-api/internal/errors"

	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, username, password string) (string, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetName() string
}

type service struct {
	repo            Repository
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(repo Repository, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return &service{repo: repo, signingKey: signingKey, tokenExpiration: tokenExpiration, logger: logger}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, username, password string) Identity {
    logger := s.logger.With(ctx, "user", username)

    // Query the database to retrieve the user based on the username
    user, err := s.repo.GetUserByUsername(ctx, username)
    if err != nil {
        logger.Errorf("error retrieving user from the database: %v", err)
        return nil
    }

    // If no user found with the provided username, return authentication failed
    if user == nil || user.ID == "" {
        logger.Infof("authentication failed: user '%s' not found", username)
        return nil
    }

    // Compare the provided password with the hashed password stored in the database
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        logger.Infof("authentication failed: password mismatch for user '%s'", username)
        return nil
    }

    logger.Infof("authentication successful")
    return user
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
