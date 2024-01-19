package yelp

import (
	"encoding/json"
	"fmt"
	"review_microservice/reviewCollectors/apiCaller"
)

type PlacesOutput struct {
	Total      int            `json:"total"`
	Businesses []PlaceDetails `json:"businesses"`
}

type PlaceDetails struct {
	Categories  []Categories `json:"categories"`
	Coordinates Coordinates  `json:"coordinates"`
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Rating      float64      `json:"rating"`
	Url         string       `json:"url"`
}

type Categories struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getPlaceDetails(key string, lat float64, lng float64) []PlaceDetails {
	lat_and_lng := fmt.Sprint(lat) + "&longitude=" + fmt.Sprint(lng)
	url := "https://api.yelp.com/v3/businesses/search?latitude=" + lat_and_lng + "&radius=10000&categories=restaurants"

	//make API call
	bodyBytes := apiCaller.MakeYelpApiCall(url, key)

	//convert byte[] to the struct we want to represent the data in
	var places_output PlacesOutput
	json.Unmarshal(bodyBytes, &places_output)

	return places_output.Businesses
}
