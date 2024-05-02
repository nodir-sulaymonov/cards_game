package service

import (
	"log"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(user *models.User)(int, error)
	GetUser(userId int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(userId int, user *models.UserDate) error
	DeleteUser(userId int) error
	GetStatistics() map[string]int
}

type service struct {
	CounterAuth int
	CounterGetUser int
	CounterGetUsers int
	CounterUpdateUser int
	CounterDeleteUser int
	repo        storage.IStorage
}

func New(repo storage.IStorage)IService{
	return &service{
		CounterAuth: 0,
		CounterGetUser: 0,
		CounterGetUsers: 0,
		CounterDeleteUser: 0,
		repo: repo,
	}
}

func (s *service)Auth(user *models.User)(int, error){
	result, err := s.repo.Auth(user)
	if err != nil {
		return 0, err
	}

	s.CounterAuth++

	return result, nil
}

func (s *service)GetUser(userId int)(*models.User, error){
	result, err := s.repo.GetUser(userId)
	if err != nil {
		return nil, err
	}

	s.CounterGetUser++

	return result, nil
}

func (s *service)GetUsers()([]*models.User, error){
	result, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	s.CounterGetUsers++

	return result, nil
}

func (s *service)UpdateUser(userId int, user *models.UserDate) error {
	err := s.repo.UpdateUser(userId, user)
	if err != nil {
		return  err
	}

	s.CounterUpdateUser++

	return nil
}

func (s *service)DeleteUser(userId int) error {
	err := s.repo.DeleteUser(userId)
	if err != nil {
		return err
	}

	s.CounterDeleteUser++

	return nil
}

func (s *service)GetStatistics() map[string]int {
	userStats := make(map[string]int)
	
	userStats["AuthStats"] = s.CounterAuth
	userStats["GetUserStats"] = s.CounterGetUser
	userStats["GetUsersStats"] = s.CounterGetUsers
	userStats["UpdateUserStats"] = s.CounterUpdateUser
	userStats["DeleteUserStats"] = s.CounterDeleteUser

	log.Println("statistics: ", userStats)

	return userStats
}