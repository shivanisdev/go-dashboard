package user

import "context"

//Service encapsulates usecase logic for user.
type Service interface {
	Get(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, input CreateUserRequest) (User, error)
}

type service struct {
	repo Repository
}

//CreateUserRequest represents user create request. it will have Validate method later
type CreateUserRequest struct {
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

//NewService create new user service
func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s service) Create(ctx context.Context, input CreateUserRequest) (User, error) {
	err := s.repo.Create(ctx, User{
		UserID: input.UserID,
		Name:   input.Name,
		Email:  input.Email,
	})

	if err != nil {
		return User{}, err
	}

	return s.Get(ctx, input.UserID)
}
