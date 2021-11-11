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

	"github.com/fullstack-lang/gongleaflet/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_LayerGroupUse_sql sql.NullBool
var dummy_LayerGroupUse_time time.Duration
var dummy_LayerGroupUse_sort sort.Float64Slice

// LayerGroupUseAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model layergroupuseAPI
type LayerGroupUseAPI struct {
	gorm.Model

	models.LayerGroupUse

	// encoding of pointers
	LayerGroupUsePointersEnconding
}

// LayerGroupUsePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LayerGroupUsePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field LayerGroup is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	LayerGroupID sql.NullInt64

	// Implementation of a reverse ID for field MapOptions{}.LayerGroupUses []*LayerGroupUse
	MapOptions_LayerGroupUsesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	MapOptions_LayerGroupUsesDBID_Index sql.NullInt64
}

// LayerGroupUseDB describes a layergroupuse in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model layergroupuseDB
type LayerGroupUseDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field layergroupuseDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field layergroupuseDB.Display bool (to be completed)
	// provide the sql storage for the boolan
	Display_Data sql.NullBool
	// encoding of pointers
	LayerGroupUsePointersEnconding
}

// LayerGroupUseDBs arrays layergroupuseDBs
// swagger:response layergroupuseDBsResponse
type LayerGroupUseDBs []LayerGroupUseDB

// LayerGroupUseDBResponse provides response
// swagger:response layergroupuseDBResponse
type LayerGroupUseDBResponse struct {
	LayerGroupUseDB
}

// LayerGroupUseWOP is a LayerGroupUse without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LayerGroupUseWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Display bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var LayerGroupUse_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Display",
}

type BackRepoLayerGroupUseStruct struct {
	// stores LayerGroupUseDB according to their gorm ID
	Map_LayerGroupUseDBID_LayerGroupUseDB *map[uint]*LayerGroupUseDB

	// stores LayerGroupUseDB ID according to LayerGroupUse address
	Map_LayerGroupUsePtr_LayerGroupUseDBID *map[*models.LayerGroupUse]uint

	// stores LayerGroupUse according to their gorm ID
	Map_LayerGroupUseDBID_LayerGroupUsePtr *map[uint]*models.LayerGroupUse

	db *gorm.DB
}

func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) GetDB() *gorm.DB {
	return backRepoLayerGroupUse.db
}

// GetLayerGroupUseDBFromLayerGroupUsePtr is a handy function to access the back repo instance from the stage instance
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) GetLayerGroupUseDBFromLayerGroupUsePtr(layergroupuse *models.LayerGroupUse) (layergroupuseDB *LayerGroupUseDB) {
	id := (*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]
	layergroupuseDB = (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[id]
	return
}

// BackRepoLayerGroupUse.Init set up the BackRepo of the LayerGroupUse
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) Init(db *gorm.DB) (Error error) {

	if backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr != nil {
		err := errors.New("In Init, backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr should be nil")
		return err
	}

	if backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB != nil {
		err := errors.New("In Init, backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB should be nil")
		return err
	}

	if backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID != nil {
		err := errors.New("In Init, backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.LayerGroupUse, 0)
	backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr = &tmp

	tmpDB := make(map[uint]*LayerGroupUseDB, 0)
	backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB = &tmpDB

	tmpID := make(map[*models.LayerGroupUse]uint, 0)
	backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID = &tmpID

	backRepoLayerGroupUse.db = db
	return
}

// BackRepoLayerGroupUse.CommitPhaseOne commits all staged instances of LayerGroupUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for layergroupuse := range stage.LayerGroupUses {
		backRepoLayerGroupUse.CommitPhaseOneInstance(layergroupuse)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, layergroupuse := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr {
		if _, ok := stage.LayerGroupUses[layergroupuse]; !ok {
			backRepoLayerGroupUse.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLayerGroupUse.CommitDeleteInstance commits deletion of LayerGroupUse to the BackRepo
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CommitDeleteInstance(id uint) (Error error) {

	layergroupuse := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[id]

	// layergroupuse is not staged anymore, remove layergroupuseDB
	layergroupuseDB := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[id]
	query := backRepoLayerGroupUse.db.Unscoped().Delete(&layergroupuseDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID), layergroupuse)
	delete((*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr), id)
	delete((*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB), id)

	return
}

// BackRepoLayerGroupUse.CommitPhaseOneInstance commits layergroupuse staged instances of LayerGroupUse to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CommitPhaseOneInstance(layergroupuse *models.LayerGroupUse) (Error error) {

	// check if the layergroupuse is not commited yet
	if _, ok := (*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]; ok {
		return
	}

	// initiate layergroupuse
	var layergroupuseDB LayerGroupUseDB
	layergroupuseDB.CopyBasicFieldsFromLayerGroupUse(layergroupuse)

	query := backRepoLayerGroupUse.db.Create(&layergroupuseDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse] = layergroupuseDB.ID
	(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[layergroupuseDB.ID] = layergroupuse
	(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[layergroupuseDB.ID] = &layergroupuseDB

	return
}

// BackRepoLayerGroupUse.CommitPhaseTwo commits all staged instances of LayerGroupUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, layergroupuse := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr {
		backRepoLayerGroupUse.CommitPhaseTwoInstance(backRepo, idx, layergroupuse)
	}

	return
}

// BackRepoLayerGroupUse.CommitPhaseTwoInstance commits {{structname }} of models.LayerGroupUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, layergroupuse *models.LayerGroupUse) (Error error) {

	// fetch matching layergroupuseDB
	if layergroupuseDB, ok := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[idx]; ok {

		layergroupuseDB.CopyBasicFieldsFromLayerGroupUse(layergroupuse)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value layergroupuse.LayerGroup translates to updating the layergroupuse.LayerGroupID
		layergroupuseDB.LayerGroupID.Valid = true // allow for a 0 value (nil association)
		if layergroupuse.LayerGroup != nil {
			if LayerGroupId, ok := (*backRepo.BackRepoLayerGroup.Map_LayerGroupPtr_LayerGroupDBID)[layergroupuse.LayerGroup]; ok {
				layergroupuseDB.LayerGroupID.Int64 = int64(LayerGroupId)
				layergroupuseDB.LayerGroupID.Valid = true
			}
		}

		query := backRepoLayerGroupUse.db.Save(&layergroupuseDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown LayerGroupUse intance %s", layergroupuse.Name))
		return err
	}

	return
}

// BackRepoLayerGroupUse.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CheckoutPhaseOne() (Error error) {

	layergroupuseDBArray := make([]LayerGroupUseDB, 0)
	query := backRepoLayerGroupUse.db.Find(&layergroupuseDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	layergroupuseInstancesToBeRemovedFromTheStage := make(map[*models.LayerGroupUse]struct{})
	for key, value := range models.Stage.LayerGroupUses {
		layergroupuseInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, layergroupuseDB := range layergroupuseDBArray {
		backRepoLayerGroupUse.CheckoutPhaseOneInstance(&layergroupuseDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		layergroupuse, ok := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[layergroupuseDB.ID]
		if ok {
			delete(layergroupuseInstancesToBeRemovedFromTheStage, layergroupuse)
		}
	}

	// remove from stage and back repo's 3 maps all layergroupuses that are not in the checkout
	for layergroupuse := range layergroupuseInstancesToBeRemovedFromTheStage {
		layergroupuse.Unstage()

		// remove instance from the back repo 3 maps
		layergroupuseID := (*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]
		delete((*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID), layergroupuse)
		delete((*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB), layergroupuseID)
		delete((*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr), layergroupuseID)
	}

	return
}

// CheckoutPhaseOneInstance takes a layergroupuseDB that has been found in the DB, updates the backRepo and stages the
// models version of the layergroupuseDB
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CheckoutPhaseOneInstance(layergroupuseDB *LayerGroupUseDB) (Error error) {

	layergroupuse, ok := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[layergroupuseDB.ID]
	if !ok {
		layergroupuse = new(models.LayerGroupUse)

		(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[layergroupuseDB.ID] = layergroupuse
		(*backRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse] = layergroupuseDB.ID

		// append model store with the new element
		layergroupuse.Name = layergroupuseDB.Name_Data.String
		layergroupuse.Stage()
	}
	layergroupuseDB.CopyBasicFieldsToLayerGroupUse(layergroupuse)

	// preserve pointer to layergroupuseDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LayerGroupUseDBID_LayerGroupUseDB)[layergroupuseDB hold variable pointers
	layergroupuseDB_Data := *layergroupuseDB
	preservedPtrToLayerGroupUse := &layergroupuseDB_Data
	(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[layergroupuseDB.ID] = preservedPtrToLayerGroupUse

	return
}

// BackRepoLayerGroupUse.CheckoutPhaseTwo Checkouts all staged instances of LayerGroupUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, layergroupuseDB := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB {
		backRepoLayerGroupUse.CheckoutPhaseTwoInstance(backRepo, layergroupuseDB)
	}
	return
}

// BackRepoLayerGroupUse.CheckoutPhaseTwoInstance Checkouts staged instances of LayerGroupUse to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, layergroupuseDB *LayerGroupUseDB) (Error error) {

	layergroupuse := (*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUsePtr)[layergroupuseDB.ID]
	_ = layergroupuse // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// LayerGroup field
	if layergroupuseDB.LayerGroupID.Int64 != 0 {
		layergroupuse.LayerGroup = (*backRepo.BackRepoLayerGroup.Map_LayerGroupDBID_LayerGroupPtr)[uint(layergroupuseDB.LayerGroupID.Int64)]
	}
	return
}

// CommitLayerGroupUse allows commit of a single layergroupuse (if already staged)
func (backRepo *BackRepoStruct) CommitLayerGroupUse(layergroupuse *models.LayerGroupUse) {
	backRepo.BackRepoLayerGroupUse.CommitPhaseOneInstance(layergroupuse)
	if id, ok := (*backRepo.BackRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]; ok {
		backRepo.BackRepoLayerGroupUse.CommitPhaseTwoInstance(backRepo, id, layergroupuse)
	}
}

// CommitLayerGroupUse allows checkout of a single layergroupuse (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLayerGroupUse(layergroupuse *models.LayerGroupUse) {
	// check if the layergroupuse is staged
	if _, ok := (*backRepo.BackRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]; ok {

		if id, ok := (*backRepo.BackRepoLayerGroupUse.Map_LayerGroupUsePtr_LayerGroupUseDBID)[layergroupuse]; ok {
			var layergroupuseDB LayerGroupUseDB
			layergroupuseDB.ID = id

			if err := backRepo.BackRepoLayerGroupUse.db.First(&layergroupuseDB, id).Error; err != nil {
				log.Panicln("CheckoutLayerGroupUse : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLayerGroupUse.CheckoutPhaseOneInstance(&layergroupuseDB)
			backRepo.BackRepoLayerGroupUse.CheckoutPhaseTwoInstance(backRepo, &layergroupuseDB)
		}
	}
}

// CopyBasicFieldsFromLayerGroupUse
func (layergroupuseDB *LayerGroupUseDB) CopyBasicFieldsFromLayerGroupUse(layergroupuse *models.LayerGroupUse) {
	// insertion point for fields commit

	layergroupuseDB.Name_Data.String = layergroupuse.Name
	layergroupuseDB.Name_Data.Valid = true

	layergroupuseDB.Display_Data.Bool = layergroupuse.Display
	layergroupuseDB.Display_Data.Valid = true
}

// CopyBasicFieldsFromLayerGroupUseWOP
func (layergroupuseDB *LayerGroupUseDB) CopyBasicFieldsFromLayerGroupUseWOP(layergroupuse *LayerGroupUseWOP) {
	// insertion point for fields commit

	layergroupuseDB.Name_Data.String = layergroupuse.Name
	layergroupuseDB.Name_Data.Valid = true

	layergroupuseDB.Display_Data.Bool = layergroupuse.Display
	layergroupuseDB.Display_Data.Valid = true
}

// CopyBasicFieldsToLayerGroupUse
func (layergroupuseDB *LayerGroupUseDB) CopyBasicFieldsToLayerGroupUse(layergroupuse *models.LayerGroupUse) {
	// insertion point for checkout of basic fields (back repo to stage)
	layergroupuse.Name = layergroupuseDB.Name_Data.String
	layergroupuse.Display = layergroupuseDB.Display_Data.Bool
}

// CopyBasicFieldsToLayerGroupUseWOP
func (layergroupuseDB *LayerGroupUseDB) CopyBasicFieldsToLayerGroupUseWOP(layergroupuse *LayerGroupUseWOP) {
	layergroupuse.ID = int(layergroupuseDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	layergroupuse.Name = layergroupuseDB.Name_Data.String
	layergroupuse.Display = layergroupuseDB.Display_Data.Bool
}

// Backup generates a json file from a slice of all LayerGroupUseDB instances in the backrepo
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LayerGroupUseDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LayerGroupUseDB, 0)
	for _, layergroupuseDB := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB {
		forBackup = append(forBackup, layergroupuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json LayerGroupUse ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json LayerGroupUse file", err.Error())
	}
}

// Backup generates a json file from a slice of all LayerGroupUseDB instances in the backrepo
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LayerGroupUseDB, 0)
	for _, layergroupuseDB := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB {
		forBackup = append(forBackup, layergroupuseDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("LayerGroupUse")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&LayerGroupUse_Fields, -1)
	for _, layergroupuseDB := range forBackup {

		var layergroupuseWOP LayerGroupUseWOP
		layergroupuseDB.CopyBasicFieldsToLayerGroupUseWOP(&layergroupuseWOP)

		row := sh.AddRow()
		row.WriteStruct(&layergroupuseWOP, -1)
	}
}

// RestoreXL from the "LayerGroupUse" sheet all LayerGroupUseDB instances
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLayerGroupUseid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["LayerGroupUse"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLayerGroupUse.rowVisitorLayerGroupUse)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) rowVisitorLayerGroupUse(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var layergroupuseWOP LayerGroupUseWOP
		row.ReadStruct(&layergroupuseWOP)

		// add the unmarshalled struct to the stage
		layergroupuseDB := new(LayerGroupUseDB)
		layergroupuseDB.CopyBasicFieldsFromLayerGroupUseWOP(&layergroupuseWOP)

		layergroupuseDB_ID_atBackupTime := layergroupuseDB.ID
		layergroupuseDB.ID = 0
		query := backRepoLayerGroupUse.db.Create(layergroupuseDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[layergroupuseDB.ID] = layergroupuseDB
		BackRepoLayerGroupUseid_atBckpTime_newID[layergroupuseDB_ID_atBackupTime] = layergroupuseDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LayerGroupUseDB.json" in dirPath that stores an array
// of LayerGroupUseDB and stores it in the database
// the map BackRepoLayerGroupUseid_atBckpTime_newID is updated accordingly
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLayerGroupUseid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LayerGroupUseDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json LayerGroupUse file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LayerGroupUseDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LayerGroupUseDBID_LayerGroupUseDB
	for _, layergroupuseDB := range forRestore {

		layergroupuseDB_ID_atBackupTime := layergroupuseDB.ID
		layergroupuseDB.ID = 0
		query := backRepoLayerGroupUse.db.Create(layergroupuseDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB)[layergroupuseDB.ID] = layergroupuseDB
		BackRepoLayerGroupUseid_atBckpTime_newID[layergroupuseDB_ID_atBackupTime] = layergroupuseDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json LayerGroupUse file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<LayerGroupUse>id_atBckpTime_newID
// to compute new index
func (backRepoLayerGroupUse *BackRepoLayerGroupUseStruct) RestorePhaseTwo() {

	for _, layergroupuseDB := range *backRepoLayerGroupUse.Map_LayerGroupUseDBID_LayerGroupUseDB {

		// next line of code is to avert unused variable compilation error
		_ = layergroupuseDB

		// insertion point for reindexing pointers encoding
		// reindexing LayerGroup field
		if layergroupuseDB.LayerGroupID.Int64 != 0 {
			layergroupuseDB.LayerGroupID.Int64 = int64(BackRepoLayerGroupid_atBckpTime_newID[uint(layergroupuseDB.LayerGroupID.Int64)])
			layergroupuseDB.LayerGroupID.Valid = true
		}

		// This reindex layergroupuse.LayerGroupUses
		if layergroupuseDB.MapOptions_LayerGroupUsesDBID.Int64 != 0 {
			layergroupuseDB.MapOptions_LayerGroupUsesDBID.Int64 =
				int64(BackRepoMapOptionsid_atBckpTime_newID[uint(layergroupuseDB.MapOptions_LayerGroupUsesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLayerGroupUse.db.Model(layergroupuseDB).Updates(*layergroupuseDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLayerGroupUseid_atBckpTime_newID map[uint]uint
