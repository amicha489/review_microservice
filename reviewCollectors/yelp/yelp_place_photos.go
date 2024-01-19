package yelp

import (
	"encoding/json"
	"review_microservice/reviewCollectors/apiCaller"
)

type PhotosOutput struct {
	Photos []string `json:"photos"`
}

func collectYelpPhotos(key string, place_id string) []string {
	// we want the photos of a particular place_id
	url := "https://api.yelp.com/v3/businesses/" + place_id

	//make API call
	bodyBytes := apiCaller.MakeYelpApiCall(url, key)

	//convert byte[] to the struct we want to represent the data
	var photos_output PhotosOutput
	json.Unmarshal(bodyBytes, &photos_output)

	return photos_output.Photos
}
