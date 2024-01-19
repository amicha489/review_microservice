package apiCaller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//why do we need these 2 functions????
//For the Yelp API, we pass the API_KEY as an authorization header
//For the GoogleMaps API, we pass the API_KEY directly in the request url

// return the byte[] of the Response Body of the Yelp API call
func MakeYelpApiCall(url string, key string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Authorization", "Bearer "+key)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	//read response Body of YelpAPI as byte[]
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

// return the byte[] of the Response Body of the GoogleMaps API call
func MakeGoogleMapsApiCall(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	//read response Body of GoogleMaps as byte[]
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}
