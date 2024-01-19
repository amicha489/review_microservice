package responseCreators

type Response struct {
	Restaurant_info []RestaurantInfo `json:"restaurant_info"`
}

type RestaurantInfo struct {
	Api_name            string      `json:"api"`
	Restaurant_id       string      `json:"restaurant_id"`
	Restaurant_name     string      `json:"restaurant_name"`
	Latitude            float64     `json:"latitude"`
	Longitude           float64     `json:"longitude"`
	Cuisine             []string    `json:"cuisine"`
	AvgRating           float64     `json:"avg_rating"`
	Reviews             []ReviewObj `json:"reviews"`
	Yelp_pictures       []string    `json:"yelp_pictures"`
	GoogleMaps_pictures [][]byte    `json:"gm_pictures"`
	Website             string      `json:"website"`
}

type ReviewObj struct {
	Username string  `json:"username"`
	Text     string  `json:"text"`
	Rating   float64 `json:"rating"`
}
