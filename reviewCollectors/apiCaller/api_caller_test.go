package apiCaller

import (
	"encoding/json"
	"fmt"
	"testing"
)

type GMO struct {
	Status string `json:"status"`
}

func TestGoogleMapsInvalidAPIKey(t *testing.T) {
	location := fmt.Sprint(-33.8670522) + "," + fmt.Sprint(151.1957362)
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + location + "&radius=10000&type=restaurant&keyword=restaurant&key="
	b := MakeGoogleMapsApiCall(url)

	var gmo GMO
	json.Unmarshal(b, &gmo)

	if gmo.Status != "REQUEST_DENIED" {
		t.Error("Fails")
	}
}

func TestGoogleMapsValidAPIKey(t *testing.T) {
	location := fmt.Sprint(-33.8670522) + "," + fmt.Sprint(151.1957362)
	key := "AIzaSyBzAt7HyaMP9D6nw4TT1pXlYta15bvJm1U"
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + location + "&radius=10000&type=restaurant&keyword=restaurant&key=" + key
	b := MakeGoogleMapsApiCall(url)

	var gmo GMO
	json.Unmarshal(b, &gmo)

	if gmo.Status == "REQUEST_DENIED" {
		t.Error("Fails")
	}
}

type YO struct {
	Err Error `json:"error"`
}

type Error struct {
	Code string `json:"code"`
}

func TestYelpInvalidAPIKey(t *testing.T) {
	location := fmt.Sprint(-33.8670522) + "," + fmt.Sprint(151.1957362)
	YELP_API_KEY := ""
	url := "https://api.yelp.com/v3/businesses/search?latitude=" + location + "&radius=10000&categories=restaurants"
	b := MakeYelpApiCall(url, YELP_API_KEY)

	var yo YO
	json.Unmarshal(b, &yo)

	if yo.Err.Code != "VALIDATION_ERROR" {
		t.Error("Fails")
	}
}

func TestYelpValidAPIKey(t *testing.T) {
	location := fmt.Sprint(-33.8670522) + "&longitude=" + fmt.Sprint(151.1957362)
	YELP_API_KEY := "bqu8D-bZ3rivKL0-qM-dWbjabmAg3CwvFUt-euCwxBrXz8EaITEq0gkEhW1VpFQNJfHu8nEt8adzipapm7UympwsgfXXbocuqavvEVK7xDvkNgACaoQeG1jKjjZ1Y3Yx"
	url := "https://api.yelp.com/v3/businesses/search?latitude=" + location + "&radius=10000&categories=restaurants"
	b := MakeYelpApiCall(url, YELP_API_KEY)

	var yo YO
	json.Unmarshal(b, &yo)

	if yo.Err.Code == "VALIDATION_ERROR" {
		t.Error("Fails")
	}
}
