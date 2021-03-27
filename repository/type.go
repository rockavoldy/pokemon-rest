package repository

import (
	"github.com/rockavoldy/pokemon-rest/model"
	"gorm.io/gorm"
)

// Create
func CreateType(db *gorm.DB, typePokemon *model.TypePokemon) (string, error) {
	result := db.Create(typePokemon)

	return typePokemon.UUID, result.Error
}

// Read list
func ListTypes(db *gorm.DB, limit int) []model.TypePokemon {
	var typePokemon []model.TypePokemon
	db.Limit(limit).Find(&typePokemon)

	return typePokemon
}

// Read one
func GetType(db *gorm.DB, uuid string) model.TypePokemon {
	var typePokemon model.TypePokemon
	db.Where("uuid = ?", uuid).Find(&typePokemon)

	return typePokemon
}

// Update
func UpdateType(db *gorm.DB, typePokemon *model.TypePokemon) (string, error) {
	result := db.Model(typePokemon).Where("uuid = ?", typePokemon.UUID).
		Updates(map[string]interface{}{"name": typePokemon.Name, "code": typePokemon.Code})

	return typePokemon.UUID, result.Error
}

// Delete
func DeleteType(db *gorm.DB, typePokemon *model.TypePokemon) error {
	result := db.Where("id = ?", typePokemon.ID).Delete(typePokemon)

	return result.Error
}