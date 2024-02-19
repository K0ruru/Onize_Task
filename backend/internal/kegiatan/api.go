package kegiatan

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/kegiatan/<id>", res.get)
	r.Get("/kegiatan", res.query)

	r.Use(authHandler)

	r.Post("/kegiatan", res.create)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	kegiatan, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.Write(kegiatan)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	kegiatan, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = kegiatan
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateKegiatanRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	kegiatan, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(kegiatan, http.StatusCreated)
}
