# go-api-simple
A simple Go API implementation that searches on Wikipedia

## About the API

Now, let's explain the code:

We define a WikipediaPage struct to represent the data we want to extract from the Wikipedia API response, including Title, Description, and Extract fields.
In the main function, we create a new Gin router and define a route "/search/:query" to search for a Wikipedia page based on the provided query.
The `searchWikipedia` function handles the API endpoint. It extracts the query parameter, calls the `getWikipediaPage` function to fetch data from the Wikipedia API, and returns the result as a JSON response.
The `getWikipediaPage` function takes a query, constructs the Wikipedia API URL with the given query, and makes an HTTP GET request to the Wikipedia API.
If the API request is successful (HTTP status code 200), it parses the JSON response into a WikipediaPage struct and returns it.

## Getting Started

To run the API, open a terminal in the same directory as the main.go file and execute the following command:

```bash
go run main.go
```

The server will start, and you can access the API endpoint using tools like curl, Postman, or web browsers. For example:

To search for a Wikipedia page about "Golang": GET `http://localhost:8080/search/Golang`

The API will return information about the Wikipedia page, including the page title, description, and extract. Note that the example above is a basic demonstration of making API requests to Wikipedia. In a real-world scenario, you may want to add more error handling, caching, or rate limiting to improve the API's reliability and performance.

