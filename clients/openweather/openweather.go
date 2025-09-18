package openweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenWeatherClient struct {
	apiKey string
}

func New(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{
		apiKey: apiKey,
	}
}

func (o OpenWeatherClient) Coordinats(city string) (Coordinates, error) {
	url := "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
	resp, err := http.Get(fmt.Sprintf(url, city, o.apiKey))
	if err != nil {
		return Coordinates{}, fmt.Errorf("error get coordinates: %w, err")
	}
	if resp.StatusCode != 200 {
		return Coordinates{}, fmt.Errorf("error fAIL get coordinates: %d", resp.StatusCode)
	}

	var coordinatesResponse []CoordinantsResponse
	err = json.NewDecoder(resp.Body).Decode(&coordinatesResponse)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error unmarshal response: %w", err)
	}
	if len(coordinatesResponse) == 0 {
		return Coordinates{}, fmt.Errorf("error empty coordinates")
	}

	return Coordinates{
		Lat: coordinatesResponse[0].Lat,
		Lon: coordinatesResponse[0].Lon,
	}, nil
}
