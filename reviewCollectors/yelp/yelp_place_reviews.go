package yelp

import (
	"encoding/json"
	"review_microservice/reviewCollectors/apiCaller"
)

type ReviewsOutput struct {
	Reviews []YelpReview `json:"reviews"`
}

type YelpReview struct {
	Text   string   `json:"text"`
	Rating int      `json:"rating"`
	User   YelpUser `json:"user"`
}

type YelpUser struct {
	Id          string `json:"id"`
	Profile_url string `json:"profile_url"`
	Image_url   string `json:"image_url"`
	Name        string `json:"name"`
}

func collectYelpReviews(key string, place_id string) ReviewsOutput {
	// we want the reviews of a particular place_id
	url := "https://api.yelp.com/v3/businesses/" + place_id + "/reviews"

	//make API call
	bodyBytes := apiCaller.MakeYelpApiCall(url, key)

	//convert byte[] to the struct we want to represent the data
	var reviews_output ReviewsOutput
	json.Unmarshal(bodyBytes, &reviews_output)

	return reviews_output
}
