package OpenWeather

type Reply struct {
	Coord      NestedCoord     `json: "coord"`
	Weather    []NestedWeather `json: "weather"`
	Base       string          `json: "base"`
	Main       NestedMain      `json: "main"`
	Visibility int             `json: "visibility"`
	Wind       NestedWind      `json: "wind"`
	Clouds     NestedClouds    `json: "clouds"`
	DT         int             `json: "DT"`
	Sys        NestedSys       `json: "sys"`
	ID         int             `json: "id"`
	Name       string          `json: "name"`
	COD        int             `json: "cod"`
}

type NestedCoord struct {
	Lat float64 `json: "lat"`
	Lon float64 `json: "lon"`
}

type NestedWeather struct {
	ID          int    `json: "id"`
	Main        string `json: "main"`
	Description string `json: "description"`
	Icon        string `json: "icon"`
}

type NestedMain struct {
	Temp     float64 `json: "temp"`
	Pressure float64 `json: "pressure"`
	Humidity float64 `json: "humidity"`
	Temp_min float64 `json: "temp_min"`
	Temp_max float64 `json: "temp_max"`
}

type NestedWind struct {
	Speed float64 `json: "speed"`
	Deg   int     `json: "deg"`
}

type NestedClouds struct {
	All int `json: "all"`
}

type NestedSys struct {
	Type    int     `json: "type"`
	ID      int     `json: "id"`
	Message float64 `json: "message"`
	Country string  `json: "country"`
	Sunrise int     `json: "sunrise"`
	Sunset  int     `json: "sunset"`
}
