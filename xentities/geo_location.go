package xentities

type GeoType string

const (
	GeoTypePoint      GeoType = "point"
	GeoTypePolygon    GeoType = "polygon"
	GeoTypeLineString GeoType = "line_string"
)

type GeoPoint struct {
	Type        GeoType   `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
type GeoGeometry struct {
	Type        GeoType     `json:"type" bson:"type"`
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

func SetPointLocation(longitude, latitude float64) GeoPoint {
	return GeoPoint{
		Type:        GeoTypePoint,
		Coordinates: []float64{longitude, latitude},
	}
}

func SetPolygonLocation(coordinates [][]float64) GeoGeometry {
	return GeoGeometry{
		Type:        GeoTypePolygon,
		Coordinates: coordinates,
	}
}

func SetLineStringLocation(coordinates [][]float64) GeoGeometry {
	return GeoGeometry{
		Type:        GeoTypeLineString,
		Coordinates: coordinates,
	}
}
