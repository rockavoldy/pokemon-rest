package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rockavoldy/pokemon-rest/model"
	"github.com/rockavoldy/pokemon-rest/repository"
	"log"
	"net/http"
	"strconv"
)

func RoutePokemon(api *gin.RouterGroup) {
	api.POST("/pokemons", createPokemon)
	api.GET("/pokemons", getPokemons)
	api.GET("/pokemons/:uuid", getPokemons)
	api.PUT("/pokemons", updatePokemon)
	api.DELETE("/pokemons/:uuid", deletePokemon)
}

// create
func createPokemon(context *gin.Context) {
	var reqPokemon model.CreatePokemonReq

	err := context.BindJSON(&reqPokemon)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var types []model.TypePokemon
	model.DB.Where("code IN (?)", reqPokemon.Types).Find(&types)

	pokemon := &model.Pokemon{
		Number: reqPokemon.Number,
		Name: reqPokemon.Name,
		Types: types,
	}

	resultId, err := repository.CreatePokemon(model.DB, pokemon)
	if err != nil {
		log.Println(err.Error())
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "OK",
		"data":    resultId,
	})
}

// read
func getPokemons(context *gin.Context) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "0"))
	uuid := context.Param("uuid")

	if uuid != "" {
		pokemon := repository.GetPokemon(model.DB, uuid)
		var pokemonRes interface{}
		var typePoke []string

		for _, typePokemon := range pokemon.Types {
			typePoke = append(typePoke, typePokemon.Name)
		}

		pokemonRes = map[string]interface{}{
			"uuid": pokemon.UUID,
			"name": pokemon.Name,
			"number": pokemon.Number,
			"types": typePoke,
		}

		context.JSON(http.StatusOK, gin.H{
			"data": pokemonRes,
		})
	} else {
		pokemons := repository.ListPokemons(model.DB, limit)
		var pokemonsRes []interface{}

		for _, pokemon := range pokemons {
			var typePoke []string
			for _, typePokemon := range pokemon.Types {
				typePoke = append(typePoke, typePokemon.Name)
			}
			pokemonsRes = append(pokemonsRes, map[string]interface{}{
				"uuid": pokemon.UUID,
				"name": pokemon.Name,
				"number": pokemon.Number,
				"types": typePoke,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"pokemons": pokemonsRes,
			},
		})
	}
}

// update
func updatePokemon(context *gin.Context) {
	var reqPokemon model.CreatePokemonReq
	context.BindJSON(&reqPokemon)

	var types []model.TypePokemon
	model.DB.Where("code IN (?)", reqPokemon.Types).Find(&types)

	pokemon := &model.Pokemon{
		UUID: reqPokemon.UUID,
		Number: reqPokemon.Number,
		Name: reqPokemon.Name,
		Types: types,
	}

	resultId, err := repository.UpdatePokemon(model.DB, pokemon)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"code": http.StatusConflict,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "OK",
			"uuid": resultId,
		})
	}
}

// delete
func deletePokemon(context *gin.Context) {
	uuid := context.Param("uuid")
	pokemon := repository.GetPokemon(model.DB, uuid)

	err := repository.DeletePokemon(model.DB, &pokemon)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "OK",
		})
	}
}