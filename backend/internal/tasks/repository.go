package tasks

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.Task, error)
	Querry(ctx context.Context) ([]entity.Task, error)
	Create(ctx context.Context, task entity.Task) (entity.Task, error)
	Update(ctx context.Context, task entity.Task) (entity.Task, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db 	  	*dbcontext.DB
	logger 	 log.Logger
}

func NewRepo(db *dbcontext.DB, logger log.Logger) repository {
	return repository{db, logger}
}

func(r repository) Get(ctx context.Context, id string) (entity.Task, error) {
	var task entity.Task

	err := r.db.With(ctx).Select().Model(id, &task)

	return task, err
}

func (r repository) Querry(ctx context.Context) ([]entity.Task, error) {
	var task []entity.Task

	err := r.db.With(ctx).Select().OrderBy("id").All(&task)

	return task, err
}

func (r repository) Create(ctx context.Context, task entity.Task) error {
	return r.db.With(ctx).Model(&task).Insert()
}

func (r repository) Update(ctx context.Context, task entity.Task) error {
	return r.db.With(ctx).Model(&task).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	task, err := r.Get(ctx, id)

	if err != nil {
		return err
	}

	return r.db.With(ctx).Model(&task).Delete()
}