package project

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type resource struct {
	service Service
	logger log.Logger
}


func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/projects/<id>", res.get)
	r.Get("/projects", res.query)
	r.Get("/projects/user/<id>", res.get_by_user)


	r.Use(authHandler)

	r.Post("/projects", res.create)
	r.Put("/projects/<id>", res.update)
	r.Delete("/projects/<id>", res.delete)

}

func (r resource) get(c *routing.Context) error {
	project, err := r.service.Get(c.Request.Context(), c.Param("id"))

	if err != nil {
		return err
	}

	return c.Write(project)
}


func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()

	projects, err := r.service.Querry(ctx)
	if err != nil {
			return err
	}

	return c.Write(projects)
}

func (r resource) get_by_user(c *routing.Context) error {
	ctx := c.Request.Context()

	projects, err := r.service.GetByUser(ctx, c.Param("id"))
	if err != nil {
			return err
	}

	return c.Write(projects)
}


func (r resource) create(c *routing.Context) error {
	var input CreateProject

	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	project, err := r.service.Create(c.Request.Context(), input)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(project, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateProject

	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	project, err := r.service.Update(c.Request.Context(), c.Param("id"), input)

	if err != nil {
		return err
	}

	return c.Write(project)
}


func (r resource) delete(c *routing.Context) error {
	project, err := r.service.Delete(c.Request.Context(), c.Param("id"))

	if err != nil {
		return err
	}

	return c.Write(project)
	
}