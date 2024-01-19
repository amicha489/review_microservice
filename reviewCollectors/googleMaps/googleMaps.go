package googleMaps

// call GoogleMaps Places API to get reviews
func GetReviews(lat float64, lng float64) (map[string]PlaceDetails, map[string]PlaceRevsAndPics, map[string][][]byte) {
	//TODO: ideally this would be an env variable
	GOOGLEMAPS_API_KEY := "AIzaSyBzAt7HyaMP9D6nw4TT1pXlYta15bvJm1U"

	//get place details (including place IDs) of restaurants near a user's location
	places := getPlaceDetails(GOOGLEMAPS_API_KEY, lat, lng)

	places_by_id := make(map[string]PlaceDetails)
	revs_by_id := make(map[string]PlaceRevsAndPics)
	pics_by_id := make(map[string][][]byte)

	for _, place := range places {
		//get reviews and pictures for each restaurant
		revs, pics := collectGoogleMapsReviewsAndPics(place.Id, GOOGLEMAPS_API_KEY)
		places_by_id[place.Id] = place
		revs_by_id[place.Id] = revs
		pics_by_id[place.Id] = pics
	}

	//forward collection result to creators
	return places_by_id, revs_by_id, pics_by_id
}
