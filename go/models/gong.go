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
	Configurations           map[*Configuration]struct{}
	Configurations_mapString map[string]*Configuration

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
	CommitConfiguration(configuration *Configuration)
	CheckoutConfiguration(configuration *Configuration)
	GetLastCommitNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation
	Configurations:           make(map[*Configuration]struct{}),
	Configurations_mapString: make(map[string]*Configuration),

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
func (stage *StageStruct) getConfigurationOrderedStructWithNameField() []*Configuration {
	// have alphabetical order generation
	configurationOrdered := []*Configuration{}
	for configuration := range stage.Configurations {
		configurationOrdered = append(configurationOrdered, configuration)
	}
	sort.Slice(configurationOrdered[:], func(i, j int) bool {
		return configurationOrdered[i].Name < configurationOrdered[j].Name
	})
	return configurationOrdered
}

// Stage puts configuration to the model stage
func (configuration *Configuration) Stage() *Configuration {
	Stage.Configurations[configuration] = __member
	Stage.Configurations_mapString[configuration.Name] = configuration

	return configuration
}

// Unstage removes configuration off the model stage
func (configuration *Configuration) Unstage() *Configuration {
	delete(Stage.Configurations, configuration)
	delete(Stage.Configurations_mapString, configuration.Name)
	return configuration
}

// commit configuration to the back repo (if it is already staged)
func (configuration *Configuration) Commit() *Configuration {
	if _, ok := Stage.Configurations[configuration]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitConfiguration(configuration)
		}
	}
	return configuration
}

// Checkout configuration to the back repo (if it is already staged)
func (configuration *Configuration) Checkout() *Configuration {
	if _, ok := Stage.Configurations[configuration]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutConfiguration(configuration)
		}
	}
	return configuration
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of configuration to the model stage
func (configuration *Configuration) StageCopy() *Configuration {
	_configuration := new(Configuration)
	*_configuration = *configuration
	_configuration.Stage()
	return _configuration
}

// StageAndCommit appends configuration to the model stage and commit to the orm repo
func (configuration *Configuration) StageAndCommit() *Configuration {
	configuration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMConfiguration(configuration)
	}
	return configuration
}

// DeleteStageAndCommit appends configuration to the model stage and commit to the orm repo
func (configuration *Configuration) DeleteStageAndCommit() *Configuration {
	configuration.Unstage()
	DeleteORMConfiguration(configuration)
	return configuration
}

// StageCopyAndCommit appends a copy of configuration to the model stage and commit to the orm repo
func (configuration *Configuration) StageCopyAndCommit() *Configuration {
	_configuration := new(Configuration)
	*_configuration = *configuration
	_configuration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMConfiguration(configuration)
	}
	return _configuration
}

// CreateORMConfiguration enables dynamic staging of a Configuration instance
func CreateORMConfiguration(configuration *Configuration) {
	configuration.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMConfiguration(configuration)
	}
}

// DeleteORMConfiguration enables dynamic staging of a Configuration instance
func DeleteORMConfiguration(configuration *Configuration) {
	configuration.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMConfiguration(configuration)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMConfiguration(Configuration *Configuration)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMConfiguration(Configuration *Configuration)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Configurations = make(map[*Configuration]struct{})
	stage.Configurations_mapString = make(map[string]*Configuration)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Configurations = nil
	stage.Configurations_mapString = nil

}
