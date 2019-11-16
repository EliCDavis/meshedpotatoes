package meshedpotatoes

import (
	"log"
	"math"

	"github.com/EliCDavis/mesh"
	"github.com/EliCDavis/vector"
)

// DrawPlaneShape attempts to draw a certain shape about the plane with a given
// number of sides. The radius is the distance from one of the vertices from
// the center
func DrawPlaneShape(center, normal vector.Vector3, radius float64, sides int) (mesh.Model, error) {
	perpendicular := normal.
		Perpendicular().
		Normalized().
		MultByConstant(radius)

	outerPoints := make([]vector.Vector3, sides)
	outerPoints[0] = perpendicular

	angleIncrement := (math.Pi * 2) / float64(sides)

	polys := make([]mesh.Polygon, sides)
	for i := 1; i < sides; i++ {
		rot := mesh.UnitQuaternionFromTheta(angleIncrement*float64(i), normal)
		outerPoints[i] = rot.Rotate(perpendicular)

		points := []vector.Vector3{
			center, outerPoints[i], outerPoints[i-1],
		}
		poly, _ := mesh.NewPolygon(points, points)
		polys[i] = poly
	}
	points := []vector.Vector3{
		center, outerPoints[0], outerPoints[sides-1],
	}
	poly, _ := mesh.NewPolygon(points, points)
	polys[0] = poly

	for i := 0; i < sides; i++ {
		log.Printf("%v", outerPoints[i])
	}

	m, _ := mesh.NewModel(polys)

	return m.Translate(center), nil
}

func GetPlaneOuterPoints(center, normal vector.Vector3, radius float64, sides int) []vector.Vector3 {
	perpendicular := normal.
		Perpendicular().
		Normalized().
		MultByConstant(radius)

	outerPoints := make([]vector.Vector3, sides)
	outerPoints[0] = perpendicular.Add(center)

	angleIncrement := (math.Pi * 2) / float64(sides)

	for i := 1; i < sides; i++ {
		rot := mesh.UnitQuaternionFromTheta(angleIncrement*float64(i), normal)
		outerPoints[i] = rot.Rotate(perpendicular).Add(center)
	}

	return outerPoints
}

// MakeSquare draws a square using two triangles
func MakeSquare(
	bottomLeft vector.Vector3,
	topLeft vector.Vector3,
	topRight vector.Vector3,
	bottomRight vector.Vector3,
) []mesh.Polygon {
	polys := make([]mesh.Polygon, 2)

	poly, _ := mesh.NewPolygon(
		[]vector.Vector3{bottomLeft, topLeft, bottomRight},
		[]vector.Vector3{bottomLeft, topLeft, bottomRight},
	)

	polys[0] = poly

	poly, _ = mesh.NewPolygon(
		[]vector.Vector3{topLeft, topRight, bottomRight},
		[]vector.Vector3{topLeft, topRight, bottomRight},
	)

	polys[1] = poly
	return polys
}
