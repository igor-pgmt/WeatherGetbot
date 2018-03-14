# WeatherGetbot

This Telegram bot works on WebHooks and shows information about location: weather description and temperature.
Default language: Russian.
You can try it: [@WeatherGetbot](https://t.me/WeatherGetbot)

See PHP version [here](https://github.com/igor-pgmt/Weather2Getbot).

### How to interact with bot
Just send to him your location (tap on the attachment button and choose Location).

### How to install it on your server
  - [Create Your bot](https://core.telegram.org/bots#6-botfather) with [@Botfather](https://t.me/BotFather)
  - Register account on the [openweathermap.org](https://home.openweathermap.org/users/sign_up) ang get your API key [here](https://home.openweathermap.org/api_keys).
  - Generate certificate pair with FQDN name (or ip) of your server. You can [use self-signed certificates](https://core.telegram.org/bots/self-signed):
```sh
openssl req -newkey rsa:2048 -sha256 -nodes -keyout YOURPRIVATE.key -x509 -days 365 -out YOURPUBLIC.pem -subj "/C=RU/ST=MSC/L=Moscow/O=No/CN=YOURDOMAIN.COM"
```
  - Set WebHook with your pubic certificate and Telegram bot's API
```sh
curl -F "url=https://YOURDOMAIN.COM:8443/weatherInfo" -F "certificate=@YOURPUBLIC.pem" "https://api.telegram.org/botXXXXXXXXX:AAAAAAAAAAAAAAAAAAAAAAAA/setwebhook"
```
  - Get files from github:
```sh
git clone https://github.com/igor-pgmt/WeatherGetbot.git
```
  - Specify some parameters in the code:
**main.go:** specify your certificates
```go
const (
	PublicKey  = "YOURPUBLIC.pem"
	PrivateKey = "YOURPRIVATE.key"
)
```
If you want, you can specify another port to listen: 80, 8080 or 443
More info: https://core.telegram.org/bots/api
```go
	err := http.ListenAndServeTLS(":8443", PublicKey, PrivateKey, nil)
```
**TelegramBot\TelegramBot.go:** specify your Telegram token
```go
token   = "XXXXXXXXX:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
```
**OpenWeather\OpenWeather.go:** specify your OpenWeather token
```go
token   = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
```
Also you can change Bot's answers language:
```go
    lang    = "ru"     // en, ru, etc.
```
  - Then just build and run it:
```bash
cd WeatherGetbot
go build WeatherGetbot
sudo ./WeatherGetbot
```

#### More information

[Telegram's API](https://core.telegram.org/bots/api)

[OpenWeatherMap's API](https://openweathermap.org/current)
