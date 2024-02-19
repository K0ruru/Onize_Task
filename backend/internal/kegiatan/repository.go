package kegiatan

import (
	"context"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Repository interface {
	Get(ctx context.Context, id string) (entity.Kegiatan, error)
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]entity.Kegiatan, error)
	Create(ctx context.Context, kegiatan entity.Kegiatan) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Kegiatan, error) {
	var kegiatan entity.Kegiatan
	err := r.db.With(ctx).Select().Model(id, &kegiatan)
	return kegiatan, err
}

func (r repository) Create(ctx context.Context, kegiatan entity.Kegiatan) error {
	return r.db.With(ctx).Model(&kegiatan).Insert()
}

func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("kegiatan").Row(&count)
	return count, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Kegiatan, error) {
	var kegiatan []entity.Kegiatan
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&kegiatan)
	return kegiatan, err
}
