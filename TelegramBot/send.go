package TelegramBot

import "strconv"

// Data to send to Telegram's API server
type SendingData struct {
	ChatID int
	Text   string
	// more parameters ...
}

// Data to send to Telegram's API server in string type
type SendingStringedData struct {
	ChatID string
	Text   string
	// more parameters ...
}

// This function converts all the data to the string type for sending by the link
func (p SendingData) ToStr() *SendingStringedData {
	ssd := new(SendingStringedData)
	ssd.ChatID = strconv.Itoa(p.ChatID)
	ssd.Text = p.Text
	// more parameters ...

	return ssd
}
