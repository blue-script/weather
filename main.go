package main

import (
	"flag"
	"fmt"

	"github.com/blue-script/weather/geo"
)

func main() {
	fmt.Println("New project")
	city := flag.String("city", "", "User city")
	format := flag.Int("format", 1, "Weather output format")
	flag.Parse()
	fmt.Println(*city, *format)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*geoData)

	// r := strings.NewReader("Hello, I'm a data stream")
	// block := make([]byte, 4)
	// for {
	// 	_, err := r.Read(block)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("%q\n", block)
	// }
}
