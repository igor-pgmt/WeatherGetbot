package main

import "strconv"
import "net/http"
import "io/ioutil"
import "weather/TelegramBot"
import "weather/OpenWeather"
import "encoding/json"

// Your certificate pair
// https://core.telegram.org/bots/self-signed
const (
	PublicKey  = "YOURPUBLIC.pem"
	PrivateKey = "YOURPRIVATE.key"
)

// Helper function
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Set single handler for Telegram webhooks
	http.HandleFunc("/weatherInfo", telegramHandler)
	// Enable web-server
	err := http.ListenAndServeTLS(":8443", PublicKey, PrivateKey, nil)
	check(err)
}

// Handle Telegram webhooks
func telegramHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/weatherInfo" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST": //Telegram sends POST only

		reqBody, err := ioutil.ReadAll(r.Body)
		check(err)

		// Populate reply struct with json
		reply := new(TelegramBot.Reply)
		err = json.Unmarshal(reqBody, &reply)

		// If there is location data, show weather info
		if (reply.Message.Location.Latitude != 0) && (reply.Message.Location.Longitude != 0) {
			info := OpenWeather.GetWeather(reply.Message.Location.Latitude, reply.Message.Location.Longitude)
			text := info.Weather[0].Description + ", " + strconv.FormatFloat(info.Main.Temp, 'f', -1, 64) + "°C"
			TelegramBot.SendMessage(TelegramBot.SendingData{ChatID: reply.Message.Chat.ID, Text: text})

		} else {
			sendMeLocation := TelegramBot.SendingData{ChatID: reply.Message.Chat.ID, Text: "Пришлите локацию!"}
			TelegramBot.SendMessage(sendMeLocation)
		}

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
