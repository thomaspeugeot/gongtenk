package visuals

import (
	"log"
	"sort"

	"github.com/thomaspeugeot/gongtenk/go/icons"
	"github.com/thomaspeugeot/gongtenk/go/models"

	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
)

// attachVisualTrack attaches a visual track to track
func attachVisualTrack(track gongleaflet_models.VisualTrackInterface,
	divIcon *gongleaflet_models.DivIcon,
	colorEnum gongleaflet_models.ColorEnum,
	displayTrackHistory bool,
	displayLevelAndSpeed bool) {

	// sometimes, the visual icon is nil (not reproductible bug)
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}

	visualTrack := new(gongleaflet_models.VisualTrack).Stage()
	visualTrack.VisualTrackInterface = track
	visualTrack.DivIcon = divIcon
	visualTrack.DisplayTrackHistory = displayTrackHistory
	visualTrack.DisplayLevelAndSpeed = displayLevelAndSpeed
	visualTrack.ColorEnum = colorEnum
	visualTrack.UpdateTrack()
}

// attach visual center to center
func attachMarker(
	visualCenterInterface gongleaflet_models.MarkerInterface,
	colorEnum gongleaflet_models.ColorEnum,
	divIcon *gongleaflet_models.DivIcon) {
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}
	visualCenter := new(gongleaflet_models.Marker).Stage()
	visualCenter.MarkerInteface = visualCenterInterface
	visualCenter.ColorEnum = colorEnum
	visualCenter.DivIcon = divIcon
	visualCenter.UpdateMarker()
}

// attach visual line to line
func attachLine(
	visualLineInterface gongleaflet_models.LineInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualLine := new(gongleaflet_models.VLine).Stage()
	visualLine.DashStyleEnum = DashStyleEnum
	visualLine.LineInterface = visualLineInterface
	visualLine.UpdateLine()
}

// attach visual circle to circle
func attachCircle(
	visualCircleInterface gongleaflet_models.CircleInterface,
	DashStyleEnum gongleaflet_models.DashStyleEnum) {
	visualCircle := new(gongleaflet_models.Circle).Stage()
	visualCircle.DashStyleEnum = DashStyleEnum
	visualCircle.Circle = visualCircleInterface
	visualCircle.UpdateCircle()
}

func AttachVisualElementsToModelElements() {

	cityOrdered := []*models.City{}
	for city := range models.Stage.Citys {
		cityOrdered = append(cityOrdered, city)
	}
	// sort cities according to their population
	sort.Slice(cityOrdered[:], func(i, j int) bool {
		return cityOrdered[i].Population < cityOrdered[j].Population
	})

	for index, city := range cityOrdered {
		_ = city
		if index > models.ConfigurationSingloton.NumberOfCitiesToDisplay {
			continue
		}

		if city.Twin {
			attachVisualTrack(city, icons.Dot_10, gongleaflet_models.GREY, false, false)
		} else {
			attachVisualTrack(city, icons.Dot_10, gongleaflet_models.GREEN, false, false)
		}
	}

}
