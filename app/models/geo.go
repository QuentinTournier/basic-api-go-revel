package models

import (
	//"github.com/kpawlik/geojson"
	//"encoding/json"
	//"fmt"
)
import (
	//"fmt"
	"encoding/json"
	//"github.com/kpawlik/geojson"
	"github.com/kpawlik/geojson"
)

type JSONPoint geojson.Point

func (p *JSONPoint) MarshalJSON() ([]byte, error) {

	if len(p.Coordinates) < 2 {
		return json.Marshal(&struct {
			Lon geojson.CoordType `json:"lon"`
			Lat geojson.CoordType  `json:"lat"`
		}{
			Lon:     0,
			Lat:     0,
		});
	}

	return json.Marshal(&struct {
		Lon geojson.CoordType `json:"lon"`
		Lat geojson.CoordType  `json:"lat"`
	}{
		Lon:     p.Coordinates[0],
		Lat:     p.Coordinates[1],
	})
}

func (p *JSONPoint) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}{
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.Coordinates = geojson.Coordinate{geojson.CoordType(aux.Lon), geojson.CoordType(aux.Lat)}
	return nil
}