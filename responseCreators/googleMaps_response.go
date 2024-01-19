package responseCreators

import (
	"review_microservice/reviewCollectors/googleMaps"
)

type GoogleMapsResponse struct {
	Response
}

func (gmr *GoogleMapsResponse) createRestInfo(place_by_id map[string]googleMaps.PlaceDetails,
	revs_by_id map[string]googleMaps.PlaceRevsAndPics, pics_by_id map[string][][]byte) {
	var rest_infos []RestaurantInfo
	for id, place := range place_by_id {
		revs_and_pics := revs_by_id[id]
		//creating the RestaurantInfo object for this particular restaurant, and adding it to the list
		rest_info := RestaurantInfo{
			Api_name:            "googleMaps",
			Restaurant_id:       id,
			Restaurant_name:     place.Name,
			Latitude:            place.Geometry.Location.Latitude,
			Longitude:           place.Geometry.Location.Longitude,
			Cuisine:             place.Types,
			AvgRating:           place.Rating,
			Reviews:             createReviewObjs(revs_and_pics.Reviews),
			GoogleMaps_pictures: pics_by_id[id],
			Website:             revs_by_id[id].Website,
		}
		rest_infos = append(rest_infos, rest_info)
	}

	gmr.Response.Restaurant_info = rest_infos
}

func createReviewObjs(revs []googleMaps.PlaceReview) []ReviewObj {
	var reviews []ReviewObj
	for _, rev := range revs {
		review := ReviewObj{
			Username: rev.Username,
			Text:     rev.Text,
			Rating:   rev.Rating,
		}
		reviews = append(reviews, review)
	}
	return reviews
}
