package responseCreators

import (
	googleMaps "review_microservice/reviewCollectors/googleMaps"
)

type GoogleMapsResponseCreator struct {
	Place_by_id map[string]googleMaps.PlaceDetails
	Revs_by_id  map[string]googleMaps.PlaceRevsAndPics
	Pics_by_id  map[string][][]byte
}

func (gmrc *GoogleMapsResponseCreator) CreateResponse() GoogleMapsResponse {
	//empty variable
	var gmr GoogleMapsResponse

	//create the list of RestaurantInfo objects in the empty variable to populate it
	gmr.createRestInfo(gmrc.Place_by_id, gmrc.Revs_by_id, gmrc.Pics_by_id)
	return gmr
}
