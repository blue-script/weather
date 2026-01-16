package geo_test

import (
	"testing"

	"github.com/blue-script/weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "Smolensk"
	expected := geo.GeoData{
		City: "Smolensk",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)

	// Assert - проверка результата с expected
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) { 
	city := "Nocity123"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("expected %v, got %v", geo.ErrNoCity, err)
	}
}