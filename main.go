package main

import (
	"flag"
	"fmt"

	"github.com/blue-script/weather/geo"
	"github.com/blue-script/weather/weather"
)

func main() {
	fmt.Println("Weather")
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 4, "Weather output format")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*geoData)

	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(weatherData)
}
