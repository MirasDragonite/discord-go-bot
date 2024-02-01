package internal

type Coordinate struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

type Main struct {
	Temp  float64 `json:"temp"`
	Feels float64 `json:"feels_like"`
}
type Weth struct {
	Descriotion string `json:"description"`
}
type Weather struct {
	Main Main   `json:"main"`
	Weth []Weth `json:"weather"`
}
