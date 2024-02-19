package tasks

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Task struct {
	entity.Task
}

type CreateTask struct {
	Title      string    `json:"title"`
	Task_desc  string    `json:"task_desc"`
	Label      string    `json:"label"`
	Priority   string    `json:"priority"`
	Tags       string    `json:"tags"`
	Project_id string    `json:"project_id"`
	Status     string    `json:"status" default:"not started"`
	Due        time.Time `json:"due"`
}

type UpdateTask struct {
	Title     string    `json:"title"`
	Task_desc string    `json:"task_desc"`
	Label     string    `json:"label"`
	Priority  string    `json:"priority"`
	Tags      string    `json:"tags"`
	Due       time.Time `json:"due"`
}

type UpdateStatus struct {
	Status string `json:"status"`
}

type service struct {
	repo   repository
	logger log.Logger
}

type Service interface {
	Get(ctx context.Context, id string) (Task, error)
	GetByProject(ctx context.Context, project_id string) ([]Task, error)
	Querry(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, input CreateTask) (Task, error)
	Update(ctx context.Context, id string, input UpdateTask) (Task, error)
	UpdateStatus(ctx context.Context, id string, input UpdateStatus) (Task, error)
	Delete(ctx context.Context, id string) (Task, error)
}

func (ct CreateTask) Validate() error {
	return validation.ValidateStruct(&ct,
		validation.Field(&ct.Title, validation.Required),
		validation.Field(&ct.Task_desc, validation.Required),
		validation.Field(&ct.Label, validation.Required),
		validation.Field(&ct.Priority, validation.Required),
		validation.Field(&ct.Tags, validation.Required),
		validation.Field(&ct.Project_id, validation.Required),
	)
}

func (ut UpdateTask) Validate() error {
	return validation.ValidateStruct(&ut,
		validation.Field(&ut.Title, validation.Required),
		validation.Field(&ut.Task_desc, validation.Required),
		validation.Field(&ut.Label, validation.Required),
		validation.Field(&ut.Priority, validation.Required),
		validation.Field(&ut.Tags, validation.Required),
	)
}

func (us UpdateStatus) Validate() error {
	return validation.ValidateStruct(&us,
		validation.Field(&us.Status, validation.Required),
	)
}

func NewService(repo repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get implements Service.
func (s service) Get(ctx context.Context, id string) (Task, error) {
	task, err := s.repo.Get(ctx, id)

	if err != nil {
		return Task{}, err
	}

	return Task{task}, nil
}

// GetByProject implements Service
func(s service) GetByProject(ctx context.Context, project_id string) ([]Task, error) {
	tasks, err := s.repo.GetByProject(ctx, project_id)

	if err != nil {
		return []Task{}, err
	}

	result := []Task{}

	for _, task := range tasks {
		result = append(result, Task{task})
	}

	return result, nil
}

// Querry implements Service.
func (s service) Querry(ctx context.Context) ([]Task, error) {
	tasks, err := s.repo.Querry(ctx)

	if err != nil {
		return nil, err
	}

	result := []Task{}

	for _, task := range tasks {
		result = append(result, Task{task})
	}

	return result, nil
}

// Create implements Service.
func (s service) Create(ctx context.Context, input CreateTask) (Task, error) {
	if err := input.Validate(); err != nil {
		return Task{}, err
	}

	id := entity.GenerateID()
	time := time.Now()

	if input.Status == "" {
		input.Status = "not started"
}

	err := s.repo.Create(ctx, entity.Task{
		ID: id,
		Title: input.Title,
		Task_desc: input.Task_desc,
		Label: input.Label,
		Priority: input.Priority,
		Status: input.Status,
		Tags: input.Tags,
		Project_id: input.Project_id,
		CreatedAt: time,
		Due: input.Due,
	})

	if err != nil {
		return Task{}, err
	}

	return s.Get(ctx, id)
}

// Update implements Service.
func (s service) Update(ctx context.Context, id string, input UpdateTask) (Task, error) {
	if err := input.Validate(); err != nil {
		return Task{}, err
	}

	task, err := s.Get(ctx, id)

	if err != nil {
		return task, err
	}

	task.Title = input.Title
	task.Task_desc = input.Task_desc
	task.Label = input.Label
	task.Priority = input.Priority
	task.Tags =  input.Tags
	task.Due = input.Due

	if err := s.repo.Update(ctx, task.Task); err != nil {
		return task, err
	}

	return task, nil

}

// UpdateStatus implements Service.
func (s service) UpdateStatus(ctx context.Context, id string, input UpdateStatus) (Task, error) {
	if err := input.Validate(); err != nil {
		return Task{}, err
	}

	task, err := s.Get(ctx, id)

	if err != nil {
		return task, err
	}

	task.Status = input.Status

	if err := s.repo.Update(ctx, task.Task); err != nil {
		return task, err
	}

	return task, nil
}

// Delete implements Service.
func (s service) Delete(ctx context.Context, id string) (Task, error) {
	task, err := s.Get(ctx, id)

	if err != nil {
		return Task{}, err
	}

	if err = s.repo.Delete(ctx, id); err != nil {
		return Task{}, err
	}

	return task, nil
}


