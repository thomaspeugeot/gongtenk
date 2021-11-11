package visuals

import (
	"gongtenk/go/icons"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var CityDummy = (&gongleaflet_models.VisualTrack{
	Name:       "Dummy",
	Lat:        18.5,
	Lng:        -72,
	Level:      100,
	Speed:      300,
	LayerGroup: IndividualLayerGroup,
	DivIcon:    icons.Dot_25,
}).StageCopy()
