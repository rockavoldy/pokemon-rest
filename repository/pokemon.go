package repository

import (
	"github.com/rockavoldy/pokemon-rest/model"
	"gorm.io/gorm"
	"log"
)

// Create
func CreatePokemon(db *gorm.DB, pokemon *model.Pokemon) (string,  error) {
	db.Create(pokemon)

	return pokemon.UUID, db.Error
}

// Read list
func ListPokemons(db *gorm.DB, limit int) []model.Pokemon {
	var pokemons []model.Pokemon
	db.Preload("Types").Limit(limit).Find(&pokemons)

	return pokemons
}

// Read one
func GetPokemon(db *gorm.DB, uuid string) model.Pokemon {
	var pokemon model.Pokemon
	db.Preload("Types").Where("uuid = ?", uuid).Find(&pokemon)

	return pokemon
}

// Update
func UpdatePokemon(db *gorm.DB, pokemon *model.Pokemon) (string, error) {
	var poke model.Pokemon
	db.Model(&model.Pokemon{}).Where("uuid = ?", pokemon.UUID).First(&poke)
	err := db.Model(&poke).Association("Types").Delete(&poke.Types)
	if err != nil {
		log.Println(err.Error())
	}

	db.Model(&poke).
		Where("uuid = ?", pokemon.UUID).
		Updates(map[string]interface{}{"name": pokemon.Name, "number": pokemon.Number})

	err = db.Model(&poke).Association("Types").Replace(pokemon.Types)

	return pokemon.UUID, err
}

// Delete
func DeletePokemon(db *gorm.DB, pokemon *model.Pokemon) error {
	db.Model(pokemon).Association("Types").Delete(pokemon.Types)
	db.Unscoped().Delete(pokemon)

	return db.Error
}