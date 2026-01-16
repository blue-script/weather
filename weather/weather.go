package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/blue-script/weather/geo"
)

var ErrInvalidFormat = errors.New("INCORRECT_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format > 4 || format < 1 {
		return "", ErrInvalidFormat
	}

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_URL")
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_HTTP")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR_READ_BODY")
	}

	return string(body), nil
}
