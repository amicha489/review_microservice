package responseCreators

import (
	"review_microservice/reviewCollectors/yelp"
	"testing"
)

func TestYelpCreateRestInfo(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	reviews_by_id, photos_by_id, places_by_id := yelp.GetReviews(lat, lng)

	var yr YelpResponse
	yr.createRestInfo(reviews_by_id, photos_by_id, places_by_id)

	if len(yr.Restaurant_info) == 0 {
		t.Error("Fails")
	}
}

func TestYelpCreateResponse(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	reviews_by_id, photos_by_id, places_by_id := yelp.GetReviews(lat, lng)

	yrc := YelpResponseCreator{
		Places_by_id: places_by_id,
		Revs_by_id:   reviews_by_id,
		Pics_by_id:   photos_by_id,
	}

	yr := yrc.CreateResponse()

	if len(yr.Restaurant_info) == 0 {
		t.Error("Fails")
	}
}
