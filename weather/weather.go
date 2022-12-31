package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"weather-widget-go/weather-widget-go/weather/request"

	"github.com/joho/godotenv"
)

// OpenWeatherMap api key
var apiKey string = getEnv("API_KEY")

// return weather depending on mode, either as a one liner or full text
func WeatherText(mode string, location string) string {
	lat, lon, locationName := geocodeLocation(location)
	weatherData := getWeatherData(lat, lon)

	switch mode {
	case "small":
		// get chosen location, temperature as float with least amount of numbers, weather description
		currentWeather := fmt.Sprintf("%s: %s°C %s", locationName,
			strconv.FormatFloat(weatherData.Current.Temp, 'f', -1, 64),
			weatherData.Current.Weather[0].Description)
		return currentWeather

	case "full":
		currentWeather := fmt.Sprintf("%s: \n\t%s, %s\n\tTemperature: %s °C\n\tHumidity: %d hpa\n\tPressure: %d\n\tWind: %f m/s from %d",
			locationName,
			weatherData.Current.Weather[0].Main,
			weatherData.Current.Weather[0].Description,
			strconv.FormatFloat(weatherData.Current.Temp, 'f', -1, 64),
			weatherData.Current.Humidity,
			weatherData.Current.Pressure,
			weatherData.Current.WindSpeed,
			weatherData.Current.WindDeg)
		return currentWeather

	case "forecast":
		// TODO make extra function for building forecast string
		return ""

	default:
		return "Error: Invalid Mode"
	}
}

// get weather data of a location, contains current, and hourly predictions
func getWeatherData(lat float64, lon float64) getWeatherResponse {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&exclude=minutely,daily,alerts&appid=%s&units=metric", lat, lon, apiKey)

	var responseData = request.Request(url)
	var responseObject getWeatherResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

// get latitude, longitude and name of given location
// name is not needed for further operations, but can be used for better understanding of what location has been detected
func geocodeLocation(location string) (lat float64, lon float64, name string) {
	// parse url
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", location, apiKey)

	var responseData = request.Request(url)

	var responseObject geocodeResponse
	json.Unmarshal(responseData, &responseObject)

	if len(responseObject) == 0 {
		fmt.Fprintf(os.Stderr, "No valid locations found! ")
		os.Exit(1)
	}

	return responseObject[0].Lat, responseObject[0].Lon, responseObject[0].Name
}

// get environment variables from .env file
// file needs to be in top folder next to main.go
func getEnv(key string) string {
	if len(os.Args) < 4 {
		log.Fatalf("path to .env file not given")
	}

	err := godotenv.Load(os.Args[3])

	if err != nil {
		log.Fatalf("Failed to load .env file: " + err.Error())
	}

	return os.Getenv(key)
}

// response when querying onecall api
type getWeatherResponse struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	Timezone       string   `json:"timezone"`
	TimezoneOffset int      `json:"timezone_offset"`
	Current        Current  `json:"current"`
	Hourly         []Hourly `json:"hourly"`
}

// type of weather data
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// current weather data
type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        int       `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

// snow specific data
type Snow struct {
	OneH float64 `json:"1h"`
}

// hourly weather data
type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        int       `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
	Snow       Snow      `json:"snow,omitempty"`
}

// struct for geocoding response
type geocodeResponse []struct {
	Name       string     `json:"name"`
	LocalNames LocalNames `json:"local_names"`
	Lat        float64    `json:"lat"`
	Lon        float64    `json:"lon"`
	Country    string     `json:"country"`
}

// localization names, use later for language specific naming
type LocalNames struct {
	Fi string `json:"fi"`
	Rm string `json:"rm"`
	Nl string `json:"nl"`
	Mg string `json:"mg"`
	Lb string `json:"lb"`
	So string `json:"so"`
	Qu string `json:"qu"`
	Fa string `json:"fa"`
	De string `json:"de"`
	Kk string `json:"kk"`
	Fy string `json:"fy"`
	It string `json:"it"`
	Et string `json:"et"`
	Ga string `json:"ga"`
	Ms string `json:"ms"`
	Ie string `json:"ie"`
	Vi string `json:"vi"`
	Da string `json:"da"`
	Br string `json:"br"`
	Sv string `json:"sv"`
	Sl string `json:"sl"`
	Ht string `json:"ht"`
	Am string `json:"am"`
	Sr string `json:"sr"`
	Mr string `json:"mr"`
	Uz string `json:"uz"`
	Bs string `json:"bs"`
	Ro string `json:"ro"`
	Ky string `json:"ky"`
	Af string `json:"af"`
	Sh string `json:"sh"`
	Sq string `json:"sq"`
	Yo string `json:"yo"`
	Kv string `json:"kv"`
	Hu string `json:"hu"`
	Ug string `json:"ug"`
	Os string `json:"os"`
	ID string `json:"id"`
	Ty string `json:"ty"`
	Cs string `json:"cs"`
	Ja string `json:"ja"`
	An string `json:"an"`
	Jv string `json:"jv"`
	Sw string `json:"sw"`
	Sk string `json:"sk"`
	Oc string `json:"oc"`
	Ln string `json:"ln"`
	Lt string `json:"lt"`
	Ps string `json:"ps"`
	Ar string `json:"ar"`
	Ur string `json:"ur"`
	Fo string `json:"fo"`
	Zu string `json:"zu"`
	Iu string `json:"iu"`
	Tg string `json:"tg"`
	Se string `json:"se"`
	Si string `json:"si"`
	Az string `json:"az"`
	Eu string `json:"eu"`
	Eo string `json:"eo"`
	Mt string `json:"mt"`
	He string `json:"he"`
	My string `json:"my"`
	Te string `json:"te"`
	Tl string `json:"tl"`
	Es string `json:"es"`
	Bo string `json:"bo"`
	Pl string `json:"pl"`
	Lv string `json:"lv"`
	El string `json:"el"`
	Hy string `json:"hy"`
	Co string `json:"co"`
	Kn string `json:"kn"`
	Tt string `json:"tt"`
	Bg string `json:"bg"`
	Gl string `json:"gl"`
	La string `json:"la"`
	Gv string `json:"gv"`
	Is string `json:"is"`
	Yi string `json:"yi"`
	Wo string `json:"wo"`
	Kw string `json:"kw"`
	Ka string `json:"ka"`
	Ia string `json:"ia"`
	Cu string `json:"cu"`
	En string `json:"en"`
	Uk string `json:"uk"`
	Na string `json:"na"`
	Be string `json:"be"`
	Hr string `json:"hr"`
	Mk string `json:"mk"`
	Nn string `json:"nn"`
	Th string `json:"th"`
	Ko string `json:"ko"`
	Mn string `json:"mn"`
	Su string `json:"su"`
	Bi string `json:"bi"`
	Ml string `json:"ml"`
	Cv string `json:"cv"`
	No string `json:"no"`
	Fr string `json:"fr"`
	Io string `json:"io"`
	Ta string `json:"ta"`
	Gd string `json:"gd"`
	Pt string `json:"pt"`
	Lg string `json:"lg"`
	Ru string `json:"ru"`
	Cy string `json:"cy"`
	Ab string `json:"ab"`
	Tr string `json:"tr"`
	Gn string `json:"gn"`
	Ba string `json:"ba"`
	Ca string `json:"ca"`
	Ku string `json:"ku"`
	Hi string `json:"hi"`
	Mi string `json:"mi"`
	Li string `json:"li"`
	Bn string `json:"bn"`
	Zh string `json:"zh"`
	Sc string `json:"sc"`
}
