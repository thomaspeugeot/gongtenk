package visuals

import (
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

var MapOptions = (&gongleaflet_models.MapOptions{

	Name:               "Whole France",
	Lat:                45,
	Lng:                4,
	ZoomLevel:          6,
	ZoomControl:        false,
	AttributionControl: true,
	ZoomSnap:           1,
	UrlTemplate:        "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
	Attribution:        "osm",
	MaxZoom:            18,
	LayerGroupUses: []*gongleaflet_models.LayerGroupUse{
		AircraftLayerGroupUse,
	},
}).Stage()

var AircraftLayerGroupUse = (&gongleaflet_models.LayerGroupUse{
	Name:       "Aircraft",
	LayerGroup: AircraftLayerGroup,
	Display:    true,
}).StageCopy()
