package googleMaps

import (
	"encoding/json"
	"fmt"
	"review_microservice/reviewCollectors/apiCaller"
)

type PlacesOutput struct {
	Results []PlaceDetails `json:"results"`
}

type PlaceDetails struct {
	Name     string   `json:"name"`
	Id       string   `json:"place_id"`
	Rating   float64  `json:"rating"`
	Geometry Geometry `json:"geometry"`
	Types    []string `json:"types"`
}

type Geometry struct {
	Location LatLng `json:"location"`
}

type LatLng struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func getPlaceDetails(key string, latitude float64, longitude float64) []PlaceDetails {
	location := fmt.Sprint(latitude) + "," + fmt.Sprint(longitude)
	url := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + location + "&radius=10000&type=restaurant&keyword=restaurant&key=" + key
	//make API call
	bodyBytes := apiCaller.MakeGoogleMapsApiCall(url)

	//convert byte[] to the struct we want to represent our data as
	var places_output PlacesOutput
	json.Unmarshal(bodyBytes, &places_output)

	return places_output.Results
}
