package responseCreators

import (
	yelp "review_microservice/reviewCollectors/yelp"
)

type YelpResponse struct {
	Response
}

func (y *YelpResponse) createRestInfo(reviews_by_id map[string]yelp.ReviewsOutput,
	photos_by_id map[string][]string, places_by_id map[string]yelp.PlaceDetails) {
	var rest_infos []RestaurantInfo
	for id, bo := range places_by_id {
		var cuisines []string
		for _, cuisine := range bo.Categories {
			cuisines = append(cuisines, cuisine.Title)
		}
		//creating the RestaurantInfo object for this particular restaurant, and adding it to the list
		rest_info := RestaurantInfo{
			Api_name:        "yelp",
			Restaurant_id:   id,
			Restaurant_name: bo.Name,
			Latitude:        bo.Coordinates.Latitude,
			Longitude:       bo.Coordinates.Longitude,
			Cuisine:         cuisines,
			AvgRating:       bo.Rating,
			Reviews:         createReviews(reviews_by_id[id]),
			Yelp_pictures:   photos_by_id[id],
			Website:         bo.Url,
		}
		rest_infos = append(rest_infos, rest_info)
	}

	y.Response.Restaurant_info = rest_infos
}

func createReviews(r yelp.ReviewsOutput) []ReviewObj {
	var reviews []ReviewObj
	for _, rev := range r.Reviews {
		review := ReviewObj{
			Username: rev.User.Name,
			Text:     rev.Text,
			Rating:   float64(rev.Rating),
		}
		reviews = append(reviews, review)
	}
	return reviews
}
