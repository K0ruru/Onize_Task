package tasks

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

	r.Get("/tasks/<id>", res.get)
	r.Get("/tasks", res.querry)
	r.Get("/tasks/project/<id>", res.get_by_project)


	r.Use(authHandler)

	r.Post("/tasks", res.create)
	r.Put("/tasks/<id>", res.update)
	r.Put("/tasks/status/<id>", res.update_status)
	r.Delete("/tasks/<id>", res.delete)
}

func (r resource) get(c *routing.Context) error {
	task, err := r.service.Get(c.Request.Context(), c.Param("id"))

	if err != nil {
		return err
	}

	return c.Write(task)
}

func (r resource) get_by_project(c *routing.Context) error {
	ctx := c.Request.Context()

	tasks, err := r.service.GetByProject(ctx, c.Param("id"))

	if err != nil {
		return err
	}

	return c.Write(tasks)
}

func (r resource) querry(c *routing.Context) error {
	ctx := c.Request.Context()

	tasks, err := r.service.Querry(ctx)

	if err != nil {
		return err
	}

	return c.Write(tasks)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTask

	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	task, err := r.service.Create(c.Request.Context(), input)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(task, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTask

	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	task, err := r.service.Update(c.Request.Context(), c.Param("id"), input)

	if err != nil {
		return err
	}

	return c.Write(task)
}

func (r resource) update_status(c *routing.Context) error {
	var input UpdateStatus

	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	status, err := r.service.UpdateStatus(c.Request.Context(), c.Param("id"), input)

	if err != nil {
		return err
	}

	return c.Write(status)
}

func (r resource) delete(c *routing.Context) error {
	task, err := r.service.Delete(c.Request.Context(), c.Param("id"))

	if err != nil {
		return err
	}

	return c.Write(task)
	
}
