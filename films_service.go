package films

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"time"
)

const filmsEndpointUrl string = "https://toolbox.palette-adv.spectrocloud.com:5002/films"
// const filmsEndpointUrl = "http://localhost:8080/films"

// Your API token. Needed to successfully authenticate when calling the films endpoint.
// Must be included in the Authorization header in the request sent to the films endpoint.
const apiToken string = "8c5996d5-fb89-46c9-8821-7063cfbc18b1"

type Film struct {
	Name         *string  `json:'name'`
	Length       *int32   `json:'length'`
	Rating       *float64 `json:'rating'`
	ReleaseDate  *string  `json:'releaseDate'`
	DirectorName *string  `json:'directorName'`
}

var films []Film
var filmMap map[string][]Film

func init() {
	// Create a new request using http
	req, getErr := http.NewRequest("GET", filmsEndpointUrl, nil)
	if getErr != nil {
		log.Fatal(getErr)
		panic(getErr)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", "Bearer "+apiToken)

	// Send req using http Client
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if getErr != nil {
		log.Fatal(respErr)
		panic(respErr)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("HTTP Request failed with status code: %d\n", resp.StatusCode)
		log.Fatal(err)
		panic(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
		panic(readErr)
	}

	jsonErr := json.Unmarshal(body, &films)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		panic(jsonErr)
	}

	filmMap = map[string][]Film{}
	for _, film := range films {
		filmMap[*film.DirectorName] = append(filmMap[*film.DirectorName], film)
	}

}

// GetFilms retrieves the data for all films by calling the https://toolbox.palette-adv.spectrocloud.com:5002/films endpoint.
func GetFilms() []Film {
	return films
}

// BestRatedFilm retrieves the name of the best rated film that was directed by the director with the given name.
// If there are no films directed by the given director, return an empty string.
// Note: there will only be one film with the best rating.
func BestRatedFilm(directorName string) string {
	highestSoFar := 0.0
	bestRatedFilm := ""

	filmArray := filmMap[directorName]
	for _, film := range filmArray {
		if *film.DirectorName == directorName && *film.Rating > highestSoFar {
			bestRatedFilm = *film.Name
			highestSoFar = *film.Rating
		}
	}

	return bestRatedFilm
}

// DirectorWithMostFilms retrieves the name of the director who has directed the most films
// in the CodeScreen Film service.
func DirectorWithMostFilms() string {
	mostFilms := 0
	directorName := ""

	for k, v := range filmMap {
		totalFilms := len(v)
		if totalFilms > mostFilms {
			mostFilms = totalFilms
			directorName = k
		}
	}

	return directorName
}

// AverageRating retrieves the average rating for the films directed by the given director, rounded to 1 decimal place.
// If there are no films directed by the given director, return 0.0.
func AverageRating(directorName string) float64 {
	films := filmMap[directorName]
	totalRatings := 0.0
	totalFilms := len(films)
	if totalFilms == 0 {
		return 0.0
	}

	for _, f := range films {
		totalRatings += *f.Rating
	}

	return roundToDecimalPlaceHelper(totalRatings/float64(totalFilms), 1)
}

/*
ShortestFilmReleaseGap retrieves the shortest number of days between any two film releases directed by the given director.
If there are no films directed by the given director, return 0.
If there is only one film directed by the given director, return 0.
Note: no director released more than one film on any given day.

For example, if the service returns the following 3 films:

	{
	    "name": "Batman Begins",
	    "length": 140,
	    "rating": 8.2,
	    "releaseDate": "2006-06-16",
	    "directorName": "Christopher Nolan"
	},

	{
	    "name": "Interstellar",
	    "length": 169,
	    "rating": 8.6,
	    "releaseDate": "2014-11-07",
	    "directorName": "Christopher Nolan"
	},

	{
	    "name": "Prestige",
	    "length": 130,
	    "rating": 8.5,
	    "releaseDate": "2006-11-10",
	    "directorName": "Christopher Nolan"
	}

Then this method should return 147 for Christopher Nolan, as Prestige was released 147 days after Batman Begins.
*/
func ShortestFilmReleaseGap(directorName string) int {
	films := filmMap[directorName]
	if len(films) == 0 {
		return 0
	}

	// Parse date strings into time.Time objects
	var dates []time.Time
	for _, film := range films {
		date, err := time.Parse("2006-01-02", *film.ReleaseDate)
		if err != nil {
			log.Fatal(fmt.Errorf("Error parsing date: %v\n", err))
			return 0
		}
		dates = append(dates, date)
	}

	// Sort the dates in ascending order
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	// Calculate the difference in days between adjacent dates
	var minDays int
	for i := 1; i < len(dates); i++ {
		diff := int(dates[i].Sub(dates[i-1]).Hours() / 24)
		if i == 1 || diff < minDays {
			minDays = diff
		}
	}

	return minDays
}

func roundToDecimalPlaceHelper(number float64, decimalPlaces int) float64 {
	// Calculate the multiplier to shift the decimal point
	multiplier := math.Pow(10, float64(decimalPlaces))

	// Round the number to the nearest integer
	rounded := math.Round(number * multiplier)

	// Shift the decimal point back to its original position
	rounded /= multiplier

	return rounded
}
