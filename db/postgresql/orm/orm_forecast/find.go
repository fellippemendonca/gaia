package orm_forecast

import (
	"database/sql"

	"github.com/fellippemendonca/gaia/services"
	log "github.com/sirupsen/logrus"
)

// GetByID model method
func GetByID(svc *services.Services, id string) (*Forecast, error) {
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
	 	FROM forecasts where id = $1;`
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	forecast := new(Forecast)
	row := stmt.QueryRow(id)
	err = row.Scan(
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
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		return forecast, nil
	default:
		panic(err)
	}
	return nil, nil
}
