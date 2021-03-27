package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypePokemon struct {
	gorm.Model
	UUID   string `json:"uuid" gorm:"unique;not null"`
	Code string `json:"code" gorm:"unique"`
	Name string `json:"name"`
}

type CreateTypePokemonReq struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (typePokemon *TypePokemon) BeforeCreate(tx *gorm.DB) (err error) {
	typePokemon.UUID = uuid.New().String()

	return err
}