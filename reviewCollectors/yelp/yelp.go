package yelp

// call Yelp API to get reviews
func GetReviews(lat float64, lng float64) (map[string]ReviewsOutput, map[string][]string, map[string]PlaceDetails) {
	//TODO: ideally this would be an env variable
	YELP_API_KEY := "bqu8D-bZ3rivKL0-qM-dWbjabmAg3CwvFUt-euCwxBrXz8EaITEq0gkEhW1VpFQNJfHu8nEt8adzipapm7UympwsgfXXbocuqavvEVK7xDvkNgACaoQeG1jKjjZ1Y3Yx"

	//get place details (including place IDs) for all restaurants near the user's location
	place_details := getPlaceDetails(YELP_API_KEY, lat, lng)

	reviews_by_id := make(map[string]ReviewsOutput)
	photos_by_id := make(map[string][]string)
	places_by_id := make(map[string]PlaceDetails)

	//for each restaurant, collect the reviews from Yelp
	for _, place := range place_details {
		reviews_output := collectYelpReviews(YELP_API_KEY, place.Id)
		photos := collectYelpPhotos(YELP_API_KEY, place.Id)
		reviews_by_id[place.Id] = reviews_output
		photos_by_id[place.Id] = photos
		places_by_id[place.Id] = place
	}

	//forward this collection output to the creators
	return reviews_by_id, photos_by_id, places_by_id
}
