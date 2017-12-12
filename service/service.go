package service

import (
	"wibble/database"
	"wibble/request"
)

func Name(r request.NameRequest, d database.Server) (string, error) {
	name, _ := d.GetName(r.Id)
	return name, nil
}
