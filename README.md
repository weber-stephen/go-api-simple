# go-api-simple
A simple Go REST API example

## Getting Started

To run the API, open a terminal in the same directory as the main.go file and execute the following command:

```bash
go run main.go
```

The server will start, and you can access the API endpoint using tools like curl, Postman, or web browsers. For example:

To search for a Wikipedia page about "Golang": GET `http://localhost:8080/search/Golang`

The API will return information about the Wikipedia page, including the page title, description, and extract. Note that the example above is a basic demonstration of making API requests to Wikipedia. In a real-world scenario, you may want to add more error handling, caching, or rate limiting to improve the API's reliability and performance.