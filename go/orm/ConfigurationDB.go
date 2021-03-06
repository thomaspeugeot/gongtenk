// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/thomaspeugeot/gongtenk/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Configuration_sql sql.NullBool
var dummy_Configuration_time time.Duration
var dummy_Configuration_sort sort.Float64Slice

// ConfigurationAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model configurationAPI
type ConfigurationAPI struct {
	gorm.Model

	models.Configuration

	// encoding of pointers
	ConfigurationPointersEnconding
}

// ConfigurationPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ConfigurationPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// ConfigurationDB describes a configuration in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model configurationDB
type ConfigurationDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field configurationDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field configurationDB.NumberOfCitiesToDisplay {{BasicKind}} (to be completed)
	NumberOfCitiesToDisplay_Data sql.NullInt64
	// encoding of pointers
	ConfigurationPointersEnconding
}

// ConfigurationDBs arrays configurationDBs
// swagger:response configurationDBsResponse
type ConfigurationDBs []ConfigurationDB

// ConfigurationDBResponse provides response
// swagger:response configurationDBResponse
type ConfigurationDBResponse struct {
	ConfigurationDB
}

// ConfigurationWOP is a Configuration without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ConfigurationWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	NumberOfCitiesToDisplay int `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Configuration_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"NumberOfCitiesToDisplay",
}

type BackRepoConfigurationStruct struct {
	// stores ConfigurationDB according to their gorm ID
	Map_ConfigurationDBID_ConfigurationDB *map[uint]*ConfigurationDB

	// stores ConfigurationDB ID according to Configuration address
	Map_ConfigurationPtr_ConfigurationDBID *map[*models.Configuration]uint

	// stores Configuration according to their gorm ID
	Map_ConfigurationDBID_ConfigurationPtr *map[uint]*models.Configuration

	db *gorm.DB
}

func (backRepoConfiguration *BackRepoConfigurationStruct) GetDB() *gorm.DB {
	return backRepoConfiguration.db
}

// GetConfigurationDBFromConfigurationPtr is a handy function to access the back repo instance from the stage instance
func (backRepoConfiguration *BackRepoConfigurationStruct) GetConfigurationDBFromConfigurationPtr(configuration *models.Configuration) (configurationDB *ConfigurationDB) {
	id := (*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]
	configurationDB = (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[id]
	return
}

// BackRepoConfiguration.Init set up the BackRepo of the Configuration
func (backRepoConfiguration *BackRepoConfigurationStruct) Init(db *gorm.DB) (Error error) {

	if backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr != nil {
		err := errors.New("In Init, backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr should be nil")
		return err
	}

	if backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB != nil {
		err := errors.New("In Init, backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB should be nil")
		return err
	}

	if backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID != nil {
		err := errors.New("In Init, backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Configuration, 0)
	backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr = &tmp

	tmpDB := make(map[uint]*ConfigurationDB, 0)
	backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB = &tmpDB

	tmpID := make(map[*models.Configuration]uint, 0)
	backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID = &tmpID

	backRepoConfiguration.db = db
	return
}

// BackRepoConfiguration.CommitPhaseOne commits all staged instances of Configuration to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoConfiguration *BackRepoConfigurationStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for configuration := range stage.Configurations {
		backRepoConfiguration.CommitPhaseOneInstance(configuration)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, configuration := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr {
		if _, ok := stage.Configurations[configuration]; !ok {
			backRepoConfiguration.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoConfiguration.CommitDeleteInstance commits deletion of Configuration to the BackRepo
func (backRepoConfiguration *BackRepoConfigurationStruct) CommitDeleteInstance(id uint) (Error error) {

	configuration := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[id]

	// configuration is not staged anymore, remove configurationDB
	configurationDB := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[id]
	query := backRepoConfiguration.db.Unscoped().Delete(&configurationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID), configuration)
	delete((*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr), id)
	delete((*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB), id)

	return
}

// BackRepoConfiguration.CommitPhaseOneInstance commits configuration staged instances of Configuration to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoConfiguration *BackRepoConfigurationStruct) CommitPhaseOneInstance(configuration *models.Configuration) (Error error) {

	// check if the configuration is not commited yet
	if _, ok := (*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]; ok {
		return
	}

	// initiate configuration
	var configurationDB ConfigurationDB
	configurationDB.CopyBasicFieldsFromConfiguration(configuration)

	query := backRepoConfiguration.db.Create(&configurationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration] = configurationDB.ID
	(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[configurationDB.ID] = configuration
	(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[configurationDB.ID] = &configurationDB

	return
}

// BackRepoConfiguration.CommitPhaseTwo commits all staged instances of Configuration to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoConfiguration *BackRepoConfigurationStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, configuration := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr {
		backRepoConfiguration.CommitPhaseTwoInstance(backRepo, idx, configuration)
	}

	return
}

// BackRepoConfiguration.CommitPhaseTwoInstance commits {{structname }} of models.Configuration to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoConfiguration *BackRepoConfigurationStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, configuration *models.Configuration) (Error error) {

	// fetch matching configurationDB
	if configurationDB, ok := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[idx]; ok {

		configurationDB.CopyBasicFieldsFromConfiguration(configuration)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoConfiguration.db.Save(&configurationDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Configuration intance %s", configuration.Name))
		return err
	}

	return
}

// BackRepoConfiguration.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoConfiguration *BackRepoConfigurationStruct) CheckoutPhaseOne() (Error error) {

	configurationDBArray := make([]ConfigurationDB, 0)
	query := backRepoConfiguration.db.Find(&configurationDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	configurationInstancesToBeRemovedFromTheStage := make(map[*models.Configuration]struct{})
	for key, value := range models.Stage.Configurations {
		configurationInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, configurationDB := range configurationDBArray {
		backRepoConfiguration.CheckoutPhaseOneInstance(&configurationDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		configuration, ok := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[configurationDB.ID]
		if ok {
			delete(configurationInstancesToBeRemovedFromTheStage, configuration)
		}
	}

	// remove from stage and back repo's 3 maps all configurations that are not in the checkout
	for configuration := range configurationInstancesToBeRemovedFromTheStage {
		configuration.Unstage()

		// remove instance from the back repo 3 maps
		configurationID := (*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]
		delete((*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID), configuration)
		delete((*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB), configurationID)
		delete((*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr), configurationID)
	}

	return
}

// CheckoutPhaseOneInstance takes a configurationDB that has been found in the DB, updates the backRepo and stages the
// models version of the configurationDB
func (backRepoConfiguration *BackRepoConfigurationStruct) CheckoutPhaseOneInstance(configurationDB *ConfigurationDB) (Error error) {

	configuration, ok := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[configurationDB.ID]
	if !ok {
		configuration = new(models.Configuration)

		(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[configurationDB.ID] = configuration
		(*backRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration] = configurationDB.ID

		// append model store with the new element
		configuration.Name = configurationDB.Name_Data.String
		configuration.Stage()
	}
	configurationDB.CopyBasicFieldsToConfiguration(configuration)

	// preserve pointer to configurationDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ConfigurationDBID_ConfigurationDB)[configurationDB hold variable pointers
	configurationDB_Data := *configurationDB
	preservedPtrToConfiguration := &configurationDB_Data
	(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[configurationDB.ID] = preservedPtrToConfiguration

	return
}

// BackRepoConfiguration.CheckoutPhaseTwo Checkouts all staged instances of Configuration to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoConfiguration *BackRepoConfigurationStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, configurationDB := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB {
		backRepoConfiguration.CheckoutPhaseTwoInstance(backRepo, configurationDB)
	}
	return
}

// BackRepoConfiguration.CheckoutPhaseTwoInstance Checkouts staged instances of Configuration to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoConfiguration *BackRepoConfigurationStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, configurationDB *ConfigurationDB) (Error error) {

	configuration := (*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationPtr)[configurationDB.ID]
	_ = configuration // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitConfiguration allows commit of a single configuration (if already staged)
func (backRepo *BackRepoStruct) CommitConfiguration(configuration *models.Configuration) {
	backRepo.BackRepoConfiguration.CommitPhaseOneInstance(configuration)
	if id, ok := (*backRepo.BackRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]; ok {
		backRepo.BackRepoConfiguration.CommitPhaseTwoInstance(backRepo, id, configuration)
	}
}

// CommitConfiguration allows checkout of a single configuration (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutConfiguration(configuration *models.Configuration) {
	// check if the configuration is staged
	if _, ok := (*backRepo.BackRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]; ok {

		if id, ok := (*backRepo.BackRepoConfiguration.Map_ConfigurationPtr_ConfigurationDBID)[configuration]; ok {
			var configurationDB ConfigurationDB
			configurationDB.ID = id

			if err := backRepo.BackRepoConfiguration.db.First(&configurationDB, id).Error; err != nil {
				log.Panicln("CheckoutConfiguration : Problem with getting object with id:", id)
			}
			backRepo.BackRepoConfiguration.CheckoutPhaseOneInstance(&configurationDB)
			backRepo.BackRepoConfiguration.CheckoutPhaseTwoInstance(backRepo, &configurationDB)
		}
	}
}

// CopyBasicFieldsFromConfiguration
func (configurationDB *ConfigurationDB) CopyBasicFieldsFromConfiguration(configuration *models.Configuration) {
	// insertion point for fields commit

	configurationDB.Name_Data.String = configuration.Name
	configurationDB.Name_Data.Valid = true

	configurationDB.NumberOfCitiesToDisplay_Data.Int64 = int64(configuration.NumberOfCitiesToDisplay)
	configurationDB.NumberOfCitiesToDisplay_Data.Valid = true
}

// CopyBasicFieldsFromConfigurationWOP
func (configurationDB *ConfigurationDB) CopyBasicFieldsFromConfigurationWOP(configuration *ConfigurationWOP) {
	// insertion point for fields commit

	configurationDB.Name_Data.String = configuration.Name
	configurationDB.Name_Data.Valid = true

	configurationDB.NumberOfCitiesToDisplay_Data.Int64 = int64(configuration.NumberOfCitiesToDisplay)
	configurationDB.NumberOfCitiesToDisplay_Data.Valid = true
}

// CopyBasicFieldsToConfiguration
func (configurationDB *ConfigurationDB) CopyBasicFieldsToConfiguration(configuration *models.Configuration) {
	// insertion point for checkout of basic fields (back repo to stage)
	configuration.Name = configurationDB.Name_Data.String
	configuration.NumberOfCitiesToDisplay = int(configurationDB.NumberOfCitiesToDisplay_Data.Int64)
}

// CopyBasicFieldsToConfigurationWOP
func (configurationDB *ConfigurationDB) CopyBasicFieldsToConfigurationWOP(configuration *ConfigurationWOP) {
	configuration.ID = int(configurationDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	configuration.Name = configurationDB.Name_Data.String
	configuration.NumberOfCitiesToDisplay = int(configurationDB.NumberOfCitiesToDisplay_Data.Int64)
}

// Backup generates a json file from a slice of all ConfigurationDB instances in the backrepo
func (backRepoConfiguration *BackRepoConfigurationStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ConfigurationDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ConfigurationDB, 0)
	for _, configurationDB := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB {
		forBackup = append(forBackup, configurationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Configuration ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Configuration file", err.Error())
	}
}

// Backup generates a json file from a slice of all ConfigurationDB instances in the backrepo
func (backRepoConfiguration *BackRepoConfigurationStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ConfigurationDB, 0)
	for _, configurationDB := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB {
		forBackup = append(forBackup, configurationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Configuration")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Configuration_Fields, -1)
	for _, configurationDB := range forBackup {

		var configurationWOP ConfigurationWOP
		configurationDB.CopyBasicFieldsToConfigurationWOP(&configurationWOP)

		row := sh.AddRow()
		row.WriteStruct(&configurationWOP, -1)
	}
}

// RestoreXL from the "Configuration" sheet all ConfigurationDB instances
func (backRepoConfiguration *BackRepoConfigurationStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoConfigurationid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Configuration"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoConfiguration.rowVisitorConfiguration)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoConfiguration *BackRepoConfigurationStruct) rowVisitorConfiguration(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var configurationWOP ConfigurationWOP
		row.ReadStruct(&configurationWOP)

		// add the unmarshalled struct to the stage
		configurationDB := new(ConfigurationDB)
		configurationDB.CopyBasicFieldsFromConfigurationWOP(&configurationWOP)

		configurationDB_ID_atBackupTime := configurationDB.ID
		configurationDB.ID = 0
		query := backRepoConfiguration.db.Create(configurationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[configurationDB.ID] = configurationDB
		BackRepoConfigurationid_atBckpTime_newID[configurationDB_ID_atBackupTime] = configurationDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ConfigurationDB.json" in dirPath that stores an array
// of ConfigurationDB and stores it in the database
// the map BackRepoConfigurationid_atBckpTime_newID is updated accordingly
func (backRepoConfiguration *BackRepoConfigurationStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoConfigurationid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ConfigurationDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Configuration file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ConfigurationDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ConfigurationDBID_ConfigurationDB
	for _, configurationDB := range forRestore {

		configurationDB_ID_atBackupTime := configurationDB.ID
		configurationDB.ID = 0
		query := backRepoConfiguration.db.Create(configurationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB)[configurationDB.ID] = configurationDB
		BackRepoConfigurationid_atBckpTime_newID[configurationDB_ID_atBackupTime] = configurationDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Configuration file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Configuration>id_atBckpTime_newID
// to compute new index
func (backRepoConfiguration *BackRepoConfigurationStruct) RestorePhaseTwo() {

	for _, configurationDB := range *backRepoConfiguration.Map_ConfigurationDBID_ConfigurationDB {

		// next line of code is to avert unused variable compilation error
		_ = configurationDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoConfiguration.db.Model(configurationDB).Updates(*configurationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoConfigurationid_atBckpTime_newID map[uint]uint
