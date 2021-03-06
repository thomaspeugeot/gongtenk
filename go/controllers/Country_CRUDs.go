// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/thomaspeugeot/gongtenk/go/models"
	"github.com/thomaspeugeot/gongtenk/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Country__dummysDeclaration__ models.Country
var __Country_time__dummyDeclaration time.Duration

// An CountryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCountry updateCountry deleteCountry
type CountryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CountryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCountry updateCountry
type CountryInput struct {
	// The Country to submit or modify
	// in: body
	Country *orm.CountryAPI
}

// GetCountrys
//
// swagger:route GET /countrys countrys getCountrys
//
// Get all countrys
//
// Responses:
//    default: genericError
//        200: countryDBsResponse
func GetCountrys(c *gin.Context) {
	db := orm.BackRepo.BackRepoCountry.GetDB()

	// source slice
	var countryDBs []orm.CountryDB
	query := db.Find(&countryDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	countryAPIs := make([]orm.CountryAPI, 0)

	// for each country, update fields from the database nullable fields
	for idx := range countryDBs {
		countryDB := &countryDBs[idx]
		_ = countryDB
		var countryAPI orm.CountryAPI

		// insertion point for updating fields
		countryAPI.ID = countryDB.ID
		countryDB.CopyBasicFieldsToCountry(&countryAPI.Country)
		countryAPI.CountryPointersEnconding = countryDB.CountryPointersEnconding
		countryAPIs = append(countryAPIs, countryAPI)
	}

	c.JSON(http.StatusOK, countryAPIs)
}

// PostCountry
//
// swagger:route POST /countrys countrys postCountry
//
// Creates a country
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: countryDBResponse
func PostCountry(c *gin.Context) {
	db := orm.BackRepo.BackRepoCountry.GetDB()

	// Validate input
	var input orm.CountryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create country
	countryDB := orm.CountryDB{}
	countryDB.CountryPointersEnconding = input.CountryPointersEnconding
	countryDB.CopyBasicFieldsFromCountry(&input.Country)

	query := db.Create(&countryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, countryDB)
}

// GetCountry
//
// swagger:route GET /countrys/{ID} countrys getCountry
//
// Gets the details for a country.
//
// Responses:
//    default: genericError
//        200: countryDBResponse
func GetCountry(c *gin.Context) {
	db := orm.BackRepo.BackRepoCountry.GetDB()

	// Get countryDB in DB
	var countryDB orm.CountryDB
	if err := db.First(&countryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var countryAPI orm.CountryAPI
	countryAPI.ID = countryDB.ID
	countryAPI.CountryPointersEnconding = countryDB.CountryPointersEnconding
	countryDB.CopyBasicFieldsToCountry(&countryAPI.Country)

	c.JSON(http.StatusOK, countryAPI)
}

// UpdateCountry
//
// swagger:route PATCH /countrys/{ID} countrys updateCountry
//
// Update a country
//
// Responses:
//    default: genericError
//        200: countryDBResponse
func UpdateCountry(c *gin.Context) {
	db := orm.BackRepo.BackRepoCountry.GetDB()

	// Get model if exist
	var countryDB orm.CountryDB

	// fetch the country
	query := db.First(&countryDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.CountryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	countryDB.CopyBasicFieldsFromCountry(&input.Country)
	countryDB.CountryPointersEnconding = input.CountryPointersEnconding

	query = db.Model(&countryDB).Updates(countryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the countryDB
	c.JSON(http.StatusOK, countryDB)
}

// DeleteCountry
//
// swagger:route DELETE /countrys/{ID} countrys deleteCountry
//
// Delete a country
//
// Responses:
//    default: genericError
func DeleteCountry(c *gin.Context) {
	db := orm.BackRepo.BackRepoCountry.GetDB()

	// Get model if exist
	var countryDB orm.CountryDB
	if err := db.First(&countryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&countryDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
