package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type WikipediaPage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Extract     string `json:"extract"`
}

func main() {
	// Initialize the Gin router.
	router := gin.Default()

	// Define the API route to search for a Wikipedia page.
	router.GET("/search/:query", searchWikipedia)

	// Run the server on port 8080.
	router.Run(":8080")
}

func searchWikipedia(c *gin.Context) {
	query := c.Param("query")
	page, err := getWikipediaPage(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from Wikipedia API"})
		return
	}

	c.JSON(http.StatusOK, page)
}

func getWikipediaPage(query string) (*WikipediaPage, error) {
	query = strings.ReplaceAll(query, " ", "_")
	apiURL := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", query)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wikipedia API returned non-200 status code: %d", resp.StatusCode)
	}

	var page WikipediaPage
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, err
	}

	return &page, nil
}
