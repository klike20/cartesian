package loader

import (
	"encoding/json"
	"entities"
	"io/ioutil"
)

var points []entities.CartesianCoordinates

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadCors(filePath string) {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	_ = json.Unmarshal([]byte(dat), &points)
}

// GetPoints loads a source database of points from the indicated path
// if reload is false, the load will not be done again if it was done before
func GetPoints(path string, reload bool) []entities.CartesianCoordinates {
	if points == nil || reload {
		loadCors(path)
	}
	return points
}
