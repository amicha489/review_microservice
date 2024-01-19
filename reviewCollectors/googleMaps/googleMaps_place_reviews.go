package googleMaps

import (
	"encoding/json"
	"review_microservice/reviewCollectors/apiCaller"
)

type PlaceDetailsResponse struct {
	Result PlaceRevsAndPics `json:"result"`
}

type PlaceRevsAndPics struct {
	Icon    string        `json:"icon"`
	Photos  []PlacePhoto  `json:"photos"`
	Reviews []PlaceReview `json:"reviews"`
	Website string        `json:"website"`
}

type PlacePhoto struct {
	Photo_reference string   `json:"photo_reference"`
	Html            []string `json:"html_attributions"`
	Height          float32  `json:"height"`
	Width           float32  `json:"width"`
}

type PlaceReview struct {
	Username string  `json:"author_name"`
	Rating   float64 `json:"rating"`
	Text     string  `json:"text"`
}

func collectGoogleMapsReviewsAndPics(place_id string, key string) (PlaceRevsAndPics, [][]byte) {
	//get reviews & pictures for a specific restaurant
	url := "https://maps.googleapis.com/maps/api/place/details/json?place_id=" + place_id + "&key=" + key
	//make API call
	bodyBytes := apiCaller.MakeGoogleMapsApiCall(url)

	//convert byte[] to struct we want to represent our data as
	var place_details PlaceDetailsResponse
	json.Unmarshal(bodyBytes, &place_details)

	//the photos are only a photo_reference, we need to convert these to byte arrays
	pictures := collectGoogleMapsPictures(place_details.Result.Photos, key)

	return place_details.Result, pictures
}
