## gunderground
Go command client for accessing wunderground.com weather data

### compile
```bash
go build gunderground.go
```

### run
You will need a Weather Underground API key file in the same directory - [Link](http://www.wunderground.com/weather/api)

The key file should be named:
> _wunderground.key_

Usage:
```bash
./gunderground seattle,wa
```
Output:
```json
{
  "response": {
  "version":"0.1",
  "termsofService":"http://www.wunderground.com/weather/api/d/terms.html",
  "features": {
  "conditions": 1
  }
	}
  ,	"current_observation": {
		"image": {
		"url":"http://icons.wxug.com/graphics/wu2/logo_130x80.png",
		"title":"Weather Underground",
		"link":"http://www.wunderground.com"
		},
		"display_location": {
		"full":"Seattle, WA",
		"city":"Seattle",
		"state":"WA",
		"state_name":"Washington",
		"country":"US",
		"country_iso3166":"US",
		"zip":"98101",
		"magic":"1",
		"wmo":"99999",
		"latitude":"47.61167908",
		"longitude":"-122.33325958",
		"elevation":"63.00000000"
		},
		"observation_location": {
		"full":"Herrera, Inc., Seattle, Washington",
		"city":"Herrera, Inc., Seattle",
		"state":"Washington",
		"country":"US",
		"country_iso3166":"US",
		"latitude":"47.616558",
		"longitude":"-122.341240",
		"elevation":"121 ft"
		},
		"estimated": {
		},
		"station_id":"KWASEATT187",
		"observation_time":"Last Updated on September 24, 10:58 AM PDT",
		"observation_time_rfc822":"Thu, 24 Sep 2015 10:58:22 -0700",
		"observation_epoch":"1443117502",
		"local_time_rfc822":"Thu, 24 Sep 2015 10:58:30 -0700",
		"local_epoch":"1443117510",
		"local_tz_short":"PDT",
		"local_tz_long":"America/Los_Angeles",
		"local_tz_offset":"-0700",
		"weather":"Mostly Cloudy",
		"temperature_string":"64.3 F (17.9 C)",
		"temp_f":64.3,
		"temp_c":17.9,
		"relative_humidity":"66%",
		"wind_string":"Calm",
		"wind_dir":"ESE",
		"wind_degrees":121,
		"wind_mph":0.0,
		"wind_gust_mph":"10.0",
		"wind_kph":0,
		"wind_gust_kph":"16.1",
		"pressure_mb":"1016",
		"pressure_in":"30.00",
		"pressure_trend":"0",
		"dewpoint_string":"53 F (12 C)",
		"dewpoint_f":53,
		"dewpoint_c":12,
		"heat_index_string":"NA",
		"heat_index_f":"NA",
		"heat_index_c":"NA",
		"windchill_string":"NA",
		"windchill_f":"NA",
		"windchill_c":"NA",
		"feelslike_string":"64.3 F (17.9 C)",
		"feelslike_f":"64.3",
		"feelslike_c":"17.9",
		"visibility_mi":"10.0",
		"visibility_km":"16.1",
		"solarradiation":"--",
		"UV":"2","precip_1hr_string":"0.00 in ( 0 mm)",
		"precip_1hr_in":"0.00",
		"precip_1hr_metric":" 0",
		"precip_today_string":"0.00 in (0 mm)",
		"precip_today_in":"0.00",
		"precip_today_metric":"0",
		"icon":"mostlycloudy",
		"icon_url":"http://icons.wxug.com/i/c/k/mostlycloudy.gif",
		"forecast_url":"http://www.wunderground.com/US/WA/Seattle.html",
		"history_url":"http://www.wunderground.com/weatherstation/WXDailyHistory.asp?ID=KWASEATT187",
		"ob_url":"http://www.wunderground.com/cgi-bin/findweather/getForecast?query=47.616558,-122.341240",
		"nowcast":""
	}
}
```
