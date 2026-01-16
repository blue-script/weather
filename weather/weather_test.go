package weather_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/blue-script/weather/geo"
	"github.com/blue-script/weather/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "Smolensk"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 4

	result, err := weather.GetWeather(geoData, format)
	isCorrectCity := strings.Contains(result, expected)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if !isCorrectCity {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetWeatherInvalidFormat(t *testing.T) {
	expected := "Smolensk"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 125

	_, err := weather.GetWeather(geoData, format)

	if err == nil {
		t.Fatalf("expected error %v, got nil", weather.ErrInvalidFormat)
	}
	if !errors.Is(err, weather.ErrInvalidFormat) {
		t.Errorf("expected %v, got %v", weather.ErrInvalidFormat, err)
	}
}
