package orm_forecast

import (
	"github.com/fellippemendonca/gaia/services"
)

const setupTable = `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE TABLE IF NOT EXISTS forecasts (
		id uuid DEFAULT uuid_generate_v4(),
		time bigint NOT NULL,
		latitude numeric NOT NULL,
		longitude numeric NOT NULL,
		city text NOT NULL,
		country text NOT NULL,
		timezone integer NOT NULL,
		sunrise bigint NOT NULL,
		sunset bigint NOT NULL,
		weather text NOT NULL,
		weather_description text NOT NULL,
		temperature numeric NOT NULL,
		feels_like numeric NOT NULL,
		temperature_min numeric NOT NULL,
		temperature_max numeric NOT NULL,
		pressure numeric NOT NULL,
		sea_level numeric NOT NULL,
		ground_level numeric NOT NULL,
		humidity numeric NOT NULL,
		clouds_coverage smallint NOT NULL,
		visibility smallint NOT NULL,
		rain_amount numeric NOT NULL
	);
`

func CreateForecastsTable(svc *services.Services) (string, error) {
	if _, err := svc.Postgresql.Exec(setupTable); err != nil {
		return "", err
	}
	return "Forecasts Table Created", nil
}
