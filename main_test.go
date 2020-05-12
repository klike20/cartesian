package main

import (
	"calculator"
	"entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var a App

func TestInit(t *testing.T) {
	a.handleRequests(false)
	req, _ := http.NewRequest("GET", "/api/points/path?path=newdata.json", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	} else {
		t.Log("Pass", t.Name())
	}
}

func checkResponse(t *testing.T, actual, expected string) {
	if strings.TrimSuffix(actual, "\n") != expected {
		t.Errorf("Expected %s. Got %s", expected, actual)
	} else {
		t.Log("Pass", t.Name())
	}
}

func TestApiPointsMissingParameter(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/points?distance=10&y=0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	checkResponse(t, response.Body.String(), "404 page not found")
}

func TestFoundApiPoints(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/points?distance=10&y=0&x=0", nil)

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponse(t, response.Body.String(), "[{\"x\":5,\"y\":5}]")
}

func TestNotFoundApiPoints(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/points?distance=5&y=0&x=0", nil)

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponse(t, response.Body.String(), "[]")
}

func TestCalculator(t *testing.T) {

	points := []entities.CartesianCoordinates{{2, 2, 0}, {3, 3, 0}}

	pointInRange := calculator.GetPointsInRange(3, 0, 0, points)

	if len(pointInRange) > 1 {
		t.Errorf("Expected 1. Got %d", len(pointInRange))
	} else {
		t.Log("Pass", t.Name())
	}
}
