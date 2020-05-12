package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"calculator"

	"loader"

	"entities"
)

type App struct {
	Router *mux.Router
}

var points []entities.CartesianCoordinates

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the api points!")
	fmt.Println("hitting api points")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	points = loader.GetPoints(path, true)
}

func pointsHandler(w http.ResponseWriter, r *http.Request) {
	maxDistance, _ := strconv.ParseFloat(r.FormValue("distance"), 64)
	xSearchPoint, _ := strconv.ParseFloat(r.FormValue("x"), 64)
	ySearchPoint, _ := strconv.ParseFloat(r.FormValue("y"), 64)

	fmt.Println("the distance! ", maxDistance)
	fmt.Println("the x! ", xSearchPoint)
	fmt.Println("the y! ", ySearchPoint)

	pointsInRange := calculator.GetPointsInRange(maxDistance, xSearchPoint, ySearchPoint, points)

	json.NewEncoder(w).Encode(pointsInRange)
}

func (a *App) handleRequests(start bool) {
	//with Mux
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/api/points/path", pathHandler).Methods("POST", "GET")
	//myRouter.Path("/api/pointsOld").Queries("distance", "{distance}", "x", "{x}", "y", "{y}").HandlerFunc(coordinates)
	myRouter.Path("/api/points").Queries("distance", "{distance}", "x", "{x}", "y", "{y}").HandlerFunc(pointsHandler)

	//with http direclty
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/cor", pointsHandler)
	//log.Fatal(http.ListenAndServe(":10000", nil))

	a.Router = myRouter
}

func (a *App) init() {
	log.Fatal(http.ListenAndServe(":10000", a.Router))
}

func loadCors() {
	//second approach
	points = loader.GetPoints("data.json", false)
}

func main() {
	a := App{}
	loadCors()
	a.handleRequests(true)
	a.init()
}
