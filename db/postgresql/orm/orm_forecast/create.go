package orm_forecast

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/fellippemendonca/gaia/services"
)

// Create model method
func Create(svc *services.Services, forecast *Forecast) (string, error) {
	query := `INSERT INTO forecasts(
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
	) VALUES(
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
		$21
	)`
	stmt, err := svc.Postgresql.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	res, err := stmt.Exec(
		forecast.Time,
		forecast.Latitude,
		forecast.Longitude,
		forecast.City,
		forecast.Country,
		forecast.Timezone,
		forecast.Sunrise,
		forecast.Sunset,
		forecast.Weather,
		forecast.WeatherDescription,
		forecast.Temperature,
		forecast.FeelsLike,
		forecast.TemperatureMin,
		forecast.TemperatureMax,
		forecast.Pressure,
		forecast.SeaLevel,
		forecast.GroundLevel,
		forecast.Humidity,
		forecast.CloudsCoverage,
		forecast.Visibility,
		forecast.RainAmount,
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Printf("row affected = %d\n", rowCnt)

	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
	case nil:
		return "", nil
	default:
		panic(err)
	}
	return "", nil
}
