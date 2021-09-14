package orm_forecast

import (
	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
)

// List model method
func List(svc *services.Services) ([]*Forecast, error) {
	query := `SELECT
		id,
		time,
		latitude,
		longitude,
		city,
		country,
		timezone,
		sunrise,
		sunset,
		weather,
		weather_description,
		temperature,
		feels_like,
		temperature_min,
		temperature_max,
		pressure,
		sea_level,
		ground_level,
		humidity,
		clouds_coverage,
		visibility,
		rain_amount
		FROM forecasts;
	`
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	forecasts := make([]*Forecast, 0)
	for rows.Next() {
		forecast := new(Forecast)
		err := rows.Scan(
			&forecast.ID,
			&forecast.Time,
			&forecast.Latitude,
			&forecast.Longitude,
			&forecast.City,
			&forecast.Country,
			&forecast.Timezone,
			&forecast.Sunrise,
			&forecast.Sunset,
			&forecast.Weather,
			&forecast.WeatherDescription,
			&forecast.Temperature,
			&forecast.FeelsLike,
			&forecast.TemperatureMin,
			&forecast.TemperatureMax,
			&forecast.Pressure,
			&forecast.SeaLevel,
			&forecast.GroundLevel,
			&forecast.Humidity,
			&forecast.CloudsCoverage,
			&forecast.Visibility,
			&forecast.RainAmount,
		)
		if err != nil {
			log.Fatal(err)
		}
		forecasts = append(forecasts, forecast)
	}
	return forecasts, nil
}
