package openWeatherMap

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Forecasts struct {
	Dt   int32 `json:"dt"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
		Humidity  float64 `json:"humidity"`
		TempKf    float64 `json:"temp_kf"`
	} `json:"main"`
	Weather []struct {
		ID          int32  `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int32 `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int32   `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Visibility int32   `json:"visibility"`
	Pop        float64 `json:"pop"`
	Rain       struct {
		ThreeHour float64 `json:"3h"`
	} `json:"rain"`
	Sys struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	Dt_txt string `json:"dt_txt"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type City struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int32  `json:"population"`
	Timezone   int32  `json:"timezone"`
	Sunrise    int32  `json:"sunrise"`
	Sunset     int32  `json:"sunset"`
}

type ResponseGetForecast struct {
	Cod     string      `json:"cod"`
	Message int32       `json:"message"`
	Cnt     int32       `json:"cnt"`
	List    []Forecasts `json:"list"`
	City    City        `json:"city"`
}

// GetForecastByCoord retrieve Forecast by geolocation coordinates
func GetForecastByCoord(lat float64, lon float64) (*ResponseGetForecast, error) {
	appId := "appid=abc"
	units := "units=metric"
	coord := "lat=" + fmt.Sprintf("%g", lat) + "&lon=" + fmt.Sprintf("%g", lon)
	query := "?" + appId + "&" + units + "&" + coord
	data, err := doRequest("GET", "/data/2.5/forecast"+query, nil)

	if err == nil {
		str := string(data)
		res := ResponseGetForecast{}
		json.Unmarshal([]byte(str), &res)
		return &res, nil
	}
	log.Println(err)
	return nil, err
}
