# weather-widget-go

## env File

The `.env`-file has to at least contain the `API_KEY` value, mapped to your OpenWeatherMap API Key!
The `location` value is optional, in case you don't want to call the script with the location. When doing so, the script needs to be called with `.env` as the location parameter!

e.g. `./weather-widget-go small .env <path to env file>`

```.env
API_KEY=<OpenWeatherMap API Key>
LOCATION=<city name,state code,country code>
```

## External Dependencies

```text
github.com/joho/godotenv
```

### OpenWeatherMap

OpenWeatherMap provides an API to geocode the given location to latitude and longitude coordinates. In addition to that it provides an API to get the current weather of a location and the forecast.

Each request takes one ticket. Calling the script currently takes two credits, one for geocoding the location, and one to get the data.
