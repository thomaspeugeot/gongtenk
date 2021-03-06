// generated by ModelGongFileTemplate
package models

import "sort"

// swagger:ignore
type __void struct{}

// needed for creating set of instances in the stage
var __member __void

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances
	XLCells           map[*XLCell]struct{}
	XLCells_mapString map[string]*XLCell

	XLFiles           map[*XLFile]struct{}
	XLFiles_mapString map[string]*XLFile

	XLRows           map[*XLRow]struct{}
	XLRows_mapString map[string]*XLRow

	XLSheets           map[*XLSheet]struct{}
	XLSheets_mapString map[string]*XLSheet

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback OnInitCommitInterface
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitXLCell(xlcell *XLCell)
	CheckoutXLCell(xlcell *XLCell)
	CommitXLFile(xlfile *XLFile)
	CheckoutXLFile(xlfile *XLFile)
	CommitXLRow(xlrow *XLRow)
	CheckoutXLRow(xlrow *XLRow)
	CommitXLSheet(xlsheet *XLSheet)
	CheckoutXLSheet(xlsheet *XLSheet)
	GetLastCommitNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation
	XLCells:           make(map[*XLCell]struct{}),
	XLCells_mapString: make(map[string]*XLCell),

	XLFiles:           make(map[*XLFile]struct{}),
	XLFiles_mapString: make(map[string]*XLFile),

	XLRows:           make(map[*XLRow]struct{}),
	XLRows_mapString: make(map[string]*XLRow),

	XLSheets:           make(map[*XLSheet]struct{}),
	XLSheets_mapString: make(map[string]*XLSheet),

	// end of insertion point
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
func (stage *StageStruct) getXLCellOrderedStructWithNameField() []*XLCell {
	// have alphabetical order generation
	xlcellOrdered := []*XLCell{}
	for xlcell := range stage.XLCells {
		xlcellOrdered = append(xlcellOrdered, xlcell)
	}
	sort.Slice(xlcellOrdered[:], func(i, j int) bool {
		return xlcellOrdered[i].Name < xlcellOrdered[j].Name
	})
	return xlcellOrdered
}

// Stage puts xlcell to the model stage
func (xlcell *XLCell) Stage() *XLCell {
	Stage.XLCells[xlcell] = __member
	Stage.XLCells_mapString[xlcell.Name] = xlcell

	return xlcell
}

// Unstage removes xlcell off the model stage
func (xlcell *XLCell) Unstage() *XLCell {
	delete(Stage.XLCells, xlcell)
	delete(Stage.XLCells_mapString, xlcell.Name)
	return xlcell
}

// commit xlcell to the back repo (if it is already staged)
func (xlcell *XLCell) Commit() *XLCell {
	if _, ok := Stage.XLCells[xlcell]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitXLCell(xlcell)
		}
	}
	return xlcell
}

// Checkout xlcell to the back repo (if it is already staged)
func (xlcell *XLCell) Checkout() *XLCell {
	if _, ok := Stage.XLCells[xlcell]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutXLCell(xlcell)
		}
	}
	return xlcell
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of xlcell to the model stage
func (xlcell *XLCell) StageCopy() *XLCell {
	_xlcell := new(XLCell)
	*_xlcell = *xlcell
	_xlcell.Stage()
	return _xlcell
}

// StageAndCommit appends xlcell to the model stage and commit to the orm repo
func (xlcell *XLCell) StageAndCommit() *XLCell {
	xlcell.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLCell(xlcell)
	}
	return xlcell
}

// DeleteStageAndCommit appends xlcell to the model stage and commit to the orm repo
func (xlcell *XLCell) DeleteStageAndCommit() *XLCell {
	xlcell.Unstage()
	DeleteORMXLCell(xlcell)
	return xlcell
}

// StageCopyAndCommit appends a copy of xlcell to the model stage and commit to the orm repo
func (xlcell *XLCell) StageCopyAndCommit() *XLCell {
	_xlcell := new(XLCell)
	*_xlcell = *xlcell
	_xlcell.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLCell(xlcell)
	}
	return _xlcell
}

// CreateORMXLCell enables dynamic staging of a XLCell instance
func CreateORMXLCell(xlcell *XLCell) {
	xlcell.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLCell(xlcell)
	}
}

// DeleteORMXLCell enables dynamic staging of a XLCell instance
func DeleteORMXLCell(xlcell *XLCell) {
	xlcell.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMXLCell(xlcell)
	}
}

func (stage *StageStruct) getXLFileOrderedStructWithNameField() []*XLFile {
	// have alphabetical order generation
	xlfileOrdered := []*XLFile{}
	for xlfile := range stage.XLFiles {
		xlfileOrdered = append(xlfileOrdered, xlfile)
	}
	sort.Slice(xlfileOrdered[:], func(i, j int) bool {
		return xlfileOrdered[i].Name < xlfileOrdered[j].Name
	})
	return xlfileOrdered
}

// Stage puts xlfile to the model stage
func (xlfile *XLFile) Stage() *XLFile {
	Stage.XLFiles[xlfile] = __member
	Stage.XLFiles_mapString[xlfile.Name] = xlfile

	return xlfile
}

// Unstage removes xlfile off the model stage
func (xlfile *XLFile) Unstage() *XLFile {
	delete(Stage.XLFiles, xlfile)
	delete(Stage.XLFiles_mapString, xlfile.Name)
	return xlfile
}

// commit xlfile to the back repo (if it is already staged)
func (xlfile *XLFile) Commit() *XLFile {
	if _, ok := Stage.XLFiles[xlfile]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitXLFile(xlfile)
		}
	}
	return xlfile
}

// Checkout xlfile to the back repo (if it is already staged)
func (xlfile *XLFile) Checkout() *XLFile {
	if _, ok := Stage.XLFiles[xlfile]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutXLFile(xlfile)
		}
	}
	return xlfile
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of xlfile to the model stage
func (xlfile *XLFile) StageCopy() *XLFile {
	_xlfile := new(XLFile)
	*_xlfile = *xlfile
	_xlfile.Stage()
	return _xlfile
}

// StageAndCommit appends xlfile to the model stage and commit to the orm repo
func (xlfile *XLFile) StageAndCommit() *XLFile {
	xlfile.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLFile(xlfile)
	}
	return xlfile
}

// DeleteStageAndCommit appends xlfile to the model stage and commit to the orm repo
func (xlfile *XLFile) DeleteStageAndCommit() *XLFile {
	xlfile.Unstage()
	DeleteORMXLFile(xlfile)
	return xlfile
}

// StageCopyAndCommit appends a copy of xlfile to the model stage and commit to the orm repo
func (xlfile *XLFile) StageCopyAndCommit() *XLFile {
	_xlfile := new(XLFile)
	*_xlfile = *xlfile
	_xlfile.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLFile(xlfile)
	}
	return _xlfile
}

// CreateORMXLFile enables dynamic staging of a XLFile instance
func CreateORMXLFile(xlfile *XLFile) {
	xlfile.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLFile(xlfile)
	}
}

// DeleteORMXLFile enables dynamic staging of a XLFile instance
func DeleteORMXLFile(xlfile *XLFile) {
	xlfile.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMXLFile(xlfile)
	}
}

func (stage *StageStruct) getXLRowOrderedStructWithNameField() []*XLRow {
	// have alphabetical order generation
	xlrowOrdered := []*XLRow{}
	for xlrow := range stage.XLRows {
		xlrowOrdered = append(xlrowOrdered, xlrow)
	}
	sort.Slice(xlrowOrdered[:], func(i, j int) bool {
		return xlrowOrdered[i].Name < xlrowOrdered[j].Name
	})
	return xlrowOrdered
}

// Stage puts xlrow to the model stage
func (xlrow *XLRow) Stage() *XLRow {
	Stage.XLRows[xlrow] = __member
	Stage.XLRows_mapString[xlrow.Name] = xlrow

	return xlrow
}

// Unstage removes xlrow off the model stage
func (xlrow *XLRow) Unstage() *XLRow {
	delete(Stage.XLRows, xlrow)
	delete(Stage.XLRows_mapString, xlrow.Name)
	return xlrow
}

// commit xlrow to the back repo (if it is already staged)
func (xlrow *XLRow) Commit() *XLRow {
	if _, ok := Stage.XLRows[xlrow]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitXLRow(xlrow)
		}
	}
	return xlrow
}

// Checkout xlrow to the back repo (if it is already staged)
func (xlrow *XLRow) Checkout() *XLRow {
	if _, ok := Stage.XLRows[xlrow]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutXLRow(xlrow)
		}
	}
	return xlrow
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of xlrow to the model stage
func (xlrow *XLRow) StageCopy() *XLRow {
	_xlrow := new(XLRow)
	*_xlrow = *xlrow
	_xlrow.Stage()
	return _xlrow
}

// StageAndCommit appends xlrow to the model stage and commit to the orm repo
func (xlrow *XLRow) StageAndCommit() *XLRow {
	xlrow.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLRow(xlrow)
	}
	return xlrow
}

// DeleteStageAndCommit appends xlrow to the model stage and commit to the orm repo
func (xlrow *XLRow) DeleteStageAndCommit() *XLRow {
	xlrow.Unstage()
	DeleteORMXLRow(xlrow)
	return xlrow
}

// StageCopyAndCommit appends a copy of xlrow to the model stage and commit to the orm repo
func (xlrow *XLRow) StageCopyAndCommit() *XLRow {
	_xlrow := new(XLRow)
	*_xlrow = *xlrow
	_xlrow.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLRow(xlrow)
	}
	return _xlrow
}

// CreateORMXLRow enables dynamic staging of a XLRow instance
func CreateORMXLRow(xlrow *XLRow) {
	xlrow.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLRow(xlrow)
	}
}

// DeleteORMXLRow enables dynamic staging of a XLRow instance
func DeleteORMXLRow(xlrow *XLRow) {
	xlrow.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMXLRow(xlrow)
	}
}

func (stage *StageStruct) getXLSheetOrderedStructWithNameField() []*XLSheet {
	// have alphabetical order generation
	xlsheetOrdered := []*XLSheet{}
	for xlsheet := range stage.XLSheets {
		xlsheetOrdered = append(xlsheetOrdered, xlsheet)
	}
	sort.Slice(xlsheetOrdered[:], func(i, j int) bool {
		return xlsheetOrdered[i].Name < xlsheetOrdered[j].Name
	})
	return xlsheetOrdered
}

// Stage puts xlsheet to the model stage
func (xlsheet *XLSheet) Stage() *XLSheet {
	Stage.XLSheets[xlsheet] = __member
	Stage.XLSheets_mapString[xlsheet.Name] = xlsheet

	return xlsheet
}

// Unstage removes xlsheet off the model stage
func (xlsheet *XLSheet) Unstage() *XLSheet {
	delete(Stage.XLSheets, xlsheet)
	delete(Stage.XLSheets_mapString, xlsheet.Name)
	return xlsheet
}

// commit xlsheet to the back repo (if it is already staged)
func (xlsheet *XLSheet) Commit() *XLSheet {
	if _, ok := Stage.XLSheets[xlsheet]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitXLSheet(xlsheet)
		}
	}
	return xlsheet
}

// Checkout xlsheet to the back repo (if it is already staged)
func (xlsheet *XLSheet) Checkout() *XLSheet {
	if _, ok := Stage.XLSheets[xlsheet]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutXLSheet(xlsheet)
		}
	}
	return xlsheet
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of xlsheet to the model stage
func (xlsheet *XLSheet) StageCopy() *XLSheet {
	_xlsheet := new(XLSheet)
	*_xlsheet = *xlsheet
	_xlsheet.Stage()
	return _xlsheet
}

// StageAndCommit appends xlsheet to the model stage and commit to the orm repo
func (xlsheet *XLSheet) StageAndCommit() *XLSheet {
	xlsheet.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLSheet(xlsheet)
	}
	return xlsheet
}

// DeleteStageAndCommit appends xlsheet to the model stage and commit to the orm repo
func (xlsheet *XLSheet) DeleteStageAndCommit() *XLSheet {
	xlsheet.Unstage()
	DeleteORMXLSheet(xlsheet)
	return xlsheet
}

// StageCopyAndCommit appends a copy of xlsheet to the model stage and commit to the orm repo
func (xlsheet *XLSheet) StageCopyAndCommit() *XLSheet {
	_xlsheet := new(XLSheet)
	*_xlsheet = *xlsheet
	_xlsheet.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLSheet(xlsheet)
	}
	return _xlsheet
}

// CreateORMXLSheet enables dynamic staging of a XLSheet instance
func CreateORMXLSheet(xlsheet *XLSheet) {
	xlsheet.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMXLSheet(xlsheet)
	}
}

// DeleteORMXLSheet enables dynamic staging of a XLSheet instance
func DeleteORMXLSheet(xlsheet *XLSheet) {
	xlsheet.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMXLSheet(xlsheet)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMXLCell(XLCell *XLCell)
	CreateORMXLFile(XLFile *XLFile)
	CreateORMXLRow(XLRow *XLRow)
	CreateORMXLSheet(XLSheet *XLSheet)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMXLCell(XLCell *XLCell)
	DeleteORMXLFile(XLFile *XLFile)
	DeleteORMXLRow(XLRow *XLRow)
	DeleteORMXLSheet(XLSheet *XLSheet)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.XLCells = make(map[*XLCell]struct{})
	stage.XLCells_mapString = make(map[string]*XLCell)

	stage.XLFiles = make(map[*XLFile]struct{})
	stage.XLFiles_mapString = make(map[string]*XLFile)

	stage.XLRows = make(map[*XLRow]struct{})
	stage.XLRows_mapString = make(map[string]*XLRow)

	stage.XLSheets = make(map[*XLSheet]struct{})
	stage.XLSheets_mapString = make(map[string]*XLSheet)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.XLCells = nil
	stage.XLCells_mapString = nil

	stage.XLFiles = nil
	stage.XLFiles_mapString = nil

	stage.XLRows = nil
	stage.XLRows_mapString = nil

	stage.XLSheets = nil
	stage.XLSheets_mapString = nil

}
