package googleMaps

import (
	"review_microservice/reviewCollectors/apiCaller"
)

func collectGoogleMapsPictures(photos []PlacePhoto, key string) [][]byte {
	var pictures [][]byte
	var count int = 0
	for _, pr := range photos {
		//each photo has a photo reference. we need another API call to get the .PNG img of this photo (represented as a byte array)
		photo_ref := pr.Photo_reference
		url := "https://maps.googleapis.com/maps/api/place/photo?photo_reference=" + photo_ref + "&key=" + key + "&maxwidth=600"
		bodyBytes := apiCaller.MakeGoogleMapsApiCall(url)
		pictures = append(pictures, bodyBytes)
		count += 1
		if count == 3 { //yelp only outputs 3 pictures, so we might as well only take the top 3 googleMaps pictures (for consistency)
			break
		}
	}
	return pictures
}
