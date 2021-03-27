package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rockavoldy/pokemon-rest/model"
	"github.com/rockavoldy/pokemon-rest/repository"
	"net/http"
	"strconv"
)

func RouteType(api *gin.RouterGroup) {
	api.POST("/types", createType)
	api.GET("/types", getTypes)
	api.GET("/types/:uuid", getTypes)
	api.PUT("/types", updateType)
	api.DELETE("/types/:uuid_type", deleteType)
}

// create
func createType(context *gin.Context) {
	var reqTypeData model.CreateTypePokemonReq
	context.BindJSON(&reqTypeData)

	typePokemon := &model.TypePokemon{
		Code: reqTypeData.Code,
		Name: reqTypeData.Name,
	}

	resultId, err := repository.CreateType(model.DB, typePokemon)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"code": http.StatusConflict,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusCreated, gin.H{
			"code": http.StatusCreated,
			"message": "OK",
			"uuid":    resultId,
		})
	}
}

// read
func getTypes(context *gin.Context) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "0"))
	uuid := context.Param("uuid")

	if uuid != "" {
		var typePokemon model.TypePokemon
		var typePoke interface{}

		typePokemon = repository.GetType(model.DB, uuid)

		typePoke = gin.H{
			"uuid": typePokemon.UUID,
			"name": typePokemon.Name,
			"code": typePokemon.Code,
		}

		context.JSON(http.StatusOK, gin.H{
			"data": typePoke,
		})
	} else {
		var typePokemons []model.TypePokemon
		var typePokes []interface{}

		typePokemons = repository.ListTypes(model.DB, limit)

		for _, typePoke := range typePokemons {
			typePokes = append(typePokes, gin.H{
				"uuid": typePoke.UUID,
				"name": typePoke.Name,
				"code": typePoke.Code,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"types": typePokes,
			},
		})
	}
}

// update
func updateType(context *gin.Context) {
	var typePokemon model.TypePokemon
	context.BindJSON(&typePokemon)

	resultId, err := repository.UpdateType(model.DB, &typePokemon)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"code": http.StatusConflict,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "OK",
			"uuid":    resultId,
		})
	}
}

// delete
func deleteType(context *gin.Context) {
	uuidType := context.Param("uuid_type")
	typePokemon := repository.GetType(model.DB, uuidType)

	err := repository.DeleteType(model.DB, &typePokemon)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}