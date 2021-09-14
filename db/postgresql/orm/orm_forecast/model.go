package orm_forecast

// Forecast struct
type Forecast struct {
	ID        string  `json:"id,omitempty"`
	Time      int32   `json:"time"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Timezone  int32   `json:"timezone"`
	Sunrise   int32   `json:"sunrise"`
	Sunset    int32   `json:"sunset"`

	Weather            string `json:"weather"`
	WeatherDescription string `json:"weatherDescription"`

	Temperature    float64 `json:"temperature"`
	FeelsLike      float64 `json:"feelsLike"`
	TemperatureMin float64 `json:"temperatureMin"`
	TemperatureMax float64 `json:"temperatureMax"`
	Pressure       float64 `json:"pressure"`
	SeaLevel       float64 `json:"seaLevel"`
	GroundLevel    float64 `json:"groundLevel"`
	Humidity       float64 `json:"humidity"`

	CloudsCoverage int32   `json:"cloudsCoverage"`
	Visibility     int32   `json:"visibility"`
	RainAmount     float64 `json:"rainAmount"`
}
