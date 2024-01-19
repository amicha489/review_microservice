package responseCreators

import (
	yelp "review_microservice/reviewCollectors/yelp"
)

type YelpResponseCreator struct {
	Places_by_id map[string]yelp.PlaceDetails
	Revs_by_id   map[string]yelp.ReviewsOutput
	Pics_by_id   map[string][]string
}

func (yrc *YelpResponseCreator) CreateResponse() YelpResponse {
	//empty variable
	var yr YelpResponse

	//create the list of RestaurantInfo objects in the empty variable to populate it
	yr.createRestInfo(yrc.Revs_by_id, yrc.Pics_by_id, yrc.Places_by_id)
	return yr
}
