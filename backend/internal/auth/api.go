package auth

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	rg.Post("/login", login(service, logger))
	rg.Post("/signup", res.signup)
}

type resource struct {
	service Service
	logger log.Logger
}

// login returns a handler that handles user login request.
func login(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var input LoginUser
		if err := c.Read(&input); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}
		token, err := service.Login(c.Request.Context(), input)
		if err != nil {
			return err
		}
		return c.Write(struct {
			Token string `json:"token"`
		}{token})
	}
}

func (r resource) signup(c *routing.Context) error {
	var input CreateUser
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	token, err := r.service.SignUp(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.Write(struct {
		Token string `json:"token"`
	}{token})
}

