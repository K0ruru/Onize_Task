package project

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Project struct {
	entity.Project
}

type CreateProject struct {
	Title     string    `json:"title"`
	User_id   string    `json:"user_id"`
}

type UpdateProject struct {
	Title string `json:"title"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

type Service interface {
	Get(ctx context.Context, id string) (Project, error)
	Querry(ctx context.Context) ([]Project, error)
	Create(ctx context.Context, input CreateProject) (Project, error)
	Update(ctx context.Context, id string, input UpdateProject) (Project, error)
	Delete(ctx context.Context, id string) (Project, error)
}

func (cp CreateProject) Validate() error {
	return validation.ValidateStruct(&cp,
		validation.Field(&cp.Title, validation.Required, validation.Length(0, 25)),
		validation.Field(&cp.User_id, validation.Required),
	)
}

func (up UpdateProject) Validate() error {
	return validation.ValidateStruct(&up,
		validation.Field(&up.Title, validation.Required, validation.Length(0, 25)),
	)
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get implements Service.
func (s service) Get(ctx context.Context, id string) (Project, error) {
	project, err := s.repo.Get(ctx, id)

	if err != nil {
		return Project{}, err
	}

	return Project{project}, nil
}

// Querry implements Service.
func (s service) Querry(ctx context.Context) ([]Project, error) {
	projects, err := s.repo.Querry(ctx)

	if err != nil {
		return nil, err
	}

	result := []Project{}

	for _, project := range projects {
		result = append(result, Project{project})
	}

	return result, nil
}

// Create implements Service.
func (s service) Create(ctx context.Context, input CreateProject) (Project, error) {
	if err := input.Validate(); err != nil {
		return Project{}, err
	}

	id := entity.GenerateID()
	time := time.Now()
	err := s.repo.Create(ctx, entity.Project{
		ID: id,
		Title: input.Title,
		User_id: input.User_id,
		CreatedAt: time,
	})

	if err != nil {
		return Project{}, err
	}

	return s.Get(ctx, id)
}

// Delete implements Service.
func (s service) Delete(ctx context.Context, id string) (Project, error) {
	project, err := s.Get(ctx, id)
	if err != nil {
		return Project{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Project{}, err
	}
	return project, nil
}


// Update implements Service.
func (s service) Update(ctx context.Context, id string, input UpdateProject) (Project, error) {
	if err := input.Validate(); err != nil {
		return Project{}, err
	}

	project, err := s.Get(ctx, id)
	if err != nil {
		return project, err
	}

	project.Title = input.Title

	if err := s.repo.Update(ctx, project.Project); err != nil {
		return project, err
	}
	return project, nil
}
