package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles API requests to the Pokemon API
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new API client with default configuration
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

// GetLocationArea fetches location data from the API
func (c *Client) GetLocationArea(url string) (LocationResponse, error) {
	if url == "" {
		url = c.baseURL + "/location"
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationResponse{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var locations LocationResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationResponse{}, err
	}

	return locations, nil
}

// ExploreArea fetches Pokemon in a specific area
func (c *Client) ExploreArea(areaName string) (AreaResponse, error) {
	url := fmt.Sprintf("%s/location-area/%s", c.baseURL, areaName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return AreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AreaResponse{}, fmt.Errorf("area not found: %s", areaName)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaResponse{}, err
	}

	var area AreaResponse
	err = json.Unmarshal(data, &area)
	if err != nil {
		return AreaResponse{}, err
	}

	return area, nil
}
