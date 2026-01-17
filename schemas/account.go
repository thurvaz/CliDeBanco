package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Name     string
	Cpf      string
	Password string
}

type OpeningResponse struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
}
