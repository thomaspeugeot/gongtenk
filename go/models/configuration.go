package models

// Configuration is the singloton
type Configuration struct {
	Name string

	// user request
	NumberOfCitiesToDisplay int

	// is the same as user request when display has been updated
	NumberOfCitiesToDisplay_real int
}

var ConfigurationSingloton = (&Configuration{
	Name:                         "Gong Tenk Configuration",
	NumberOfCitiesToDisplay:      200,
	NumberOfCitiesToDisplay_real: 200,
}).Stage()
