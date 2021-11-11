package visuals

import (
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var AircraftLayerGroup = (&gongleaflet_models.LayerGroup{
	Name: "Aircraft",
}).StageCopy()
