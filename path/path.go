package path

import (
	"github.com/EliCDavis/mesh"
	"github.com/EliCDavis/vector"
)

// Path is a series of ordered points that make up a path
// through 3D space.
type Path []vector.Vector3

func (ls Path) Combine(path Path) Path {
	return append(ls, path...)
}

func (ls Path) Rotate(pivot vector.Vector3, rot mesh.Quaternion) Path {
	results := make([]vector.Vector3, len(ls))

	for i := 0; i < len(ls); i++ {
		results[i] = rot.Rotate(ls[i].Sub(pivot)).Add(pivot)
	}

	return results
}

func (ls Path) Translate(amt vector.Vector3) Path {
	results := make([]vector.Vector3, len(ls))

	for i := 0; i < len(ls); i++ {
		results[i] = ls[i].Add(amt)
	}

	return results
}

func (ls Path) Reverse() Path {
	results := make([]vector.Vector3, len(ls))

	for i := 0; i < len(ls); i++ {
		results[i] = ls[len(ls)-i-1]
	}

	return results
}
