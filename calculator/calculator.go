package calculator

import (
	"entities"
	"math"
	"sort"
)

// GetPointsInRange Return a coordinates structure with the points which are within the maxDistance
// from the point indicated by xSearhPoint and ySearchPoint
func GetPointsInRange(maxDistance float64, xSearhPoint float64, ySearchPoint float64, points []entities.CartesianCoordinates) []entities.CartesianCoordinates {

	pointsInRange := make([]entities.CartesianCoordinates, 0)

	/*for i := 0; i < len(points); i++ {
		points[i].Distance = math.Abs(xSearhPoint-points[i].X) + math.Abs(ySearchPoint-points[i].Y)
		if distance := points[i].Distance; distance <= maxDistance {
			pointsInRange = append(pointsInRange, points[i])
		}
	}*/

	for _, point := range points {
		point.Distance = calculateDistance(point.X, xSearhPoint, point.Y, ySearchPoint) //math.Abs(xSearhPoint-point.X) + math.Abs(ySearchPoint-point.Y)
		if distance := point.Distance; distance <= maxDistance {
			pointsInRange = append(pointsInRange, point)
		}
	}

	/*for _, point := range points {
		if distance := point.DistanceTo(entities.CartesianCoordinates{xSearhPoint, ySearchPoint, 0}); distance <= maxDistance {
			pointsInRange = append(pointsInRange, point)
		}
	}*/

	sort.SliceStable(pointsInRange, func(i, j int) bool {
		return pointsInRange[i].Distance < pointsInRange[j].Distance
	})

	return pointsInRange
}

func calculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Abs(x2-x1) + math.Abs(y2-y1)
}
