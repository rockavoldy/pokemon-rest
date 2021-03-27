package model

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
)

type Pokemon struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"unique;not null;"`
	Number string   `json:"number" gorm:"type:integer;not null;unique"`
	Name   string `json:"name" gorm:"not null"`

	Types []TypePokemon `json:"types" gorm:"many2many:pokemon_types;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreatePokemonReq struct {
	UUID string `json:"uuid"`
	Number string `json:"number"`
	Name string `json:"name"`
	Types []string `json:"types"`
}

func (pokemon *Pokemon) BeforeCreate(tx *gorm.DB) (err error) {
	pokemon.UUID = uuid.New().String()

	return err
}

func (pokemon *Pokemon) AfterFind(tx *gorm.DB) (err error) {
	number, _ := strconv.Atoi(pokemon.Number)
	pokemon.Number = fmt.Sprintf("%03d", number)

	return err
}