# sa8-review
Review Microservice

This microservice is written entirely in Golang. The code is split up into 3 core parts:
- The main thread which receives and consumes the API call from the Suggestion microservice. This thread also has 2 child threads (goroutines) for each of the 2 external APIs used
- The data collection code in the reviewCollectors/ folder. This is where each of the external APIs are called and their outputs are structured so that they can be passed to
- The response creator code in the responseCreators/ folder. This is where the Factory Method DP is used to create the response objects that are sent back to the Suggestion microservice.

## External APIs Used
- Yelp API
- GoogleMaps API

## File Structure

    .
    ├── responseCreators         
    │   ├── iResponse.go                    # part of Factory Method DP -> interface that each concrete object must extend
    │   ├── iResponseCreator.go             # part of Factory Method DP -> interface that each creator subclass must extend
    │   ├── Response.go                     # part of Factory Method DP -> generic Response object which GoogleMapsResponse and YelpResponse inherit from
    │   ├── googleMaps_response_creator.go  # part of Factory Method DP -> one of the creator subclasses
    │   ├── googleMaps_response.go          # part of Factory Method DP -> one of the concrete object subclasses
    │   ├── googleMaps_creator_test.go      # testing file for both the concrete GoogleMapsResponse object and the creator
    │   ├── yelp_response_creator.go        # part of Factory Method DP -> one of the creator subclasses
    │   ├── yelp_response.go                # part of Factory Method DP -> one of the concrete object subclasses
    │   ├── yelp_creator_test.go            # testing file for both the concrete YelpResponse object and the creator
    |               
    └── reviewCollectors
    |   └── apiCaller
    │   │   ├── apiCaller.go                  # helper functions which actually make the API calls to Yelp or GoogleMaps
    │   │   └── apiCaller_test.go             # testing file for these helper functions
    │   ├── googleMaps   
    │   │   ├── googleMaps.go                 # has the generic flow for what data we are collecting for GoogleMaps
    │   │   ├── googleMaps_place_details.go   # first API call to get restaurant IDs
    │   │   ├── googleMaps_place_photos.go    # subsequent API calls to get an image based on a photo_reference
    │   │   ├── googleMaps_place_reviews.go   # second API call to get the reviews and pictures for a specific restaurant ID
    │   │   └── googleMaps_collector_test.go  # testing file for all functions in this package
    │   ├── yelp   
    │   │   ├── yelp.go                       # has the generic flow for what data we are collecting for Yelp
    │   │   ├── yelp_place_details.go         # first API call to get restaurant IDs
    │   │   ├── yelp_place_photos.go          # third API call to get 3 image URLs from a specific restaurant ID
    │   │   ├── yelp_place_reviews.go         # second API call to get the reviews for a specific restaurant ID
    │   │   └── yelp_collector_test.go        # testing file for all functions in this package

## API
#### GET /reviews?latitude=[latitude]&longitude=[longitude]

## Run 
Instructions for running this service locally:
```powershell
$ go run main.go
```
This microservice will now listen on port 3000 and wait for a request (in the format specified above). When it has received a request, a text "Request Received" will print, and the text "Response Sent" will print once the response has been sent. This can be tested using POSTMAN by doing the following GET request: 
```powershell
GET localhost:3000/reviews?latitude=[latitude]&longitude=[longitude]
```

## Tests

Test files are spread throughout the repo so that each test file can test the package that it is in (this is a recommended practice in golang rather than putting all test files in one test/ folder). All test files end with _test.go and the tests can be run from the root of the directory by typing the following command:
```powershell
$ go test ./... -cover
```
The test coverage is at 94%.

## Deployment Information
This microservice has also been deployed to the following AWS EC2 Instance:
- Public IPv4 Address: 35.89.207.167
- Public IPv4 DNS address: ec2-35-89-207-167.us-west-2.compute.amazonaws.com
