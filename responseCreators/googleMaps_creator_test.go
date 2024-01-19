package responseCreators

import (
	"review_microservice/reviewCollectors/googleMaps"
	"testing"
)

func TestGoogleMapsCreateRestInfo(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	place_by_id, revs_by_id, pics_by_id := googleMaps.GetReviews(lat, lng)

	var gmr GoogleMapsResponse
	gmr.createRestInfo(place_by_id, revs_by_id, pics_by_id)

	if len(gmr.Restaurant_info) == 0 {
		t.Error("Fails")
	}
}

func TestGoogleMapsCreateResponse(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	place_by_id, revs_by_id, pics_by_id := googleMaps.GetReviews(lat, lng)

	gmrc := GoogleMapsResponseCreator{
		Place_by_id: place_by_id,
		Revs_by_id:  revs_by_id,
		Pics_by_id:  pics_by_id,
	}

	gmr := gmrc.CreateResponse()

	if len(gmr.Restaurant_info) == 0 {
		t.Error("Fails")
	}
}
