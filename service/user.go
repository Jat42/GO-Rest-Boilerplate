package service

import (
	"users-rest/model"
	"users-rest/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user model.User) (*model.User, error) {
	return s.repo.Insert(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
