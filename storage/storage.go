package storage

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Auth(user *models.User)(int, error)
	GetUser(userId int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(userId int, user *models.UserDate) error
	DeleteUser(userId int) error
}

type storage struct {
	ids int
	Host string
	Port int
	TLL int
	conn *Conn
	db map[int]*models.User
}

func New(host string, port, ttl int, conn *Conn)IStorage {
	return &storage{
		db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

func (s *storage) Auth(user *models.User)(int, error) {
	err := s.conn.Open()
	if err != nil {
		return 0, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	if user == nil{
		return 0, errors.New("Not founded")
	}

	s.ids ++
	user.Id = s.ids
	s.db[s.ids] = user
	log.Printf("user %v is add: %v", s.ids, user)

	for k, v := range s.db {
		if v.Login == user.Login {
			if v.HashPassword == user.HashPassword {
				return k, nil
			} else {
				return 0, errors.New("Wrong password")
			}
		} else {
			return k, errors.New("user exists in the database")
		}
	}

	return s.ids, nil
}


func (s *storage) GetUser(userId int)(*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

		user, isOk := s.db[userId]
		if isOk {
			return user, nil
		}

		return nil, errors.New("User not found")
}

func (s *storage) GetUsers()([]*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	users := make([]*models.User, 0, len(s.db))

	for _, v := range s.db {
		 users = append(users, v)
		 fmt.Println("users", users)
	}

		return users, nil
}

func (s *storage) UpdateUser(userId int, user *models.UserDate) error {
	if user == nil {
		return errors.New("data is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	userDb, isok := s.db[userId]
	
	if !isok {
		return errors.New("fall update user")
	} 

	userDb.Name = user.Name
	userDb.Surname = user.Surname
	userDb.Status = user.Status
	userDb.Role = user.Role
	userDb.UpdateDate = time.Now()
	log.Printf("update user %v", user)

	return nil

}

func (s *storage) DeleteUser(userId int) error {
	err := s.conn.Open()
	if err != nil {
		return err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

		_, isOk := s.db[userId]
		if !isOk {
			return errors.New("not found userId")
		} else {
			delete(s.db, userId)
			return nil
		}
}

func NewConnect() *Conn {
	return &Conn{}
}

type Conn struct {
	val bool
}

func (c *Conn) Close() error {
	if !c.val {
		return errors.New("failed to close")
	}

	return nil
}

func (c *Conn) Open() error {
	if c.val {
		return errors.New("failed to open Conn")
	}

	return nil
}

