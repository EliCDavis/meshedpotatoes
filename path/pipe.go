package path

import (
	"errors"

	"github.com/EliCDavis/mesh"
	"github.com/EliCDavis/meshedpotatoes"
	"github.com/EliCDavis/vector"
)

// CreatePipe draws a pipe using the line segment as guides
func (ls Path) CreatePipe(pipeThickness float64, sides int) (mesh.Model, error) {
	if len(ls) < 2 {
		return mesh.Model{}, errors.New("Unable to create a pipe with less than 2 points")
	}

	points := make([][]vector.Vector3, len(ls))
	for i, p := range ls {
		var dir vector.Vector3

		if i == 0 {
			dir = ls[1].Sub(ls[0])
		} else {
			dir = ls[i].Sub(ls[i-1])
		}

		points[i] = meshedpotatoes.GetPlaneOuterPoints(p, dir, pipeThickness, sides)
	}

	polygons := make([]mesh.Polygon, 0)

	for i := 1; i < len(points); i++ {
		for p := 0; p < len(points[i]); p++ {
			polygons = append(polygons, meshedpotatoes.MakeSquare(
				points[i-1][p],
				points[i][p],
				points[i][(p+1)%len(points[i])],
				points[i-1][(p+1)%len(points[i])],
			)...)
		}
	}

	return mesh.NewModel(polygons)
}

// CreatePipeWithVarryingThickness draws a pipe using the line segment as guides
func (ls Path) CreatePipeWithVarryingThickness(thicknesses []float64, sides int) (mesh.Model, error) {
	if len(ls) < 2 {
		return mesh.Model{}, errors.New("Unable to create a pipe with less than 2 points")
	}

	if len(ls) != len(thicknesses) {
		return mesh.Model{}, errors.New("We need a thickness value per point on the line")
	}

	points := make([][]vector.Vector3, len(ls))
	for i, p := range ls {
		var dir vector.Vector3

		if i == 0 {
			dir = ls[1].Sub(ls[0])
		} else {
			dir = ls[i].Sub(ls[i-1])
		}

		points[i] = meshedpotatoes.GetPlaneOuterPoints(p, dir, thicknesses[i], sides)
	}

	polygons := make([]mesh.Polygon, 0)

	for i := 1; i < len(points); i++ {
		for p := 0; p < len(points[i]); p++ {
			polygons = append(polygons, meshedpotatoes.MakeSquare(
				points[i-1][p],
				points[i][p],
				points[i][(p+1)%len(points[i])],
				points[i-1][(p+1)%len(points[i])],
			)...)
		}
	}

	return mesh.NewModel(polygons)
}
