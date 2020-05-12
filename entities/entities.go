package entities

// Coordinates interface to calculate the distance for any coordinate system
/*type Coordinates interface {
	DistanceTo(searchPoint Coordinates) float64
	GetDistance() float64
}*/

// CartesianCoordinates represents a point in a Cartesian coordinate system
type CartesianCoordinates struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Distance float64 `json:"-"`
}

// DistanceTo get the Manhattan distance between searchPoint and c
/*func (c CartesianCoordinates) DistanceTo(searchPoint Coordinates) float64 {
	cartesianSearchPoint := searchPoint.(CartesianCoordinates)
	c.Distance = math.Abs(cartesianSearchPoint.X-c.X) + math.Abs(cartesianSearchPoint.Y-c.Y)
	return c.Distance
}*/

// GetDistance return the distance from this point to the searchPoint
/*func (c CartesianCoordinates) GetDistance() float64 {
	return c.Distance
}*/

// SetDistance modified the distance from this point to the searchPoint
/*func (c CartesianCoordinates) SetDistance(distance float64) {
	c.Distance = distance
}*/

// PointDistance could be used as an enhancement to keep Coordinates as a defined unit
/*type PointDistance struct {
	Point    CartesianCoordinates
	Distance float64
}*/
