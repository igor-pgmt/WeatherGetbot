package TelegramBot

import "net/url"
import "net/http"

// Helper function
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Your bot's token
const (
	token   = "XXXXXXXXX:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	apiLink = "https://api.telegram.org/bot" + token + "/"
)

func sendRequest(method string, data SendingStringedData) {

	// Creating link for Telegram requests
	link := apiLink + method + "?"
	searchLink, err := url.Parse(link)
	searchLink.Scheme = "https"
	check(err)
	parameters := url.Values{}
	parameters.Add("chat_id", data.ChatID)
	parameters.Add("text", data.Text)
	parameters.Add("parse_mode", "HTML")
	searchLink.RawQuery = parameters.Encode()

	// Sending link to Telegram's API server
	_, err = http.Get(searchLink.String())
	check(err)

}

func SendMessage(data SendingData) {
	strData := data.ToStr()
	sendRequest("sendMessage", *strData)
}
