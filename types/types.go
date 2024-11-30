package types

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v4"
)

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

type List struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
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

type NewMessage struct {
	Topic    string `json:"topic"`
	Title    string `json:"title"`
	Tags     string `json:"tags"`
	Payload  string `json:"payload"`
	Priority int    `json:"priority"`
}

type Message struct {
	Id	     int32  `json:"id"`
	Topic    string `json:"topic"`
	Title    string `json:"title"`
	Tags     string `json:"tags"`
	Payload  string `json:"payload"`
	Priority int    `json:"priority"`
}

type Messages struct {
	Id	  int32  `json:"id"`
	Title string `json:"title"`
}

type Alert struct {
	// Id   int32  `json:"id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	DeviceId int32 `json:"device_id"`
	SensorId int32 `json:"sensor_id"`
	AttributeId int32 `json:"attribute_id"`
	OperatorId int32 `json:"operator_id"`
	MessageId int32 `json:"message_id"`
	Frequency int32 `json:"frequency"`
	Value1 float64 `json:"value1"`
	Value2 sql.NullFloat64 `json:"value2"`
}