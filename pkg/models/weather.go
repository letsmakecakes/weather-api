package models

type Weather struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	TZOffset        float64 `json:"tzoffset"`
	Days            []Day   `json:"days"`
}

type Day struct {
	DateTime       string   `json:"datetime"`
	DateTimeEpoch  int64    `json:"datetimeEpoch"`
	TempMax        float64  `json:"tempmax"`
	TempMin        float64  `json:"tempmin"`
	Temp           float64  `json:"temp"`
	FeelsLikeMax   float64  `json:"feelslikemax"`
	FeelsLikeMin   float64  `json:"feelslikemin"`
	FeelsLike      float64  `json:"feelslike"`
	Dew            float64  `json:"dew"`
	Humidity       float64  `json:"humidity"`
	Precip         float64  `json:"precip"`
	PrecipProb     float64  `json:"precipprob"`
	PrecipCover    float64  `json:"precipcover"`
	PrecipType     []string `json:"preciptype"`
	Snow           float64  `json:"snow"`
	SnowDepth      float64  `json:"snowdepth"`
	WindGust       float64  `json:"windgust"`
	WindSpeed      float64  `json:"windspeed"`
	WindDir        float64  `json:"winddir"`
	Pressure       float64  `json:"pressure"`
	CloudCover     float64  `json:"cloudcover"`
	Visibility     float64  `json:"visibility"`
	SolarRadiation float64  `json:"solarradiation"`
	SolarEnergy    float64  `json:"solarenergy"`
	UVIndex        float64  `json:"uvindex"`
	SevereRisk     float64  `json:"severerisk"`
	Sunrise        string   `json:"sunrise"`
	SunriseEpoch   int64    `json:"sunriseEpoch"`
	Sunset         string   `json:"sunset"`
	SunsetEpoch    int64    `json:"sunsetEpoch"`
	MoonPhase      float64  `json:"moonphase"`
	Conditions     string   `json:"conditions"`
	Description    string   `json:"description"`
	Icon           string   `json:"icon"`
	Stations       []string `json:"stations"`
	Source         string   `json:"source"`
}
