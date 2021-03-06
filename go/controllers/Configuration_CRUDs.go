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
var __Configuration__dummysDeclaration__ models.Configuration
var __Configuration_time__dummyDeclaration time.Duration

// An ConfigurationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getConfiguration updateConfiguration deleteConfiguration
type ConfigurationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ConfigurationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postConfiguration updateConfiguration
type ConfigurationInput struct {
	// The Configuration to submit or modify
	// in: body
	Configuration *orm.ConfigurationAPI
}

// GetConfigurations
//
// swagger:route GET /configurations configurations getConfigurations
//
// Get all configurations
//
// Responses:
//    default: genericError
//        200: configurationDBsResponse
func GetConfigurations(c *gin.Context) {
	db := orm.BackRepo.BackRepoConfiguration.GetDB()

	// source slice
	var configurationDBs []orm.ConfigurationDB
	query := db.Find(&configurationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	configurationAPIs := make([]orm.ConfigurationAPI, 0)

	// for each configuration, update fields from the database nullable fields
	for idx := range configurationDBs {
		configurationDB := &configurationDBs[idx]
		_ = configurationDB
		var configurationAPI orm.ConfigurationAPI

		// insertion point for updating fields
		configurationAPI.ID = configurationDB.ID
		configurationDB.CopyBasicFieldsToConfiguration(&configurationAPI.Configuration)
		configurationAPI.ConfigurationPointersEnconding = configurationDB.ConfigurationPointersEnconding
		configurationAPIs = append(configurationAPIs, configurationAPI)
	}

	c.JSON(http.StatusOK, configurationAPIs)
}

// PostConfiguration
//
// swagger:route POST /configurations configurations postConfiguration
//
// Creates a configuration
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: configurationDBResponse
func PostConfiguration(c *gin.Context) {
	db := orm.BackRepo.BackRepoConfiguration.GetDB()

	// Validate input
	var input orm.ConfigurationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create configuration
	configurationDB := orm.ConfigurationDB{}
	configurationDB.ConfigurationPointersEnconding = input.ConfigurationPointersEnconding
	configurationDB.CopyBasicFieldsFromConfiguration(&input.Configuration)

	query := db.Create(&configurationDB)
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

	c.JSON(http.StatusOK, configurationDB)
}

// GetConfiguration
//
// swagger:route GET /configurations/{ID} configurations getConfiguration
//
// Gets the details for a configuration.
//
// Responses:
//    default: genericError
//        200: configurationDBResponse
func GetConfiguration(c *gin.Context) {
	db := orm.BackRepo.BackRepoConfiguration.GetDB()

	// Get configurationDB in DB
	var configurationDB orm.ConfigurationDB
	if err := db.First(&configurationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var configurationAPI orm.ConfigurationAPI
	configurationAPI.ID = configurationDB.ID
	configurationAPI.ConfigurationPointersEnconding = configurationDB.ConfigurationPointersEnconding
	configurationDB.CopyBasicFieldsToConfiguration(&configurationAPI.Configuration)

	c.JSON(http.StatusOK, configurationAPI)
}

// UpdateConfiguration
//
// swagger:route PATCH /configurations/{ID} configurations updateConfiguration
//
// Update a configuration
//
// Responses:
//    default: genericError
//        200: configurationDBResponse
func UpdateConfiguration(c *gin.Context) {
	db := orm.BackRepo.BackRepoConfiguration.GetDB()

	// Get model if exist
	var configurationDB orm.ConfigurationDB

	// fetch the configuration
	query := db.First(&configurationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ConfigurationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	configurationDB.CopyBasicFieldsFromConfiguration(&input.Configuration)
	configurationDB.ConfigurationPointersEnconding = input.ConfigurationPointersEnconding

	query = db.Model(&configurationDB).Updates(configurationDB)
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

	// return status OK with the marshalling of the the configurationDB
	c.JSON(http.StatusOK, configurationDB)
}

// DeleteConfiguration
//
// swagger:route DELETE /configurations/{ID} configurations deleteConfiguration
//
// Delete a configuration
//
// Responses:
//    default: genericError
func DeleteConfiguration(c *gin.Context) {
	db := orm.BackRepo.BackRepoConfiguration.GetDB()

	// Get model if exist
	var configurationDB orm.ConfigurationDB
	if err := db.First(&configurationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&configurationDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
