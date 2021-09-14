package forecasts

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/db/postgresql/orm/orm_forecast"
	"github.com/fellippemendonca/gaia/services"
	"github.com/gin-gonic/gin"

	openWeatherMap "github.com/fellippemendonca/gaia/lib/openWeatherMap"
)

// GetForecastByCoord get weather forecasts from external service by Geolocation coordinates
func GetForecastByCoord(svc *services.Services) func(c *gin.Context) {
	return func(c *gin.Context) {
		lat, parseError := strconv.ParseFloat(c.DefaultQuery("latitude", "0.0"), 64)
		if parseError != nil {
			log.Println(parseError)
		}
		lon, parseError := strconv.ParseFloat(c.DefaultQuery("longitude", "0.0"), 64)
		if parseError != nil {
			log.Println(parseError)
		}

		var resWeather *openWeatherMap.ResponseGetForecast
		resWeather, openWeatherError := openWeatherMap.GetForecastByCoord(lat, lon)
		if openWeatherError != nil {
			log.Println(openWeatherError)
		}

		forecast := orm_forecast.Forecast{
			Time:      resWeather.List[0].Dt,
			Latitude:  resWeather.City.Coord.Lat,
			Longitude: resWeather.City.Coord.Lon,
			City:      resWeather.City.Name,
			Country:   resWeather.City.Country,
			Timezone:  resWeather.City.Timezone,
			Sunrise:   resWeather.City.Sunrise,
			Sunset:    resWeather.City.Sunset,

			Weather:            resWeather.List[0].Weather[0].Main,
			WeatherDescription: resWeather.List[0].Weather[0].Description,

			Temperature:    resWeather.List[0].Main.Temp,
			FeelsLike:      resWeather.List[0].Main.FeelsLike,
			TemperatureMin: resWeather.List[0].Main.TempMin,
			TemperatureMax: resWeather.List[0].Main.TempMax,
			Pressure:       resWeather.List[0].Main.Pressure,
			SeaLevel:       resWeather.List[0].Main.SeaLevel,
			GroundLevel:    resWeather.List[0].Main.GrndLevel,
			Humidity:       resWeather.List[0].Main.Humidity,

			CloudsCoverage: resWeather.List[0].Clouds.All,
			Visibility:     resWeather.List[0].Visibility,
			RainAmount:     resWeather.List[0].Rain.ThreeHour,
		}

		_, err := orm_forecast.Create(svc, &forecast)

		if err == nil {
			c.JSON(200, resWeather)
			return
		}
		log.Printf("The HTTP request failed with error %s\n", err)
		c.JSON(400, nil)
	}
}
