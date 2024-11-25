package types

import "github.com/golang-jwt/jwt/v4"

type Message struct {
	Topic    string   `json:"topic"`
	Title    string   `json:"title"`
	Tags     []string `json:"tags"`
	Payload  string   `json:"payload"`
	Priority int      `json:"priority"`
}

type OpenWeatherLocation struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Param struct {
	Name         string
	DefaultValue string
	Optional     bool
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OkJsonMessage struct {
	Message string `json:"message"`
}

type OkCreatedJsonMessage struct {
	Message string `json:"message"`
	Data	interface{} `json:"data"`
}

type TokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Attribute struct {
	Name  string `json:"name"`
	Id    int32  `json:"id"`
	Desc  string `json:"description"`
}

type Location struct {
	Name string `json:"name"`
	Id   int32  `json:"id"`
	Desc string `json:"description"`
}

type Sensor struct {
	Name string `json:"name"`
	Id   int32  `json:"id"`
	Desc string `json:"description"`
}

type NewDevice struct {
	Name       string  `json:"name"`
	Id	       int32   `json:"id"`
	Desc	   string  `json:"description"`
	IPAddress  string  `json:"ip_address"`
	Port       int32   `json:"port"`
	MACAddress string  `json:"mac_address"`
	Location   int32   `json:"location"`
	Sensors    []int32 `json:"sensors"`
}