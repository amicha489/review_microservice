package yelp

import (
	"testing"
)

var YELP_API_KEY = "bqu8D-bZ3rivKL0-qM-dWbjabmAg3CwvFUt-euCwxBrXz8EaITEq0gkEhW1VpFQNJfHu8nEt8adzipapm7UympwsgfXXbocuqavvEVK7xDvkNgACaoQeG1jKjjZ1Y3Yx"

func TestYelpDetails(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362

	details := getPlaceDetails(YELP_API_KEY, lat, lng)

	if len(details) < 1 {
		t.Error("Fails")
	}
}

func TestYelpPhotos(t *testing.T) {
	pics := collectYelpPhotos(YELP_API_KEY, "NqGNHWrxfxEySH4mNoSK2Q")

	if len(pics) != 3 {
		t.Error("Fails")
	}
}

func TestYelpReviews(t *testing.T) {
	revs := collectYelpReviews(YELP_API_KEY, "NqGNHWrxfxEySH4mNoSK2Q")

	if len(revs.Reviews) < 1 {
		t.Error("Fails")
	}
}

func TestYelpCollector(t *testing.T) {
	lat := -33.8670522
	lng := 151.1957362
	reviews_by_id, photos_by_id, places_by_id := GetReviews(lat, lng)

	if len(reviews_by_id) < 1 {
		t.Error("Fails")
	}
	if len(photos_by_id) < 1 {
		t.Error("Fails")
	}
	if len(places_by_id) < 1 {
		t.Error("Fails")
	}
}
