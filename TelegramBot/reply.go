package TelegramBot

type Reply struct {
	Message   NestedMsg `json: "message"`
	Update_ID int       `json: "update_id"`
}

type NestedMsg struct {
	Chat       NestedChat     `json: "chat"`
	Date       int            `json: "date"`
	From       NestedFrom     `json: "from"`
	Message_ID int            `json: "message_id"`
	Text       string         `json: "text"`
	Location   NestedLocation `json: "location"`
}

type NestedChat struct {
	First_Name string `json: "first_name"`
	ID         int    `json: "id"`
	Last_Name  string `json: "last_name"`
	Type       string `json: "type"`
	Username   string `json: "username"`
}

type NestedFrom struct {
	First_Name    string `json: "first_name"`
	ID            int    `json: "id"`
	IsBot         bool   `json: "is_bot"`
	Language_Code string `json: "language_code"`
	Last_Name     string `json: "last_name"`
	Username      string `json: "username"`
}

type NestedLocation struct {
	Latitude  float64 `json: "latitude"`
	Longitude float64 `json: "longitude"`
}
