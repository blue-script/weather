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

var testCases = []struct {
	name   string
	format int
}{
	{name: "147 format", format: 147},
	{name: "0 format", format: 0},
	{name: "-7 format", format: -7},
}

func TestGetWeatherInvalidFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Smolensk"
			geoData := geo.GeoData{
				City: expected,
			}

			_, err := weather.GetWeather(geoData, tc.format)

			if err == nil {
				t.Fatalf("expected error %v, got nil", weather.ErrInvalidFormat)
			}
			if !errors.Is(err, weather.ErrInvalidFormat) {
				t.Errorf("expected %v, got %v", weather.ErrInvalidFormat, err)
			}
		})
	}
}
