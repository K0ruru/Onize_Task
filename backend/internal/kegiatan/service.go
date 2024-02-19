package kegiatan

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for albums.
type Service interface {
	Create(ctx context.Context, input CreateKegiatanRequest) (Kegiatan, error)
	Get(ctx context.Context, id string) (Kegiatan, error)
	Count(ctx context.Context) (int, error)
	Query(ctx context.Context, offset, limit int) ([]Kegiatan, error)
}

// Album represents the data about an album.
type Kegiatan struct {
	entity.Kegiatan
}

// CreateAlbumRequest represents an album creation request.
type CreateKegiatanRequest struct {
	Name       string    `json:"name"`
	Deskripsi  string    `json:"deskripsi"`
	DeadlineAt time.Time `json:"deadline_at"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateKegiatanRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Length(0, 128)),
		validation.Field(&m.Deskripsi, validation.Length(0, 500)),
		validation.Field(&m.DeadlineAt, validation.Required),
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

func (s service) Get(ctx context.Context, id string) (Kegiatan, error) {
	kegiatan, err := s.repo.Get(ctx, id)
	if err != nil {
		return Kegiatan{}, err
	}
	return Kegiatan{kegiatan}, nil
}

func (s service) Create(ctx context.Context, req CreateKegiatanRequest) (Kegiatan, error) {
	if err := req.Validate(); err != nil {
		return Kegiatan{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Kegiatan{
		ID:         id,
		Name:       req.Name,
		Deskripsi:  req.Deskripsi,
		CreatedAt:  now,
		DeadlineAt: req.DeadlineAt,
	})
	if err != nil {
		return Kegiatan{}, err
	}
	return s.Get(ctx, id)
}

func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

func (s service) Query(ctx context.Context, offset, limit int) ([]Kegiatan, error) {
	Items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Kegiatan{}
	for _, item := range Items {
		result = append(result, Kegiatan{item})
	}
	return result, nil
}
