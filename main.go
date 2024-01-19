package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	response "review_microservice/responseCreators"
	"review_microservice/reviewCollectors/googleMaps"
	"review_microservice/reviewCollectors/yelp"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/reviews", ParallelReviewsHandler)
	log.Fatal(http.ListenAndServe(":6000", router))
}

func ParallelReviewsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	//parse Request Parameters -> need latitude + longitude
	query := r.URL.Query()
	latitude := query.Get("latitude")
	longitude := query.Get("longitude")
	lat, _ := strconv.ParseFloat(latitude, 64)
	lng, _ := strconv.ParseFloat(longitude, 64)
	//if need test location: "-33.8670522", "151.1957362"

	//define WaitGroup to ensure both goroutines finish before line 76
	var wg sync.WaitGroup

	//first goroutine
	wg.Add(1)
	yelp_channel := make(chan response.YelpResponse)
	go func(lat float64, lng float64) {
		defer wg.Done()
		//get reviews from Yelp
		reviews_by_id, photos_by_id, places_by_id := yelp.GetReviews(lat, lng)
		yrc := response.YelpResponseCreator{
			Places_by_id: places_by_id,
			Revs_by_id:   reviews_by_id,
			Pics_by_id:   photos_by_id,
		}
		yr := yrc.CreateResponse()
		yelp_channel <- yr
	}(lat, lng)
	yr := <-yelp_channel

	//second goroutine
	wg.Add(1)
	googleMaps_channel := make(chan response.GoogleMapsResponse)
	go func(lat float64, lng float64) {
		defer wg.Done()
		//get reviews from GoogleMaps
		place_by_id, revs_by_id, pics_by_id := googleMaps.GetReviews(lat, lng)
		gmrc := response.GoogleMapsResponseCreator{
			Place_by_id: place_by_id,
			Revs_by_id:  revs_by_id,
			Pics_by_id:  pics_by_id,
		}
		gmr := gmrc.CreateResponse()
		googleMaps_channel <- gmr
	}(lat, lng)
	gmr := <-googleMaps_channel

	//ensures that both goroutines finish before we continue from here
	wg.Wait()

	//combine them
	rest_info := append(yr.Restaurant_info, gmr.Restaurant_info...)

	//finalize the response
	response := response.Response{
		Restaurant_info: rest_info,
	}
	json.NewEncoder(w).Encode(response)
	fmt.Println("Response Sent")
}
