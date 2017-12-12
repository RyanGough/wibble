package database

import (
	"errors"
)

type Server interface {
	GetName(id int) (string, error)
}

type MySqlServer struct{}

func (s MySqlServer) GetName(id int) (string, error) {
	if id != 1234 {
		return "", errors.New("unknown id")
	}
	return "Ray Purchase", nil
}
