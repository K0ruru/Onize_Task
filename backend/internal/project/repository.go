package project

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.Project, error)
	Querry(ctx context.Context) ([]entity.Project, error)
	GetByUser(ctx context.Context, id string) ([]entity.Project, error)
	Create(ctx context.Context, project entity.Project) error
	Update(ctx context.Context, project entity.Project) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db 		 *dbcontext.DB
	logger  log.Logger
}

func NewRepo(db *dbcontext.DB, logger log.Logger) repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Project, error){
	var project entity.Project

	err := r.db.With(ctx).Select().Model(id, &project)
	return project, err
}


func (r repository) Querry(ctx context.Context) ([]entity.Project, error) {
	var project []entity.Project

	err := r.db.With(ctx).
	Select().
	OrderBy("id").
	All(&project)

	return project, err
}

func (r repository) GetByUser(ctx context.Context, id string) ([]entity.Project, error) {
	var project []entity.Project

	err := r.db.With(ctx).
	Select().
	Where(dbx.HashExp{"user_id": id}).
	All(&project)

	return project, err
}

func (r repository) Create(ctx context.Context, project entity.Project) error {
	return r.db.With(ctx).Model(&project).Insert()
}

func (r repository) Update(ctx context.Context, project entity.Project) error {
	return r.db.With(ctx).Model(&project).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	project, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&project).Delete()
}
