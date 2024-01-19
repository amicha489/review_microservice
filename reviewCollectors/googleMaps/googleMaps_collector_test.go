package googleMaps

import (
	"testing"
)

var GOOGLEMAPS_API_KEY = "AIzaSyBzAt7HyaMP9D6nw4TT1pXlYta15bvJm1U"

func TestGoogleMapsDetails(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362

	details := getPlaceDetails(GOOGLEMAPS_API_KEY, lat, lng)

	if len(details) <= 0 {
		t.Error("Fails")
	}
}

func TestGoogleMapsPhotos(t *testing.T) {
	var photos []PlacePhoto
	p := PlacePhoto{
		Photo_reference: "CnRvAAAAwMpdHeWlXl-lH0vp7lez4znKPIWSWvgvZFISdKx45AwJVP1Qp37YOrH7sqHMJ8C-vBDC546decipPHchJhHZL94RcTUfPa1jWzo-rSHaTlbNtjh-N68RkcToUCuY9v2HNpo5mziqkir37WU8FJEqVBIQ4k938TI3e7bf8xq-uwDZcxoUbO_ZJzPxremiQurAYzCTwRhE_V0",
	}
	photos = append(photos, p)

	pics := collectGoogleMapsPictures(photos, GOOGLEMAPS_API_KEY)

	if len(pics) != 1 {
		t.Error("Fails")
	}
}

func TestGoogleMapsRevsAndPics(t *testing.T) {
	revs, pics := collectGoogleMapsReviewsAndPics("ChIJN1t_tDeuEmsRUsoyG83frY4", GOOGLEMAPS_API_KEY)

	if len(revs.Photos) < 1 {
		t.Error("Fails")
	}
	if len(revs.Reviews) < 1 {
		t.Error("Fails")
	}
	if len(pics) < 1 {
		t.Error("Fails")
	}
}

func TestGoogleMapsCollector(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	place_by_id, revs_by_id, pics_by_id := GetReviews(lat, lng)

	if len(place_by_id) < 1 {
		t.Error("Fails")
	}
	if len(revs_by_id) < 1 {
		t.Error("Fails")
	}
	if len(pics_by_id) < 1 {
		t.Error("Fails")
	}
}
