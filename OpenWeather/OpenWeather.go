package OpenWeather

import "net/url"
import "net/http"
import "strconv"
import "io/ioutil"
import "encoding/json"

const (
    token   = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
    apiLink = "https://api.openweathermap.org/data/2.5/weather"
    lang    = "ru"     // en, ru, etc.
    mode    = "json"   // html, xml, json
    units   = "metric" // metric, imperial
)

// Helper function
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func GetWeather(lat, lon float64) *Reply {

    // Creating link for OpenWeather requests
    searchLink, err := url.Parse(apiLink)
    searchLink.Scheme = "https"
    check(err)
    parameters := url.Values{}
    parameters.Add("lang", lang)
    parameters.Add("mode", mode)
    parameters.Add("units", units)
    parameters.Add("appid", token)
    parameters.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
    parameters.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))
    searchLink.RawQuery = parameters.Encode()

    // Sending data and getting response
    response, err := http.Get(searchLink.String())
    check(err)

    resBody, err := ioutil.ReadAll(response.Body)
    check(err)

    // Populating struct with response data
    weatherInfo := new(Reply)
    err = json.Unmarshal(resBody, &weatherInfo)
    check(err)

    return weatherInfo

}
