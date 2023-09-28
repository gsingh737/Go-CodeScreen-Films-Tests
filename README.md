# Go-CodeScreen-Films-Test

The CodeScreen Films API is a service that contains one endpoint,<br/>
`GET https://toolbox.palette-adv.spectrocloud.com:5002/films`, which returns the details of a large number of different films.

When you send an `HTTP GET` request to the endpoint, the response will be a `200 OK`, which includes a body containing a list of film data in `JSON` format. 
<br>

For authentication, you need to send your API token in the `Authorization HTTP header` using the [Bearer authentication scheme](https://tools.ietf.org/html/draft-ietf-oauth-v2-bearer-20#section-2.1). Your API token is `8c5996d5-fb89-46c9-8821-7063cfbc18b1`.

Here is an example of how to send the request from cURL:

    curl -H "Authorization: Bearer 8c5996d5-fb89-46c9-8821-7063cfbc18b1" \
    https://toolbox.palette-adv.spectrocloud.com:5002/films
    
An example response is the following:

     [
       {
         "name": "Batman Begins",
         "length": 140,
         "rating": 8.2,
         "releaseDate": "2006-06-16",
         "directorName": "Christopher Nolan"
       },
       {
         "name": "Alien",
         "length": 117,
         "rating": 8.4,
         "releaseDate": "1979-09-06",
         "directorName": "Ridley Scott"
       }
     ]


The `name` field represents the name of the film. The `length` field represents the duration of the film in minutes. The `rating` is the <a href="https://www.imdb.com/" target="_blank">`IMDb`</a> rating for the film, out of 10.
The `releaseDate` is the date in which the film was released in the United Kingdom, and the `directorName` field is the name of the film's director.

## Your Task

You are required to implement all the functions marked with `TODO Implement` in [films_service.go](films_service.go) in such a way that all the unit tests in the [films_service_test.go](films_service_test.go) file pass.

The functions in `films_service.go` should be implemented in such a way that you only need to call out to the CodeScreen Films API endpoint once per full run of the `films_service_test.go` test suite.

## Requirements

The [films_service_test.go](films_service_test.go) file should not be modified. If you would like to add your own unit tests, you
can add these in a separate file.

You are free to use whichever libraries you want when implementing your solution. </br>

The `go.mod` file should only be modified in order to add any third-party dependencies required for your solution. The `ginkgo` and `gomega` versions should not be changed.

The [.github/workflows/go.yml](.github/workflows/go.yml) file should also not be modified.

Your solution must use/be compatible with `Go` version `1.19`.

## Tests
Run `go test` to run the unit tests. These should all pass if your solution has been implemented correctly.

##

This test should take no longer than 1 hour to complete successfully.

Good luck!
