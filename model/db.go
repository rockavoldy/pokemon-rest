package model

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func dsn(dbUser string, dbPass string, dbHost string, dbPort string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
}

func Connect(dbUser string, dbPass string, dbHost string, dbPort string, dbName string) {
	dsn := dsn(dbUser, dbPass, dbHost, dbPort, dbName)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to DB\n", err.Error())
		return
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println(err.Error())
	}

	log.Println("Connected to DB")

	// connect with gorm
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
	}

	DB = gormDB

	DB.AutoMigrate(&TypePokemon{}, &Pokemon{})

	typesMigrate := []TypePokemon{
		{Code: "ghost", Name: "Ghost"},
		{Code: "flying", Name: "Flying"},
		{Code: "fairy", Name: "Fairy"},
		{Code: "normal", Name: "Normal"},
		{Code: "bug", Name: "Bug"},
		{Code: "electric", Name: "Electric"},
		{Code: "ice", Name: "Ice"},
		{Code: "fire", Name: "Fire"},
		{Code: "rock", Name: "Rock"},
		{Code: "psychic", Name: "Psychic"},
		{Code: "ground", Name: "Ground"},
		{Code: "dragon", Name: "Dragon"},
		{Code: "dark", Name: "Dark"},
		{Code: "steel", Name: "Steel"},
		{Code: "fighting", Name: "Fighting"},
		{Code: "poison", Name: "Water"},
		{Code: "water", Name: "Water"},
		{Code: "earth", Name: "Earth"},
	}

	//pokemonsMigrate := []Pokemon{
	//	{Number: "001", Name: "Bulbasaur", Types: },
	//}

	for _, elType := range typesMigrate {
		DB.Create(&elType)
	}

	if err != nil {
		log.Println(err.Error())
	}

}
